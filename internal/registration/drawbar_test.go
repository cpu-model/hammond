package registration

import (
 "math"
 "testing"
)

func TestDrawbarGain(t *testing.T) {
 if DrawbarGain(0)!=0 || DrawbarGain(8)!=1 { t.Fatal("endpoints") }
 for p:=1;p<8;p++ {
  got:=DrawbarGain(p)
  want:=math.Pow(10,-3*float64(8-p)/20)
  if math.Abs(got-want)>1e-12 { t.Fatalf("position %d: %v",p,got) }
  if !(got<DrawbarGain(p+1)) { t.Fatalf("not monotonic at %d",p) }
 }
}

func TestManualUsesSteppedGain(t *testing.T) {
 var routing Routing
 routing[0][0]=46
 m,err:=NewManual(routing)
 if err!=nil { t.Fatal(err) }
 _=m.Key(0,true)
 _=m.SetDrawbar(0,6)
 frame:=make([]float64,91)
 frame[45]=1
 if math.Abs(m.Mix(frame)-DrawbarGain(6))>1e-12 { t.Fatal("incorrect mixer gain") }
}
