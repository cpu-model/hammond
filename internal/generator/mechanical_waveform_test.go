package generator

import (
 "math"
 "testing"
)

func TestWheelSecondHarmonicIndependent(t *testing.T) {
 a,err:=NewMechanicalBank(48000)
 if err!=nil {t.Fatal(err)}
 b,err:=NewMechanicalBank(48000)
 if err!=nil {t.Fatal(err)}
 if err:=b.SetWheelSecondHarmonic(37,0.35);err!=nil {t.Fatal(err)}
 var x,y [91]float64
 changed:=false
 for sample:=0;sample<128;sample++ {
  a.Next(x[:])
  b.Next(y[:])
  for wheel:=range x {
   if wheel==36 {
    if math.Abs(x[wheel]-y[wheel])>1e-8 {changed=true}
   } else if math.Float64bits(x[wheel])!=math.Float64bits(y[wheel]) {
    t.Fatalf("wheel %d changed at sample %d",wheel+1,sample)
   }
  }
 }
 if !changed {t.Fatal("configured wheel waveform did not change")}
 for _,wheel:=range []int{0,92} {
  if err:=b.SetWheelSecondHarmonic(wheel,0.5);err==nil {t.Fatalf("accepted wheel %d",wheel)}
 }
 for _,level:=range []float64{-0.1,1.1,math.NaN(),math.Inf(1)} {
  if err:=b.SetWheelSecondHarmonic(37,level);err==nil {t.Fatalf("accepted invalid level %g",level)}
 }
}
