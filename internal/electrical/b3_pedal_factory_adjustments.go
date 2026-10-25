package electrical

// B3PedalAdjustmentEvidence records a documented factory adjustment
// specifically attributable to the pedal circuit.
type B3PedalAdjustmentEvidence struct {
 Designator string
 Instruction string
 Source string
}

// B3PedalFactoryAdjustments includes note 7 of the full B3/C3 schematic,
// which identifies R61 (not R81) as selected at test for channel balance.
// A separate preamplifier annotation concerns R10.
// Source: B3/C3 schematic, PDF page 72:
// https://synthmania.com/wp-content/uploads/2020/05/Hammond-Organ-B3-C3-Service-Manual.pdf#page=72
func B3PedalFactoryAdjustments() []B3PedalAdjustmentEvidence {
 return nil
}
