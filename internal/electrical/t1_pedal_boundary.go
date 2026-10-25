package electrical

import "fmt"

// PedalTransformerInputs records the pedal contribution entering T1.
// This is an electrical interface boundary, not a pedal generator,
// drawbar circuit, or transformer mixing model.
type PedalTransformerInputs struct {
 Connected bool
 Voltage float64
 Source Parameter
}

func (p PedalTransformerInputs) Validate() error {
 if !p.Connected { return nil }
 if !finite(p.Voltage) { return fmt.Errorf("pedal voltage must be finite") }
 if err:=p.Source.Validate();err!=nil { return fmt.Errorf("pedal source: %w",err) }
 if p.Source.Unit!="ohm" || !positiveFinite(p.Source.Value) {
  return fmt.Errorf("pedal source resistance must be positive ohms")
 }
 return nil
}

// T1InputBoundary distinguishes the lower manual's nine tap
// connections from the pedal input. No summation or transformer
// transfer is performed.
type T1InputBoundary struct {
 LowerManual TransformerInputs
 Pedals PedalTransformerInputs
}

func NewT1InputBoundary(lower TransformerInputs,pedals PedalTransformerInputs) (T1InputBoundary,error) {
 if lower.Transformer!=TransformerT1 {
  return T1InputBoundary{},fmt.Errorf("T1 requires lower manual transformer inputs")
 }
 if err:=pedals.Validate();err!=nil { return T1InputBoundary{},err }
 return T1InputBoundary{LowerManual:lower,Pedals:pedals},nil
}
