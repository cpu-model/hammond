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

// Offline single-key demonstration using the transcribed B3 contact matrix.
// This is a linear ideal-sine mix, not an authentic B3 audio emulation.
func main() {
 out:=flag.String("out","manual.wav","output WAV file")
 key:=flag.Int("key",37,"key number 1..61 (37 is middle C)")
 seconds:=flag.Int("seconds",3,"duration 1..30 seconds")
 flag.Parse()
 if *key<1 || *key>61 || *seconds<1 || *seconds>30 { log.Fatal("key must be 1..61 and seconds 1..30") }
 const rate=48000
 routing:=registration.B3ManualRouting()
 e,err:=engine.New(rate,routing,routing)
 if err!=nil { log.Fatal(err) }
 upper,_:=e.Manual(0)
 if err:=upper.Key(*key-1,true);err!=nil { log.Fatal(err) }
 // 888000000: first three drawbars fully out.
 for d:=0;d<3;d++ { if err:=upper.SetDrawbar(d,8);err!=nil { log.Fatal(err) } }
 samples:=make([]float64,rate**seconds)
 e.Render(samples)
 // Conservative peak normalization for this offline demonstration.
 peak:=0.0
 for _,x:=range samples { if math.Abs(x)>peak { peak=math.Abs(x) } }
 if peak>0 { for i:=range samples { samples[i]*=0.75/peak } }
 f,err:=os.Create(*out)
 if err!=nil { log.Fatal(err) }
 if err:=wav.WriteMono16(f,rate,samples);err!=nil { _=f.Close();log.Fatal(err) }
 if err:=f.Close();err!=nil { log.Fatal(err) }
 log.Printf("wrote %s (key %d, registration 888000000, ideal sine mix)",*out,*key)
}
