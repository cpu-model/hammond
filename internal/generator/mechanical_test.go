package generator

import (
 "math"
 "testing"
)

func TestMechanicalBank(t *testing.T) {
 b,err:=NewMechanicalBank(48000)
 if err!=nil { t.Fatal(err) }
 frame:=make([]float64,91)
 for n:=0;n<48000;n++ {
  b.Next(frame)
  // Frequency number 46: 440 Hz, derived from speed group A.
  want:=math.Sin(2*math.Pi*440*float64(n)/48000)
  if math.Abs(frame[45]-want)>1e-8 {
   t.Fatalf("sample %d: got %.12f want %.12f",n,frame[45],want)
  }
 }
}

func TestMechanicalBankContinuousBlocks(t *testing.T) {
 a,_:=NewMechanicalBank(48000)
 b,_:=NewMechanicalBank(48000)
 x,y:=make([]float64,91),make([]float64,91)
 for i:=0;i<1000;i++ { a.Next(x) }
 for i:=0;i<400;i++ { b.Next(y) }
 for i:=0;i<600;i++ { b.Next(y) }
 a.Next(x);b.Next(y)
 for i:=range x { if x[i]!=y[i] { t.Fatalf("wheel %d differs across block boundaries",i+1) } }
}

func TestMechanicalBankNoAllocations(t *testing.T) {
 b,_:=NewMechanicalBank(48000)
 frame:=make([]float64,91)
 if n:=testing.AllocsPerRun(100,func(){ b.Next(frame) });n!=0 {
  t.Fatalf("render allocations: %v",n)
 }
}

func TestMechanicalBankRejectLowRate(t *testing.T) {
 if _,err:=NewMechanicalBank(8000);err==nil { t.Fatal("expected Nyquist error") }
}
