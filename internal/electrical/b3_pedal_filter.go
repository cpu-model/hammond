package electrical

// B3PedalFilterTopology inventories R80 and L20 from the pedal-switch
// detail. This name does not imply a separate physical filter assembly:
// the schematic places these designators alongside the pedal resistor
// network. No junction between their terminals is asserted here.
// Source: Hammond B3/C3 Service Manual, full schematic, PDF page 72.
func B3PedalFilterTopology() PedalCircuitTopology {
 const source="B3/C3 schematic, pedal switch, PDF page 72"
 return PedalCircuitTopology{Components:[]PedalComponentPlacement{
  {Designator:"R80",Kind:PedalResistor,TerminalA:"filter-R80-a",TerminalB:"filter-R80-b",Source:source},
  {Designator:"L20",Kind:PedalReactor,TerminalA:"filter-L20-a",TerminalB:"filter-L20-b",Source:source},
 }}
}

// B3PedalAnalogComponents inventories the pedal components without
// introducing unverified electrical connections between assemblies.
func B3PedalAnalogComponents() PedalCircuitTopology {
 panel:=B3PedalResistorPanelTopology()
 filter:=B3PedalFilterTopology()
 panel.Components=append(panel.Components,filter.Components...)
 return panel
}
