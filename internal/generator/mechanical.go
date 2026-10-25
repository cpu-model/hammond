package generator

import (
 "errors"
 "math"
)

// MechanicalBank advances 12 shared speed-group phases. It models nominal
// gearing, not individual shaft offsets, pickup shapes or electrical filtering.
// Wheel output numbers are 1..91, indexed 0..90 in the returned frame.
type MechanicalBank struct {
 phase [12]float64
 step [12]float64
 harmonic2 [91]float64
 phaseOffset [91]float64
 secondAllowed [91]bool
}

var driveTeeth = [12]float64{85,71,67,105,103,84,74,98,96,88,67,108}
var drivenTeeth = [12]float64{104,82,73,108,100,77,64,80,74,64,46,70}

// NewMechanicalBank uses the US 60 Hz generator's nominal 1200 RPM drive.
// This reference is fixed independently of the audio sample rate.
func NewMechanicalBank(sampleRate float64) (*MechanicalBank, error) {
 if sampleRate <= 0 || math.IsNaN(sampleRate) || math.IsInf(sampleRate,0) {
  return nil, errors.New("invalid sample rate")
 }
 b := &MechanicalBank{}
 for n:=0;n<91;n++ {
  group:=n%12
  teeth:=2 << (n/12)
  if n>=84 {group=n-84+5;teeth=192}
  fundamental:=20*(driveTeeth[group]/drivenTeeth[group])*float64(teeth)
  b.secondAllowed[n]=2*fundamental<sampleRate/2
 }
 for i := range b.step {
  b.step[i] = (1200.0/60.0)*(driveTeeth[i]/drivenTeeth[i])/sampleRate
 }
 // All modeled outputs must be below Nyquist.
 if sampleRate <= 2*20*192*108.0/70.0 {
  return nil, errors.New("sample rate below highest generator signal Nyquist")
 }
 return b,nil
}

// SetSecondHarmonic enables a bounded, uncalibrated pickup-shape experiment.
// Second harmonics above Nyquist are suppressed rather than aliased.
// Zero preserves the original ideal sine output; no factory spectrum is implied.
func (b *MechanicalBank) SetSecondHarmonic(level float64) error {
 if math.IsNaN(level)||math.IsInf(level,0)||level<0||level>1 {
  return errors.New("second harmonic level must be finite and between 0 and 1")
 }
 for i:=range b.harmonic2 {b.harmonic2[i]=level}
 return nil
}

// SetWheelSecondHarmonic sets an individual wheel (1..91) without
// changing other wheels. Values are experimental, not measured B3 data.
func (b *MechanicalBank) SetWheelSecondHarmonic(wheel int,level float64) error {
 if wheel<1||wheel>91 {return errors.New("wheel must be between 1 and 91")}
 if math.IsNaN(level)||math.IsInf(level,0)||level<0||level>1 {
  return errors.New("second harmonic level must be finite and between 0 and 1")
 }
 b.harmonic2[wheel-1]=level
 return nil
}

// SetWheelPhaseOffset changes the initial phase of a wheel in cycles [0,1).
// The shared gearing and frequency remain unchanged. Values are experimental.
func (b *MechanicalBank) SetWheelPhaseOffset(wheel int,cycles float64) error {
 if wheel<1||wheel>91 {return errors.New("wheel must be between 1 and 91")}
 if math.IsNaN(cycles)||math.IsInf(cycles,0)||cycles<0||cycles>=1 {
  return errors.New("wheel phase offset must be finite and in [0,1)")
 }
 b.phaseOffset[wheel-1]=cycles
 return nil
}

// Next writes 91 phase-continuous signals. Every wheel advances
// even when no keys are held. The 192-tooth wheels 85..91 use gear groups 5..11.
func (b *MechanicalBank) Next(dst []float64) {
 if len(dst)<91 { panic("generator: destination too short") }
 for n:=0;n<91;n++ {
  group := n%12
  teeth := 2 << (n/12)
  if n>=84 { group=n-84+5; teeth=192 }
  cycles := math.Mod(float64(teeth)*b.phase[group]+b.phaseOffset[n],1)
  angle:=2*math.Pi*cycles
  level:=b.harmonic2[n]
  if !b.secondAllowed[n] {level=0}
  dst[n] = (math.Sin(angle)+level*math.Sin(2*angle))/(1+level)
 }
 for i:=range b.phase {
  p:=b.phase[i]+b.step[i]
  if p>=1 { p-=1 }
  b.phase[i]=p
 }
}
