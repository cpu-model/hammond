package electrical

import "fmt"

// TransformerEvidence records a traceable source for each category of
// physical transformer data. A connection identity alone is not enough
// to authorize an electrical transfer calculation.
type TransformerEvidence struct {
 ID TransformerID
 WindingTopologySource string
 TapTopologySource string
 ImpedanceSource string
 TransferSource string
}

// ReadyForCircuitModel reports missing evidence needed before claiming
// a physical T1/T2 model. It does not certify source correctness.
func (e TransformerEvidence) ReadyForCircuitModel() error {
 if e.ID!=TransformerT1&&e.ID!=TransformerT2 {
  return fmt.Errorf("invalid transformer identity %q",e.ID)
 }
 if e.WindingTopologySource=="" { return fmt.Errorf("%s: winding topology source missing",e.ID) }
 if e.TapTopologySource=="" { return fmt.Errorf("%s: tap topology source missing",e.ID) }
 if e.ImpedanceSource=="" { return fmt.Errorf("%s: impedance source missing",e.ID) }
 if e.TransferSource=="" { return fmt.Errorf("%s: transfer source missing",e.ID) }
 return nil
}

// TransformerEvidencePair requires separate evidence for T1 and T2.
type TransformerEvidencePair struct {
 UpperT2 TransformerEvidence
 LowerT1 TransformerEvidence
}

func (p TransformerEvidencePair) ReadyForCircuitModel() error {
 if p.UpperT2.ID!=TransformerT2||p.LowerT1.ID!=TransformerT1 {
  return fmt.Errorf("transformer evidence identities must be upper T2 and lower T1")
 }
 if err:=p.UpperT2.ReadyForCircuitModel();err!=nil { return err }
 return p.LowerT1.ReadyForCircuitModel()
}
