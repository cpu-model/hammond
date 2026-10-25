package engine

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
)

// A delayed second contact must introduce an additional harmonic only
// after its make time. This is a synthetic timing contract, not a
// measurement of B3 contact travel.
func TestTimedManualDelayedSecondContact(t *testing.T) {
 route:=registration.B3ManualRouting()
 var timing registration.ContactTiming
 timing.Make[1]=48
 e,err:=NewTimed(48000,route,route,timing,registration.ContactTiming{})
 if err!=nil {t.Fatal(err)}
 upper,err:=e.Manual(0)
 if err!=nil {t.Fatal(err)}
 for _,d:=range []int{0,1} {
  if err:=upper.SetDrawbar(d,8);err!=nil {t.Fatal(err)}
 }
 if err:=upper.Key(36,true);err!=nil {t.Fatal(err)}
 // Compare against an engine with only the first drawbar active.
 baseline,err:=NewTimed(48000,route,route,timing,registration.ContactTiming{})
 if err!=nil {t.Fatal(err)}
 first,_:=baseline.Manual(0)
 if err:=first.SetDrawbar(0,8);err!=nil {t.Fatal(err)}
 if err:=first.Key(36,true);err!=nil {t.Fatal(err)}
 differenceAfterDelay:=false
 for i:=0;i<512;i++ {
  got,want:=e.Next(),baseline.Next()
  if i<48 && math.Abs(got-want)>1e-10 {
   t.Fatalf("second contact audible too early at sample %d: delta=%g",i,got-want)
  }
  if i>=48 && math.Abs(got-want)>1e-8 {
   differenceAfterDelay=true
  }
 }
 if !differenceAfterDelay {t.Fatal("second contact never became audible after its make delay")}
}
