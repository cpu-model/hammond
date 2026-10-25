package electrical

// B3PedalResistorPanelTopology represents a provisional interpretation of three resistor pairs
// shown in the lower-left resistor panel of the 1995 B3/C3
// schematic (PDF page 72). It deliberately does not assert where the
// external busbar, drawbar, filter, or ground wires terminate.
//
// The pairs are R81--R84, R82--R85, and R83--R86.
// The three right-hand ends join the same rail in the interpreted drawing.
// Their left-hand inputs remain distinct; no other junctions are assumed.
// This partial reconstruction is not yet a verified factory netlist.
//
// Source: https://synthmania.com/wp-content/uploads/2020/05/Hammond-Organ-B3-C3-Service-Manual.pdf
// page 72, lower-left PEDAL SWITCH schematic.
func B3PedalResistorPanelTopology() PedalCircuitTopology {
 const source="B3/C3 schematic, December 11 1995, PDF page 72, pedal resistor panel"
 return PedalCircuitTopology{Components:[]PedalComponentPlacement{
  {Designator:"R81",Kind:PedalResistor,TerminalA:"panel-row-1-in",TerminalB:"panel-row-1-mid",Source:source},
  {Designator:"R84",Kind:PedalResistor,TerminalA:"panel-row-1-mid",TerminalB:"panel-common-out",Source:source},
  {Designator:"R82",Kind:PedalResistor,TerminalA:"panel-row-2-in",TerminalB:"panel-row-2-mid",Source:source},
  {Designator:"R85",Kind:PedalResistor,TerminalA:"panel-row-2-mid",TerminalB:"panel-common-out",Source:source},
  {Designator:"R83",Kind:PedalResistor,TerminalA:"panel-row-3-in",TerminalB:"panel-row-3-mid",Source:source},
  {Designator:"R86",Kind:PedalResistor,TerminalA:"panel-row-3-mid",TerminalB:"panel-common-out",Source:source},
 }}
}
