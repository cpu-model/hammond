package registration

import (
 "math"
 "testing"
)

// Measure the additional discontinuity from ninth-drawbar switching
// against the same wheel waveform without a switch.
func TestNinthDrawbarTransitionReducesSwitchingResidual(t *testing.T) {
 measure:=func(samples int)float64 {
  var route Routing
  route[0][8]=1
  m,err:=NewTimedManual(route,ContactTiming{})
  if err!=nil{t.Fatal(err)}
  if err:=m.SetDrawbar(8,8);err!=nil{t.Fatal(err)}
  if err:=m.Key(0,true);err!=nil{t.Fatal(err)}
  m.SetNinthDrawbarTransition(samples)
  frame:=make([]float64,91)
  // Constant generator isolates the discontinuity caused by switching.
  frame[0]=1
  m.Mix(frame)
  m.SetPercussionDrawbarCut(true)
  previous:=1.0
  peak:=0.0
  for i:=0;i<samples+2;i++ {
   current:=m.Mix(frame)
   jump:=math.Abs(current-previous)
   if jump>peak {peak=jump}
   previous=current
  }
  return peak
 }
 immediate:=measure(1)
 threeMS:=measure(144)
 tenMS:=measure(480)
 if math.Abs(immediate-1)>1e-12 {t.Fatalf("immediate jump %v",immediate)}
 if math.Abs(threeMS-1.0/144)>1e-12 {t.Fatalf("3ms jump %v",threeMS)}
 if math.Abs(tenMS-1.0/480)>1e-12 {t.Fatalf("10ms jump %v",tenMS)}
 if !(tenMS<threeMS&&threeMS<immediate) {t.Fatalf("unexpected jump order %v %v %v",immediate,threeMS,tenMS)}
}
