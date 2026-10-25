package main

import (
 "encoding/json"
 "flag"
 "log"
 "math"
 "path/filepath"
 "os"
 "strconv"
 "strings"

 "github.com/cpu-model/hammond/internal/engine"
 "github.com/cpu-model/hammond/internal/registration"
 "github.com/cpu-model/hammond/internal/wav"
)

// Offline single-key demonstration using the transcribed B3 contact matrix.
// This is a linear ideal-sine mix, not an authentic B3 audio emulation.
func main() {
 out:=flag.String("out","manual.wav","output WAV file")
 key:=flag.Int("key",37,"key number 1..61 (37 is middle C)")
 seconds:=flag.Int("seconds",3,"duration 1..30 seconds")
 drawbars:=flag.String("drawbars","888000000","nine drawbar positions 0..8")
 release:=flag.Float64("release",2,"key release time in seconds (0 means held throughout)")
 gain:=flag.Float64("gain",0.02,"fixed linear output gain")
 step:=flag.Int("step",0,"additional contact closure delay per drawbar in samples (0..100)")
 harmonic2:=flag.Float64("harmonic2",0,"experimental second-harmonic level for all tonewheels (0..1)")
 metrics:=flag.String("metrics","","optional JSON path for reproducible render measurements")
 wheelHarmonics:=flag.String("wheel-harmonics","","comma-separated wheel:level overrides, e.g. 37:0.2,49:0.4")
 wheelPhases:=flag.String("wheel-phases","","comma-separated wheel:phase-in-cycles overrides, e.g. 37:0.25,49:0.5")
 flag.Parse()
 if *key<1 || *key>61 || *seconds<1 || *seconds>30 { log.Fatal("key must be 1..61 and seconds 1..30") }
 if len(*drawbars)!=9 {log.Fatal("drawbars must contain nine digits 0..8")}
 for _,ch:=range *drawbars {if ch<'0'||ch>'8' {log.Fatal("drawbars must contain nine digits 0..8")}}
 if math.IsNaN(*release)||math.IsInf(*release,0)||*release<0 || *release>float64(*seconds) {log.Fatal("release must be between 0 and duration")}
 if math.IsNaN(*gain)||math.IsInf(*gain,0)||*gain<0 || *gain>1 {log.Fatal("gain must be between 0 and 1")}
 if *step<0||*step>100 {log.Fatal("step must be between 0 and 100")}
 if math.IsNaN(*harmonic2)||math.IsInf(*harmonic2,0)||*harmonic2<0||*harmonic2>1 {log.Fatal("harmonic2 must be finite and between 0 and 1")}
 const rate=48000
 routing:=registration.B3ManualRouting()
 var timing registration.ContactTiming
 for d:=0;d<registration.Harmonics;d++ {timing.Make[d]=uint32(d**step)}
 e,err:=engine.NewTimed(rate,routing,routing,timing,registration.ContactTiming{})
 if err!=nil { log.Fatal(err) }
 if err:=e.SetSecondHarmonic(*harmonic2);err!=nil {log.Fatal(err)}
 if *wheelHarmonics!="" {
  seen:=make(map[int]bool)
  for _,item:=range strings.Split(*wheelHarmonics,",") {
   parts:=strings.Split(item,":")
   if len(parts)!=2 {log.Fatalf("invalid wheel-harmonics entry %q",item)}
   wheel,err:=strconv.Atoi(parts[0])
   if err!=nil {log.Fatalf("invalid wheel number %q",parts[0])}
   level,err:=strconv.ParseFloat(parts[1],64)
   if err!=nil {log.Fatalf("invalid harmonic level %q",parts[1])}
   if seen[wheel] {log.Fatalf("duplicate wheel %d",wheel)}
   seen[wheel]=true
   if err:=e.SetWheelSecondHarmonic(wheel,level);err!=nil {log.Fatal(err)}
  }
 }
 if *wheelPhases!="" {
  seen:=make(map[int]bool)
  for _,item:=range strings.Split(*wheelPhases,",") {
   parts:=strings.Split(item,":")
   if len(parts)!=2 {log.Fatalf("invalid wheel-phases entry %q",item)}
   wheel,err:=strconv.Atoi(parts[0])
   if err!=nil {log.Fatalf("invalid phase wheel number %q",parts[0])}
   phase,err:=strconv.ParseFloat(parts[1],64)
   if err!=nil {log.Fatalf("invalid wheel phase %q",parts[1])}
   if seen[wheel] {log.Fatalf("duplicate phase wheel %d",wheel)}
   seen[wheel]=true
   if err:=e.SetWheelPhaseOffset(wheel,phase);err!=nil {log.Fatal(err)}
  }
 }
 upper,_:=e.Manual(0)
 if err:=upper.Key(*key-1,true);err!=nil { log.Fatal(err) }
 for d,ch:=range *drawbars {if err:=upper.SetDrawbar(d,int(ch-'0'));err!=nil {log.Fatal(err)}}
 samples:=make([]float64,rate**seconds)
 peak:=0.0
 clipped:=0
 sumSquares:=0.0
 releaseSample:=int(*release*rate)
 if *release==0 {releaseSample=len(samples)}
 for i:=range samples {
  if i==releaseSample {if err:=upper.Key(*key-1,false);err!=nil {log.Fatal(err)}}
  samples[i]=e.Next()
  samples[i]*=*gain
  sumSquares+=samples[i]*samples[i]
  if math.Abs(samples[i])>peak {peak=math.Abs(samples[i])}
  if math.Abs(samples[i])>1 {clipped++}
 }
 if *metrics!="" {
  record:=struct {
   SampleRate int `json:"sample_rate"`
   Key int `json:"key"`
   Drawbars string `json:"drawbars"`
   DurationSeconds int `json:"duration_seconds"`
   ReleaseSeconds float64 `json:"release_seconds"`
   Gain float64 `json:"gain"`
   ContactStepSamples int `json:"contact_step_samples"`
   SecondHarmonic float64 `json:"second_harmonic"`
   WheelHarmonics string `json:"wheel_harmonics"`
   WheelPhases string `json:"wheel_phases"`
   Peak float64 `json:"peak"`
   RMS float64 `json:"rms"`
   OverRangeSamples int `json:"over_range_samples"`
  }{rate,*key,*drawbars,*seconds,*release,*gain,*step,*harmonic2,*wheelHarmonics,*wheelPhases,peak,math.Sqrt(sumSquares/float64(len(samples))),clipped}
  data,err:=json.MarshalIndent(record,"","  ")
  if err!=nil {log.Fatal(err)}
  if err:=os.WriteFile(filepath.Clean(*metrics),append(data,'\n'),0644);err!=nil {log.Fatal(err)}
 }
 if clipped>0 {log.Printf("WARNING: %d samples exceed PCM full scale; lower -gain",clipped)}
 f,err:=os.Create(*out)
 if err!=nil { log.Fatal(err) }
 if err:=wav.WriteMono16(f,rate,samples);err!=nil { _=f.Close();log.Fatal(err) }
 if err:=f.Close();err!=nil { log.Fatal(err) }
 log.Printf("peak %.6f, over-range samples %d",peak,clipped)
 log.Printf("wrote %s (key %d, registration %s, release %.3fs, gain %.4f, contact step %d samples, second harmonic %.3f, wheel overrides %q, phase overrides %q, experimental mix)",*out,*key,*drawbars,*release,*gain,*step,*harmonic2,*wheelHarmonics,*wheelPhases)
}
