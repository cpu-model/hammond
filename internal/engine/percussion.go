package engine

import (
 "errors"
 "math"
)

// Percussion is an experimental decaying tap of a shared generator wheel.
// The oscillator phase belongs to MechanicalBank, not this envelope.
type Percussion struct {
 wheel int
 level float64
 decay float64
 envelope float64
}

func (p *Percussion) Configure(wheel int,level,decay float64) error {
 if wheel<1||wheel>91||math.IsNaN(level)||math.IsInf(level,0)||level<0||level>1||math.IsNaN(decay)||math.IsInf(decay,0)||decay<0||decay>=1 {
  return errors.New("invalid percussion wheel, level or decay")
 }
 p.wheel=wheel
 p.level=level
 p.decay=decay
 p.envelope=0
 return nil
}

func (p *Percussion) Trigger() {
 if p.wheel!=0 {p.envelope=p.level}
}

func (p *Percussion) Next(frame []float64) float64 {
 if p.wheel==0 {return 0}
 sample:=p.envelope*frame[p.wheel-1]
 p.envelope*=p.decay
 return sample
}
