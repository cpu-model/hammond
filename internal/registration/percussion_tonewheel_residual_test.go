package registration

import (
 "math"
 "testing"
)

// Compare a switching drawbar to the identical uninterrupted oscillator.
// The residual removes ordinary tonewheel movement from the measurement.
func TestTonewheelSwitchingResidual(t *testing.T) {
 const sampleRate=48000.0
 const frequency=440.0
 const switchAt=240
 measure:=func(rampSamples int)float64 {
  var route Routing
  route[0][8]=1
  m,err:=NewTimedManual(route,ContactTiming{})
  if err!=nil{t.Fatal(err)}
  if err:=m.SetDrawbar(8,8);err!=nil{t.Fatal(err)}
  if err:=m.Key(0,true);err!=nil{t.Fatal(err)}
  m.SetNinthDrawbarTransition(rampSamples)
  frame:=make([]float64,91)
  var peak float64
  var previousResidual float64
  var previousOscillator float64
  for i:=0;i<switchAt+600;i++ {
   frame[0]=math.Sin(2*math.Pi*frequency*float64(i)/sampleRate)
   if i==switchAt {m.SetPercussionDrawbarCut(true)}
   actual:=m.Mix(frame)
   // A constant gain of one would produce exactly frame[0].
   residual:=actual-frame[0]
   if i>switchAt && i<switchAt+rampSamples+1 {
    // Remove the derivative expected from the changing oscillator
    // at the previous sample gain. The remainder is the gain change.
    previousGain:=1.0
    if previousOscillator!=0 {previousGain=1+previousResidual/previousOscillator}
    step:=math.Abs((residual-previousResidual)-(previousGain-1)*(frame[0]-previousOscillator))
    if step>peak {peak=step}
   }
   previousResidual=residual
   previousOscillator=frame[0]
  }
  return peak
 }
 immediate:=measure(1)
 smooth:=measure(144)
 if !(smooth<immediate) {
  t.Fatalf("3ms transition did not reduce residual peak: immediate=%v smooth=%v",immediate,smooth)
 }
 t.Logf("440 Hz residual peak derivative: immediate=%g, 3ms=%g, ratio=%g",immediate,smooth,immediate/smooth)
}
