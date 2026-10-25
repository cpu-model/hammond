package electrical

import (
 "math"
 "testing"
)

func TestLowpassStepResponse(t *testing.T) {
 const rate=48000.0
 const cutoff=1000.0
 f,err:=NewLowpass(rate,cutoff)
 if err!=nil { t.Fatal(err) }
 a:=math.Exp(-2*math.Pi*cutoff/rate)
 for n:=1;n<=200;n++ {
  got:=f.Next(1)
  want:=1-math.Pow(a,float64(n))
  if math.Abs(got-want)>1e-12 { t.Fatalf("sample %d: %.15f != %.15f",n,got,want) }
 }
 f.Reset()
 if f.Next(0)!=0 { t.Fatal("reset failed") }
}

func TestLowpassRejectsInvalidParameters(t *testing.T) {
 for _,p:=range [][2]float64{{0,1000},{48000,0},{48000,24000},{math.NaN(),1000},{48000,math.Inf(1)}} {
  if _,err:=NewLowpass(p[0],p[1]);err==nil { t.Fatalf("accepted invalid parameters %v",p) }
 }
}

func TestLowpassNoAllocations(t *testing.T) {
 f,_:=NewLowpass(48000,1000)
 if n:=testing.AllocsPerRun(100,func(){f.Next(0.5)});n!=0 { t.Fatalf("allocations: %v",n) }
}
