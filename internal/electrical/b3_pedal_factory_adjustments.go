package electrical

// B3PedalAdjustmentEvidence records a factory adjustment instruction,
// not a fixed component value. The B3/C3 schematic note 6 identifies
// R81 as selected at test to balance the channels.
type B3PedalAdjustmentEvidence struct {
 Designator string
 Instruction string
 Source string
}

func B3PedalFactoryAdjustments() []B3PedalAdjustmentEvidence {
 return []B3PedalAdjustmentEvidence{{
  Designator:"R81",
  Instruction:"SELECTED AT TEST TO BALANCE THE CHANNELS",
  Source:"https://synthmania.com/wp-content/uploads/2020/05/Hammond-Organ-B3-C3-Service-Manual.pdf#page=72, note 6",
 }}
}
