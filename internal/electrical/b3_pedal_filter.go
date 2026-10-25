package electrical

// B3PedalFilterTopology captures only the two components identified
// in the separate pedal-filter assembly: R80 and reactor L20.
// The diagram connects the three resistor-panel rows to successive
// taps of the R80/L20 branch. Pedal busbar connections remain unresolved.
//
// Source: Hammond B3/C3 Service Manual, pedal switch schematic,
// https://synthmania.com/wp-content/uploads/2020/05/Hammond-Organ-B3-C3-Service-Manual.pdf
func B3PedalFilterTopology() PedalCircuitTopology {
 const source="B3/C3 pedal switch schematic, pedal filter assembly, PDF page 72"
 return PedalCircuitTopology{Wires:[]TerminalEdge{{From:"filter-R80-b",To:"filter-L20-a",Source:source}},Components:[]PedalComponentPlacement{
  {Designator:"R80",Kind:PedalResistor,TerminalA:"filter-R80-a",TerminalB:"filter-R80-b",Source:source},
  {Designator:"L20",Kind:PedalReactor,TerminalA:"filter-L20-a",TerminalB:"filter-L20-b",Source:source},
 }}
}

// B3PedalAnalogComponents combines the resistor-panel and filter
// component inventories and the three drawn row taps.
// No pedal busbar assignment or component value is inferred.
func B3PedalAnalogComponents() PedalCircuitTopology {
 panel:=B3PedalResistorPanelTopology()
 filter:=B3PedalFilterTopology()
 panel.Components=append(panel.Components,filter.Components...)
 panel.Wires=append(panel.Wires,filter.Wires...)
 const source="B3/C3 pedal switch schematic, PDF page 72, resistor row taps"
 panel.Wires=append(panel.Wires,
  TerminalEdge{From:"filter-R80-a",To:"panel-row-1-in",Source:source},
  TerminalEdge{From:"filter-R80-b",To:"panel-row-2-in",Source:source},
  TerminalEdge{From:"filter-L20-b",To:"panel-row-3-in",Source:source},
 )
 return panel
}
