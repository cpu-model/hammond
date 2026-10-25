package electrical

import "fmt"

// ConnectPedalPanelInputs attaches the three resistor-panel inputs to
// caller-selected pedal buses. The mapping is explicit because the B3
// schematic's panel terminals must not be guessed from drawbar order.
// A mapping is ordered by panel row (R81, R82, R83), using bus indices 0..3.
// This is a PARTIAL circuit assembly: the fourth pedal bus can feed the
// separate R80/L20 path visible on the source schematic. An unmapped bus
// is not necessarily silent in the complete B3 pedal circuit.
// Source identifies the evidence for this particular wiring decision.
func ConnectPedalPanelInputs(panel PedalCircuitTopology, buses [3]int, source string) (PedalCircuitTopology,error) {
 if source=="" {return PedalCircuitTopology{},fmt.Errorf("pedal panel bus mapping requires evidence source")}
 if err:=panel.Validate();err!=nil {return PedalCircuitTopology{},err}
 result:=PedalCircuitTopology{
  Wires:append([]TerminalEdge(nil),panel.Wires...),
  Components:append([]PedalComponentPlacement(nil),panel.Components...),
 }
 for row,bus:=range buses {
  if bus<0||bus>3 {return PedalCircuitTopology{},fmt.Errorf("panel row %d: invalid pedal bus %d",row+1,bus)}
  result.Wires=append(result.Wires,TerminalEdge{
   From:fmt.Sprintf("panel-row-%d-in",row+1),
   To:fmt.Sprintf("pedal-bus-%d",bus),
   Source:source,
  })
 }
 if err:=result.Validate();err!=nil {return PedalCircuitTopology{},err}
 return result,nil
}
