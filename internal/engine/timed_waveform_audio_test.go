package engine

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
)

func TestTimedEngineWheelHarmonicAffectsOnlySelectedWheel(t *testing.T) {
 route:=registration.B3ManualRouting()
 a,err:=NewTimed(48000,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil {t.Fatal(err)}
 b,err:=NewTimed(48000,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil {t.Fatal(err)}
 if err:=b.SetWheelSecondHarmonic(int(route[36][0]),0.3);err!=nil {t.Fatal(err)}
 for _,e:=range []*TimedEngine{a,b} {
  upper,err:=e.Manual(0)
  if err!=nil {t.Fatal(err)}
  if err:=upper.SetDrawbar(0,8);err!=nil {t.Fatal(err)}
  if err:=upper.Key(36,true);err!=nil {t.Fatal(err)}
 }
 changed:=false
 for i:=0;i<512;i++ {
  x,y:=a.Next(),b.Next()
  if math.Abs(x-y)>1e-8 {changed=true}
 }
 if !changed {t.Fatal("selected wheel harmonic did not affect rendered audio")}
}
