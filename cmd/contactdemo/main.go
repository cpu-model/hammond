package main

import (
 "flag"
 "log"
 "math"
 "os"

 "github.com/cpu-model/hammond/internal/engine"
 "github.com/cpu-model/hammond/internal/registration"
 "github.com/cpu-model/hammond/internal/wav"
)

// A controlled contact-timing experiment, NOT measured B3 key click.
// All nine contacts use the same user-selected step between closures.
func main() {
 out:=flag.String("out","contacts.wav","output WAV file")
 key:=flag.Int("key",37,"key number 1..61")
 step:=flag.Int("step",0,"additional make delay per successive contact, in samples (0..100)")
 seconds:=flag.Int("seconds",3,"total duration in seconds (1..30)")
 flag.Parse()
 if *key<1||*key>61||*step<0||*step>100||*seconds<1||*seconds>30 {
  log.Fatal("invalid key, step or seconds")
 }
 const rate=48000
 var timing registration.ContactTiming
 for d:=0;d<registration.Harmonics;d++ {
  timing.Make[d]=uint32(d**step)
 }
 route:=registration.B3ManualRouting()
 e,err:=engine.NewTimed(rate,route,route,timing,registration.ContactTiming{})
 if err!=nil { log.Fatal(err) }
 upper,_:=e.Manual(0)
 for d:=0;d<3;d++ {
  if err:=upper.SetDrawbar(d,8);err!=nil { log.Fatal(err) }
 }
 if err:=upper.Key(*key-1,true);err!=nil { log.Fatal(err) }
 samples:=make([]float64,rate**seconds)
 release:=len(samples)*2/3
 for n:=range samples {
  if n==release {
   if err:=upper.Key(*key-1,false);err!=nil { log.Fatal(err) }
  }
  samples[n]=e.Next()
 }
 peak:=0.0
 for _,x:=range samples { if math.Abs(x)>peak { peak=math.Abs(x) } }
 if peak>0 { for n:=range samples { samples[n]*=0.75/peak } }
 f,err:=os.Create(*out)
 if err!=nil { log.Fatal(err) }
 if err:=wav.WriteMono16(f,rate,samples);err!=nil { _=f.Close();log.Fatal(err) }
 if err:=f.Close();err!=nil { log.Fatal(err) }
 log.Printf("wrote %s (key %d, step %d samples, uncalibrated contact timing)",*out,*key,*step)
}
