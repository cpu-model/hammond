package electrical

import "testing"

func TestPedalComponentTopologyPreservesImpedance(t *testing.T) {
 topology:=PedalCircuitTopology{
  Wires:[]TerminalEdge{
   {From:"bus",To:"R80.a",Source:"synthetic fixture"},
   {From:"R80.b",To:"L20.a",Source:"synthetic fixture"},
   {From:"L20.b",To:"out",Source:"synthetic fixture"},
  },
  Components:[]PedalComponentPlacement{
   {Designator:"R80",Kind:PedalResistor,TerminalA:"R80.a",TerminalB:"R80.b",Source:"synthetic fixture"},
   {Designator:"L20",Kind:PedalReactor,TerminalA:"L20.a",TerminalB:"L20.b",Source:"synthetic fixture"},
  },
 }
 if err:=topology.Validate();err!=nil {t.Fatal(err)}
 connected,err:=(TerminalGraph{Edges:topology.Wires}).Connected("bus","out")
 if err!=nil {t.Fatal(err)}
 if connected {t.Fatal("component impedance incorrectly treated as wire")}
 topology.Wires=append(topology.Wires,TerminalEdge{From:"R80.a",To:"R80.b",Source:"synthetic short"})
 if err:=topology.Validate();err==nil {t.Fatal("accepted short-circuited resistor")}
}

func TestPedalComponentTopologyRejectsDuplicateAndInvalidComponents(t *testing.T) {
 c:=PedalComponentPlacement{Designator:"R80",Kind:PedalResistor,TerminalA:"a",TerminalB:"b",Source:"fixture"}
 topology:=PedalCircuitTopology{Components:[]PedalComponentPlacement{c,c}}
 if err:=topology.Validate();err==nil {t.Fatal("accepted duplicate component")}
 topology.Components=topology.Components[:1]
 topology.Components[0].Kind="unknown"
 if err:=topology.Validate();err==nil {t.Fatal("accepted unknown component")}
}
