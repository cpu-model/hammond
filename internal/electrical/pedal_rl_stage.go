package electrical

import (
 "errors"
 "math"
)

// PedalRLConfig is an experimental series reactor feeding a resistive load.
// It is NOT a transcription of the B3 pedal filter schematic. All component
// values must be supplied explicitly with provenance.
type PedalRLConfig struct {
 InductanceHenry Parameter
 SeriesResistanceOhm Parameter
 LoadResistanceOhm Parameter
}

// PedalRLStage models an ideal voltage source driving a series R-L branch
// into a grounded resistive load. Backward Euler is unconditionally stable
// for positive component values.
type PedalRLStage struct {
 alpha float64
 gain float64
 output float64
}

func NewPedalRLStage(sampleRate float64,c PedalRLConfig)(*PedalRLStage,error) {
 if !positiveFinite(sampleRate) { return nil,errors.New("invalid pedal RL sample rate") }
 for _,p:=range []Parameter{c.InductanceHenry,c.SeriesResistanceOhm,c.LoadResistanceOhm} {
  if err:=p.Validate();err!=nil { return nil,err }
  if !positiveFinite(p.Value) { return nil,errors.New("pedal RL components must be positive") }
 }
 if c.InductanceHenry.Unit!="H"||c.SeriesResistanceOhm.Unit!="ohm"||c.LoadResistanceOhm.Unit!="ohm" {
  return nil,errors.New("pedal RL component units must be H and ohm")
 }
 resistance:=c.SeriesResistanceOhm.Value+c.LoadResistanceOhm.Value
 if !positiveFinite(resistance) { return nil,errors.New("pedal RL resistance overflow") }
 dt:=1/sampleRate
 // For a series R-L circuit, L di/dt + R i = Vin.
 // Discrete current: i[n] = alpha*i[n-1] + (1-alpha)*Vin/R.
 alpha:=1/(1+dt*resistance/c.InductanceHenry.Value)
 gain:=c.LoadResistanceOhm.Value/resistance
 if math.IsNaN(alpha)||math.IsInf(alpha,0)||alpha<0||alpha>=1||!positiveFinite(gain) {
  return nil,errors.New("pedal RL coefficients not representable")
 }
 return &PedalRLStage{alpha:alpha,gain:gain},nil
}

func (s *PedalRLStage) Next(input float64)(float64,error) {
 if !finite(input) { return 0,errors.New("nonfinite pedal RL input") }
 s.output=s.alpha*s.output+(1-s.alpha)*s.gain*input
 return s.output,nil
}

func (s *PedalRLStage) Reset(){s.output=0}
