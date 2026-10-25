// Command wavcompare reports amplitude and spectral differences between
// two mono 16-bit PCM WAV recordings. It does not align recordings.
package main

import (
 "encoding/binary"
 "encoding/json"
 "errors"
 "flag"
 "fmt"
 "math"
 "os"
)

type audio struct {rate uint32;samples []float64}
func readWAV(path string)(audio,error){
 data,err:=os.ReadFile(path);if err!=nil{return audio{},err}
 if len(data)<44||string(data[:4])!="RIFF"||string(data[8:12])!="WAVE" {return audio{},errors.New("not a RIFF/WAVE file")}
 var format,channels,bits uint16
 var rate uint32
 var pcm []byte
 for pos:=12;pos+8<=len(data);{
  size:=int(binary.LittleEndian.Uint32(data[pos+4:pos+8]))
  start:=pos+8
  if size<0||start+size>len(data){return audio{},errors.New("invalid WAV chunk")}
  switch string(data[pos:pos+4]){
  case "fmt ":
   if size<16{return audio{},errors.New("short WAV format")}
   format=binary.LittleEndian.Uint16(data[start:])
   channels=binary.LittleEndian.Uint16(data[start+2:])
   rate=binary.LittleEndian.Uint32(data[start+4:])
   bits=binary.LittleEndian.Uint16(data[start+14:])
  case "data":pcm=data[start:start+size]
  }
  pos=start+size+(size&1)
 }
 if format!=1||channels!=1||bits!=16||rate==0||len(pcm)==0||len(pcm)%2!=0{return audio{},errors.New("requires mono 16-bit PCM WAV")}
 out:=audio{rate:rate,samples:make([]float64,len(pcm)/2)}
 for i:=range out.samples {out.samples[i]=float64(int16(binary.LittleEndian.Uint16(pcm[i*2:])))/32768}
 return out,nil
}
func rms(x []float64)float64{
 var s float64
 for _,v:=range x{s+=v*v}
 return math.Sqrt(s/float64(len(x)))
}
func magnitude(x []float64,frequency float64,rate uint32)float64{
 var re,im float64
 for i,v:=range x {
  angle:=2*math.Pi*frequency*float64(i)/float64(rate)
  re+=v*math.Cos(angle)
  im-=v*math.Sin(angle)
 }
 return 2*math.Hypot(re,im)/float64(len(x))
}
func main(){
 reference:=flag.String("reference","","reference mono WAV")
 candidate:=flag.String("candidate","","candidate mono WAV")
 frequency:=flag.Float64("frequency",0,"fundamental frequency in Hz")
 harmonics:=flag.Int("harmonics",8,"number of harmonics to compare")
 flag.Parse()
 if *reference==""||*candidate==""||math.IsNaN(*frequency)||math.IsInf(*frequency,0)||*frequency<=0||*harmonics<1||*harmonics>32 {fmt.Fprintln(os.Stderr,"invalid arguments");os.Exit(2)}
 a,err:=readWAV(*reference);if err!=nil{fmt.Fprintln(os.Stderr,err);os.Exit(1)}
 b,err:=readWAV(*candidate);if err!=nil{fmt.Fprintln(os.Stderr,err);os.Exit(1)}
 if a.rate!=b.rate||len(a.samples)!=len(b.samples)||*frequency*float64(*harmonics)>=float64(a.rate)/2 {fmt.Fprintln(os.Stderr,"sample rates, durations or Nyquist constraint mismatch");os.Exit(1)}
 type partial struct {Harmonic int `json:"harmonic"`; FrequencyHz float64 `json:"frequency_hz"`; Reference float64 `json:"reference_amplitude"`; Candidate float64 `json:"candidate_amplitude"`}
 result:=struct{SampleRate uint32 `json:"sample_rate"`;Samples int `json:"samples"`;ReferenceRMS float64 `json:"reference_rms"`;CandidateRMS float64 `json:"candidate_rms"`;Harmonics []partial `json:"harmonics"`}{a.rate,len(a.samples),rms(a.samples),rms(b.samples),nil}
 for h:=1;h<=*harmonics;h++ {
  f:=float64(h)**frequency
  result.Harmonics=append(result.Harmonics,partial{h,f,magnitude(a.samples,f,a.rate),magnitude(b.samples,f,b.rate)})
 }
 if err:=json.NewEncoder(os.Stdout).Encode(result);err!=nil{fmt.Fprintln(os.Stderr,err);os.Exit(1)}
}
