package engine

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
)

func TestTimedPercussionIndependentOfSampleRate(t *testing.T) {
 route:=registration.B3ManualRouting()
 for _,rate:=range []float64{44100,48000,96000} {
  e,err:=NewTimed(rate,route,route,registration.ContactTiming{},registration.ContactTiming{})
  if err!=nil{t.Fatal(err)}
  if err:=e.SetTimedPercussion(0,3,0.8,0.2);err!=nil{t.Fatal(err)}
  p:=&e.percussion[0]
  p.Trigger()
  count:=int(rate*0.2)
  for i:=0;i<count;i++ {p.NextHarmonic(1)}
  want:=0.8/math.E
  if math.Abs(p.envelope-want)>1e-10 {t.Fatalf("rate %.0f: envelope %v want %v",rate,p.envelope,want)}
 }
}

func TestTimedPercussionRejectsInvalidDuration(t *testing.T) {
 route:=registration.B3ManualRouting()
 e,err:=NewTimed(48000,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil{t.Fatal(err)}
 for _,duration:=range []float64{0,-1,math.NaN(),math.Inf(1)} {
  if err:=e.SetTimedPercussion(0,2,0.5,duration);err==nil{t.Fatalf("accepted %v",duration)}
 }
}
