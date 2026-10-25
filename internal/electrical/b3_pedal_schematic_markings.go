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
// The nearby (1000) annotation has not been conclusively identified\n// as L20 inductance; do not treat it as a physical value.
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
