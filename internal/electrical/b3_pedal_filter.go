package electrical

// B3PedalFilterTopology represents the vertical R80/L20 ladder shown
// immediately to the left of the three horizontal resistor-panel rows.
// The upper, middle and lower junctions are the panel row inputs.
// The four pedal busbar connections remain to be traced.
// Source: B3/C3 full schematic, PDF page 72, PEDAL SWITCH detail.
func B3PedalFilterTopology() PedalCircuitTopology {
 const source="B3/C3 full schematic PDF page 72, PEDAL SWITCH detail"
 return PedalCircuitTopology{Components:[]PedalComponentPlacement{
  {Designator:"R80",Kind:PedalResistor,TerminalA:"panel-row-1-in",TerminalB:"panel-row-2-in",Source:source},
  {Designator:"L20",Kind:PedalReactor,TerminalA:"panel-row-2-in",TerminalB:"panel-row-3-in",Source:source},
 }}
}

// B3PedalAnalogComponents combines the three panel branches with the
// R80/L20 vertical ladder. No busbar wiring or numerical values are inferred.
func B3PedalAnalogComponents() PedalCircuitTopology {
 panel:=B3PedalResistorPanelTopology()
 filter:=B3PedalFilterTopology()
 panel.Components=append(panel.Components,filter.Components...)
 return panel
}
