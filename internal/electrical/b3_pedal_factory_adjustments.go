package electrical

// B3PedalAdjustmentEvidence records a documented factory adjustment
// specifically attributable to the pedal circuit.
type B3PedalAdjustmentEvidence struct {
 Designator string
 Instruction string
 Source string
}

// B3PedalFactoryAdjustments currently has no verified pedal-specific
// adjustments. The schematic's "R10 - SELECT AT TEST TO BALANCE CHANNELS"
// refers to preamplifier R10, NOT pedal resistor R81.
// Source: B3/C3 schematic, PDF page 72:
// https://synthmania.com/wp-content/uploads/2020/05/Hammond-Organ-B3-C3-Service-Manual.pdf#page=72
func B3PedalFactoryAdjustments() []B3PedalAdjustmentEvidence {
 return nil
}
