package electrical

import (
 "errors"
 "math"
)

// Lowpass is a single-pole, sample-by-sample RC-equivalent filter.
// It is a controlled electrical-response building block, not a
// circuit-derived model of any specific Hammond B3 component.
type Lowpass struct {
 a float64
 y float64
}

func NewLowpass(sampleRate, cutoffHz float64) (*Lowpass,error) {
 if math.IsNaN(sampleRate)||math.IsInf(sampleRate,0)||sampleRate<=0 ||
  math.IsNaN(cutoffHz)||math.IsInf(cutoffHz,0)||cutoffHz<=0||cutoffHz>=sampleRate/2 {
  return nil,errors.New("invalid sample rate or cutoff")
 }
 // Exact discretization of an analog first-order RC response.
 a:=math.Exp(-2*math.Pi*cutoffHz/sampleRate)
 return &Lowpass{a:a},nil
}

func (f *Lowpass) Next(x float64) float64 {
 f.y=(1-f.a)*x+f.a*f.y
 return f.y
}

func (f *Lowpass) Reset() { f.y=0 }
