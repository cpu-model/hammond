package engine

import (
 "errors"
 "math"
)

// PercussionControls models the four familiar Hammond percussion choices.
// Time constants and amplitudes are experimental calibration values.
type PercussionControls struct {
 Harmonic int
 Fast bool
 Soft bool
 FastSeconds float64
 SlowSeconds float64
 NormalLevel float64
 SoftLevel float64
}

func DefaultPercussionControls() PercussionControls {
 return PercussionControls{
  Harmonic:3,
  Fast:true,
  Soft:false,
  FastSeconds:0.18,
  SlowSeconds:0.65,
  NormalLevel:0.5,
  SoftLevel:0.25,
 }
}

func (e *TimedEngine) SetPercussionControls(manual int,c PercussionControls) error {
 if c.Harmonic!=2&&c.Harmonic!=3 {return errors.New("percussion harmonic must be 2 or 3")}
 for _,value:=range []float64{c.FastSeconds,c.SlowSeconds,c.NormalLevel,c.SoftLevel} {
  if math.IsNaN(value)||math.IsInf(value,0) {return errors.New("percussion control values must be finite")}
 }
 if c.FastSeconds<=0||c.SlowSeconds<=0||c.NormalLevel<0||c.NormalLevel>1||c.SoftLevel<0||c.SoftLevel>1 {
  return errors.New("invalid percussion time or level")
 }
 seconds:=c.SlowSeconds
 if c.Fast {seconds=c.FastSeconds}
 level:=c.NormalLevel
 if c.Soft {level=c.SoftLevel}
 return e.SetTimedPercussion(manual,c.Harmonic,level,seconds)
}
