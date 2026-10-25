package registration

import (
 "math"
 "testing"
)

// High-frequency energy of the switching modulation is measured after
// subtracting the uninterrupted oscillator. A short DFT is sufficient
// here and avoids external FFT dependencies.
func TestNinthDrawbarSwitchingHighFrequencyEnergy(t *testing.T) {
 const rate=48000.0
 const n=1024
 const switchAt=256
 makeResidual:=func(samples int)[]float64 {
  var route Routing
  route[0][8]=1
  m,err:=NewTimedManual(route,ContactTiming{})
  if err!=nil{t.Fatal(err)}
  _=m.SetDrawbar(8,8)
  _=m.Key(0,true)
  m.SetNinthDrawbarTransition(samples)
  frame:=make([]float64,91)
  residual:=make([]float64,n)
  for i:=0;i<n;i++ {
   frame[0]=math.Sin(2*math.Pi*440*float64(i)/rate)
   if i==switchAt {m.SetPercussionDrawbarCut(true)}
   residual[i]=m.Mix(frame)-frame[0]
  }
  return residual
 }
 // Subtract the steady-state post-switch signal, leaving only the
 // finite switching transient (rather than the sustained removed tone).
 transient:=func(residual []float64)[]float64 {
  out:=make([]float64,len(residual))
  for i,v:=range residual {
   if i<switchAt {out[i]=v} else {
    oscillator:=math.Sin(2*math.Pi*440*float64(i)/rate)
    out[i]=v+oscillator
   }
  }
  return out
 }
 highEnergy:=func(x []float64)float64 {
  var energy float64
  // Bins 128..512 correspond to 6..24 kHz at 48 kHz.
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
 immediate:=highEnergy(transient(makeResidual(1)))
 smooth:=highEnergy(transient(makeResidual(144)))
 if !(smooth<immediate) {
  t.Fatalf("high-frequency transition energy not reduced: immediate=%g smooth=%g",immediate,smooth)
 }
 t.Logf("switching energy above 6kHz: immediate=%g, 3ms=%g, reduction=%.2f dB",immediate,smooth,10*math.Log10(immediate/smooth))
}
