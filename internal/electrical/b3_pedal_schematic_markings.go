package electrical

// B3PedalSchematicMarking preserves the numeric marking printed beside
// a component on the B3/C3 pedal-switch schematic. A marking is not a
// validated SI value: units, scale, tolerances and revisions must be
// established before supplying PedalComponentValue to the RL solver.
type B3PedalSchematicMarking struct {
 Designator string
 PrintedNumber string
 Source string
}

// B3PedalPrintedMarkings transcribes the component labels from the
// lower-left PEDAL SWITCH detail on PDF page 72 (zero-based page 71).
// The (1000) annotation appears beside the optional heater in the motor
// section, not as an L20 inductance specification. L20 has no verified
// numerical inductance here.
func B3PedalPrintedMarkings() []B3PedalSchematicMarking {
 const source="https://synthmania.com/wp-content/uploads/2020/05/Hammond-Organ-B3-C3-Service-Manual.pdf#page=72"
 return []B3PedalSchematicMarking{
  {Designator:"R80",PrintedNumber:"10",Source:source},
  {Designator:"R81",PrintedNumber:"10",Source:source},
  {Designator:"R82",PrintedNumber:"47",Source:source},
  {Designator:"R83",PrintedNumber:"470",Source:source},
  {Designator:"R84",PrintedNumber:"5",Source:source},
  {Designator:"R85",PrintedNumber:"5",Source:source},
  {Designator:"R86",PrintedNumber:"20",Source:source},
  {Designator:"L20",PrintedNumber:"",Source:source},
 }
}
