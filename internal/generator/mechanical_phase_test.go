package generator

import (
 "math"
 "testing"
)

func TestWheelPhaseOffsetIndependent(t *testing.T) {
 a,err:=NewMechanicalBank(48000);if err!=nil{t.Fatal(err)}
 b,err:=NewMechanicalBank(48000);if err!=nil{t.Fatal(err)}
 if err:=b.SetWheelPhaseOffset(1,0.25);err!=nil{t.Fatal(err)}
 var x,y [91]float64
 for sample:=0;sample<64;sample++ {
  a.Next(x[:]);b.Next(y[:])
  if math.Abs(x[0]-y[0])<0.1 {t.Fatalf("offset not applied at sample %d",sample)}
  for wheel:=1;wheel<91;wheel++ {
   if x[wheel]!=y[wheel] {t.Fatalf("wheel %d changed at sample %d",wheel+1,sample)}
  }
 }
}

func TestWheelPhaseOffsetValidation(t *testing.T) {
 b,err:=NewMechanicalBank(48000);if err!=nil{t.Fatal(err)}
 for _,v:=range []float64{-0.1,1,math.NaN(),math.Inf(1)} {
  if err:=b.SetWheelPhaseOffset(1,v);err==nil{t.Fatalf("accepted %v",v)}
 }
 for _,wheel:=range []int{0,92} {
  if err:=b.SetWheelPhaseOffset(wheel,0.25);err==nil{t.Fatalf("accepted wheel %d",wheel)}
 }
}
