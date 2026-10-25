package registration

import (
 "math"
 "testing"
)

// A deterministic frequency sweep compares the emitted gated tone
// across multiple wheel frequencies and fade lengths.
func TestNinthDrawbarSpectralSweep(t *testing.T) {
 const rate=48000.0
 const n=1024
 const switchAt=256
 frequencies:=[]float64{110,220,440,880,1760}
 fades:=[]int{1,48,144,480}
 highEnergy:=func(x []float64)float64 {
  var energy float64
  for k:=128;k<=n/2;k++ {
   var re,im float64
   for i,v:=range x {
    phase:=2*math.Pi*float64(k*i)/n
    re+=v*math.Cos(phase)
    im-=v*math.Sin(phase)
   }
   energy+=re*re+im*im
  }
  return energy
 }
 render:=func(frequency float64,fade int)[]float64 {
  var route Routing
  route[0][8]=1
  m,err:=NewTimedManual(route,ContactTiming{})
  if err!=nil{t.Fatal(err)}
  if err:=m.SetDrawbar(8,8);err!=nil{t.Fatal(err)}
  if err:=m.Key(0,true);err!=nil{t.Fatal(err)}
  m.SetNinthDrawbarTransition(fade)
  frame:=make([]float64,91)
  signal:=make([]float64,n)
  for i:=0;i<n;i++ {
   frame[0]=math.Sin(2*math.Pi*frequency*float64(i)/rate)
   if i==switchAt {m.SetPercussionDrawbarCut(true)}
   signal[i]=m.Mix(frame)
  }
  return signal
 }
 for _,frequency:=range frequencies {
  baseline:=highEnergy(render(frequency,1))
  if baseline<=0 {t.Fatalf("zero baseline at %g Hz",frequency)}
  for _,fade:=range fades[1:] {
   measured:=highEnergy(render(frequency,fade))
   reduction:=10*math.Log10(baseline/measured)
   t.Logf("wheel %.0f Hz, fade %.1f ms: high-band reduction %.2f dB",frequency,float64(fade)/rate*1000,reduction)
   if math.IsNaN(reduction)||math.IsInf(reduction,0) {t.Fatalf("invalid energy ratio at %g Hz, %d samples",frequency,fade)}
   if fade==144 && measured>=baseline {t.Errorf("3ms fade increases high-band energy at %.0f Hz: immediate=%g fade=%g",frequency,baseline,measured)}
  }
 }
}
