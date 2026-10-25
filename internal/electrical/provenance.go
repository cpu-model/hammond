package electrical

import (
 "errors"
 "math"
)

// Provenance identifies the basis for an electrical model parameter.
type Provenance string

const (
 ProvenanceSchematic Provenance = "schematic"
 ProvenanceMeasurement Provenance = "measurement"
 ProvenanceExperimental Provenance = "experimental"
)

// Parameter is a value accompanied by its evidence and units.
// A source citation is mandatory for schematic and measured values.
type Parameter struct {
 Value float64
 Unit string
 Provenance Provenance
 Source string
}

func (p Parameter) Validate() error {
 if math.IsNaN(p.Value)||math.IsInf(p.Value,0) {
  return errors.New("parameter must be finite")
 }
 if p.Unit=="" { return errors.New("parameter unit required") }
 switch p.Provenance {
 case ProvenanceSchematic,ProvenanceMeasurement:
  if p.Source=="" { return errors.New("documented parameter requires source") }
 case ProvenanceExperimental:
  // Experimental parameters must never masquerade as documented values.
 default:
  return errors.New("unknown parameter provenance")
 }
 return nil
}

// ManualResponseConfig records explicitly uncalibrated or evidenced
// per-manual lowpass parameters. It does not claim to model T1/T2.
type ManualResponseConfig struct {
 UpperCutoffHz Parameter
 LowerCutoffHz Parameter
}

func (c ManualResponseConfig) Validate(sampleRate float64) error {
 for _,p:=range []Parameter{c.UpperCutoffHz,c.LowerCutoffHz} {
  if err:=p.Validate();err!=nil { return err }
  if p.Unit!="Hz"||p.Value<=0||p.Value>=sampleRate/2 {
   return errors.New("cutoff must be in Hz and within (0, Nyquist)")
  }
 }
 if math.IsNaN(sampleRate)||math.IsInf(sampleRate,0)||sampleRate<=0 {
  return errors.New("invalid sample rate")
 }
 return nil
}
