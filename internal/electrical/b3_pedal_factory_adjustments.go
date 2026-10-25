package electrical

// B3PedalAdjustmentEvidence records a documented factory adjustment
// specifically attributable to the pedal circuit.
type B3PedalAdjustmentEvidence struct {
 Designator string
 Instruction string
 Source string
}

// B3PedalFactoryAdjustments includes note 7 of the full B3/C3 schematic,
// which explicitly states that R81 is selected at test to balance channels.
// A separate preamplifier annotation concerns R10.
// Source: B3/C3 schematic, PDF page 72:
// https://synthmania.com/wp-content/uploads/2020/05/Hammond-Organ-B3-C3-Service-Manual.pdf#page=72
func B3PedalFactoryAdjustments() []B3PedalAdjustmentEvidence {
 return []B3PedalAdjustmentEvidence{{Designator:"R81",Instruction:"SELECTED AT TEST TO BALANCE THE CHANNELS",Source:"B3/C3 schematic PDF page 72, note 7"}}
}
