package engine

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
)

func TestPercussionControlSelections(t *testing.T) {
 route:=registration.B3ManualRouting()
 e,err:=NewTimed(48000,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil{t.Fatal(err)}
 c:=DefaultPercussionControls()
 for _,tc:=range []struct{fast,soft bool;seconds,level float64}{
  {true,false,c.FastSeconds,c.NormalLevel},
  {true,true,c.FastSeconds,c.SoftLevel},
  {false,false,c.SlowSeconds,c.NormalLevel},
  {false,true,c.SlowSeconds,c.SoftLevel},
 } {
  c.Fast=tc.fast;c.Soft=tc.soft
  if err:=e.SetPercussionControls(0,c);err!=nil{t.Fatal(err)}
  p:=e.percussion[0]
  if p.harmonic!=3||p.level!=tc.level||math.Abs(p.decay-math.Exp(-1/(48000*tc.seconds)))>1e-14 {
   t.Fatalf("incorrect controls %+v: %+v",tc,p)
  }
 }
}

func TestPercussionControlsRejectInvalidValues(t *testing.T) {
 route:=registration.B3ManualRouting()
 e,err:=NewTimed(48000,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil{t.Fatal(err)}
 c:=DefaultPercussionControls()
 c.SlowSeconds=math.NaN()
 if err:=e.SetPercussionControls(0,c);err==nil{t.Fatal("accepted NaN")}
}
