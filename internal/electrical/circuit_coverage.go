package electrical

import "errors"

// CircuitCoverage records whether a transformer circuit representation
// includes the physical signal contributors shown in the B3/C3
// schematic. It is a completeness check, not a calibrated model.
type CircuitCoverage struct {
 UpperManual bool
 LowerManual bool
 Pedals bool
}

// ValidateFullTransformerCoverage refuses to treat a manual-only
// experiment as a complete T1/T2 electrical circuit. T1 serves
// lower manual AND pedals; T2 serves upper manual.
func (c CircuitCoverage) ValidateFullTransformerCoverage() error {
 if !c.UpperManual { return errors.New("T2 upper-manual circuit missing") }
 if !c.LowerManual { return errors.New("T1 lower-manual circuit missing") }
 if !c.Pedals { return errors.New("T1 pedal circuit missing") }
 return nil
}

// ExperimentalManualOnlyCoverage documents the existing engine scope.
// Its ValidateFullTransformerCoverage result must be an error.
func ExperimentalManualOnlyCoverage() CircuitCoverage {
 return CircuitCoverage{UpperManual:true,LowerManual:true,Pedals:false}
}
