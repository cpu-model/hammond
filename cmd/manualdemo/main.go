package main

import (
 "flag"
 "log"
 "os"

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
 flag.Parse()
 if *key<1 || *key>61 || *seconds<1 || *seconds>30 { log.Fatal("key must be 1..61 and seconds 1..30") }
 if len(*drawbars)!=9 {log.Fatal("drawbars must contain nine digits 0..8")}
 for _,ch:=range *drawbars {if ch<'0'||ch>'8' {log.Fatal("drawbars must contain nine digits 0..8")}}
 if *release<0 || *release>float64(*seconds) {log.Fatal("release must be between 0 and duration")}
 if *gain<0 || *gain>1 {log.Fatal("gain must be between 0 and 1")}
 const rate=48000
 routing:=registration.B3ManualRouting()
 e,err:=engine.New(rate,routing,routing)
 if err!=nil { log.Fatal(err) }
 upper,_:=e.Manual(0)
 if err:=upper.Key(*key-1,true);err!=nil { log.Fatal(err) }
 for d,ch:=range *drawbars {if err:=upper.SetDrawbar(d,int(ch-'0'));err!=nil {log.Fatal(err)}}
 samples:=make([]float64,rate**seconds)
 releaseSample:=int(*release*rate)
 if *release==0 {releaseSample=len(samples)}
 for i:=range samples {
  if i==releaseSample {if err:=upper.Key(*key-1,false);err!=nil {log.Fatal(err)}}
  e.Render(samples[i:i+1])
  samples[i]*=*gain
 }
 f,err:=os.Create(*out)
 if err!=nil { log.Fatal(err) }
 if err:=wav.WriteMono16(f,rate,samples);err!=nil { _=f.Close();log.Fatal(err) }
 if err:=f.Close();err!=nil { log.Fatal(err) }
 log.Printf("wrote %s (key %d, registration %s, release %.3fs, gain %.4f, ideal sine mix)",*out,*key,*drawbars,*release,*gain)
}
