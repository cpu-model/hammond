package engine

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
)

// Contact delays must not alter the shared tonewheel clock. With a
// single active drawbar, changing timing for an inactive drawbar must
// leave the complete output bit-for-bit identical.
func TestInactiveContactTimingDoesNotChangeAudio(t *testing.T) {
 route:=registration.B3ManualRouting()
 var delayed registration.ContactTiming
 delayed.Make[1]=72
 delayed.Break[1]=48
 a,err:=NewTimed(48000,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil {t.Fatal(err)}
 b,err:=NewTimed(48000,route,route,delayed,registration.ContactTiming{})
 if err!=nil {t.Fatal(err)}
 for _,e:=range []*TimedEngine{a,b} {
  m,err:=e.Manual(0)
  if err!=nil {t.Fatal(err)}
  if err:=m.SetDrawbar(0,8);err!=nil {t.Fatal(err)}
  if err:=m.Key(36,true);err!=nil {t.Fatal(err)}
 }
 for i:=0;i<2048;i++ {
  x,y:=a.Next(),b.Next()
  if math.Float64bits(x)!=math.Float64bits(y) {
   t.Fatalf("inactive contact timing changed sample %d: %g vs %g",i,x,y)
  }
 }
}
