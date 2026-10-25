package engine

import (
 "bytes"
 "encoding/binary"
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
 "github.com/cpu-model/hammond/internal/wav"
)

// Full path: documented B3 contact routing -> shared mechanical generator
// -> drawbar gain -> two-manual mix -> PCM WAV. No invented routing fixture.
func TestB3AudioToWAV(t *testing.T) {
 route:=registration.B3ManualRouting()
 e,err:=NewTimed(48000,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil { t.Fatal(err) }
 upper,_:=e.Manual(0)
 lower,_:=e.Manual(1)
 for _,m:=range []*registration.TimedManual{upper,lower} {
  if err:=m.Key(36,true);err!=nil { t.Fatal(err) }
  if err:=m.SetDrawbar(2,8);err!=nil { t.Fatal(err) }
 }
 samples:=make([]float64,480)
 e.Render(samples)
 if samples[0]!=0 { t.Fatal("expected zero initial phase") }
 peak:=0.0
 for _,s:=range samples {
  if math.IsNaN(s)||math.IsInf(s,0) { t.Fatal("non-finite output") }
  if math.Abs(s)>peak { peak=math.Abs(s) }
 }
 if peak<0.1 { t.Fatalf("unexpectedly silent: peak %v",peak) }
 // The same shared wheel on two manuals should double the signal.
 single,err:=NewTimed(48000,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil { t.Fatal(err) }
 one,_:=single.Manual(0)
 _=one.Key(36,true)
 _=one.SetDrawbar(2,8)
 for n:=0;n<len(samples);n++ {
  if math.Abs(samples[n]-2*single.Next())>1e-9 {
   t.Fatalf("manual sharing mismatch at sample %d",n)
  }
 }
 // Normalize before PCM conversion to avoid deliberate clipping.
 for i:=range samples { samples[i]*=0.4/peak }
 var out bytes.Buffer
 if err:=wav.WriteMono16(&out,48000,samples);err!=nil { t.Fatal(err) }
 b:=out.Bytes()
 if len(b)!=44+len(samples)*2 { t.Fatalf("WAV length %d",len(b)) }
 if string(b[:4])!="RIFF" || string(b[8:12])!="WAVE" || string(b[36:40])!="data" { t.Fatal("bad WAV header") }
 if got:=binary.LittleEndian.Uint32(b[40:44]);got!=uint32(len(samples)*2) { t.Fatalf("data size %d",got) }
}
