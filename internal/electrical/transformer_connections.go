package electrical

import "fmt"

// TransformerID distinguishes the two B3 manual transformer identities.
// Identity does not establish a winding model or any circuit values.
type TransformerID string

const (
 TransformerT1 TransformerID = "T1"
 TransformerT2 TransformerID = "T2"
)

// ManualTransformerConnection preserves a manual's nine independent
// conductor voltages and its selected tap positions. No voltages are
// transformed, loaded or summed.
type ManualTransformerConnection struct {
 ID TransformerID
 Selection DrawbarTapSelection
}

func (c ManualTransformerConnection) Validate() error {
 if c.ID!=TransformerT1&&c.ID!=TransformerT2 {
  return fmt.Errorf("unknown manual transformer %q",c.ID)
 }
 return c.Selection.Validate()
}

// TransformerInputs is a diagnostic wiring record, not an electrical
// transformer output. A zero-position drawbar is marked disconnected;
// connected entries carry raw bus voltages and selected tap indices.
type TransformerInputs struct {
 Transformer TransformerID
 Connected [9]bool
 Tap [9]uint8
 Voltage [9]float64
}

func (c ManualTransformerConnection) Inputs(bus [9]float64) (TransformerInputs,error) {
 if err:=c.Validate();err!=nil { return TransformerInputs{},err }
 out:=TransformerInputs{Transformer:c.ID}
 for i,tap:=range c.Selection {
  if tap==0 { continue }
  out.Connected[i]=true
  out.Tap[i]=tap
  out.Voltage[i]=bus[i]
 }
 return out,nil
}

// ManualTransformerPair records T2 for upper and T1 for lower manual.
type ManualTransformerPair struct {
 Upper ManualTransformerConnection
 Lower ManualTransformerConnection
}

func (p ManualTransformerPair) Validate() error {
 if p.Upper.ID!=TransformerT2||p.Lower.ID!=TransformerT1 {
  return fmt.Errorf("upper manual requires T2 and lower manual requires T1")
 }
 if err:=p.Upper.Validate();err!=nil { return err }
 return p.Lower.Validate()
}

func (p ManualTransformerPair) Inputs(upper,lower [9]float64) (TransformerInputs,TransformerInputs,error) {
 if err:=p.Validate();err!=nil { return TransformerInputs{},TransformerInputs{},err }
 u,_:=p.Upper.Inputs(upper)
 l,_:=p.Lower.Inputs(lower)
 return u,l,nil
}
