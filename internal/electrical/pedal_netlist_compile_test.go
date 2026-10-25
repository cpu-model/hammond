package electrical

import (
 "math"
 "testing"
)

func TestCompilePedalRLNetworkPreservesWiresAndComponents(t *testing.T) {
 topology:=PedalCircuitTopology{
  Wires:[]TerminalEdge{
   {From:"source",To:"R1.left",Source:"test fixture"},
   {From:"R1.right",To:"L1.left",Source:"test fixture"},
   {From:"L1.right",To:"ground",Source:"test fixture"},
  },
  Components:[]PedalComponentPlacement{
   {Designator:"R1",Kind:PedalResistor,TerminalA:"R1.left",TerminalB:"R1.right",Source:"test fixture"},
   {Designator:"L1",Kind:PedalReactor,TerminalA:"L1.left",TerminalB:"L1.right",Source:"test fixture"},
  },
 }
 values:=[]PedalComponentValue{
  {Designator:"R1",Parameter:Parameter{Value:100,Unit:"ohm",Provenance:ProvenanceExperimental}},
  {Designator:"L1",Parameter:Parameter{Value:1,Unit:"H",Provenance:ProvenanceExperimental}},
 }
 n,nodes,err:=CompilePedalRLNetwork(1000,topology,values,"ground")
 if err!=nil {t.Fatal(err)}
 if nodes["source"]!=nodes["R1.left"]||nodes["R1.right"]!=nodes["L1.left"]||nodes["ground"]!="" {
  t.Fatalf("incorrect wire equivalence: %v",nodes)
 }
 _,err=n.Next(map[string]float64{nodes["source"]:10})
 if err!=nil {t.Fatal(err)}
 if math.Abs(n.currents[0]-10.0/1100)>1e-10 {t.Fatalf("unexpected reactor current: %g",n.currents[0])}
}

func TestCompilePedalRLNetworkRejectsUnsourcedOrIncorrectValues(t *testing.T) {
 topology:=PedalCircuitTopology{Components:[]PedalComponentPlacement{
  {Designator:"R1",Kind:PedalResistor,TerminalA:"a",TerminalB:"b",Source:"test"},
 }}
 for _,values:=range [][]PedalComponentValue{
  nil,
  {{Designator:"R1",Parameter:Parameter{Value:100,Unit:"H",Provenance:ProvenanceExperimental}}},
  {{Designator:"R1",Parameter:Parameter{Value:100,Unit:"ohm",Provenance:ProvenanceSchematic}}},
  {{Designator:"R1",Parameter:Parameter{Value:-1,Unit:"ohm",Provenance:ProvenanceExperimental}}},
  {{Designator:"other",Parameter:Parameter{Value:100,Unit:"ohm",Provenance:ProvenanceExperimental}}},
 } {
  if _,_,err:=CompilePedalRLNetwork(1000,topology,values,"ground");err==nil {t.Fatalf("accepted invalid values: %+v",values)}
 }
}

func TestCompilePedalRLNetworkPreservesPedalSourceIdentity(t *testing.T) {
 topology:=PedalCircuitTopology{
  Wires:[]TerminalEdge{{From:"pedal-bus-0",To:"R1.left",Source:"test"}},
  Components:[]PedalComponentPlacement{{Designator:"R1",Kind:PedalResistor,TerminalA:"R1.left",TerminalB:"ground",Source:"test"}},
 }
 values:=[]PedalComponentValue{{Designator:"R1",Parameter:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental}}}
 n,nodes,err:=CompilePedalRLNetwork(48000,topology,values,"ground")
 if err!=nil {t.Fatal(err)}
 if nodes["pedal-bus-0"]!="pedal-bus-0"||nodes["R1.left"]!="pedal-bus-0" {t.Fatalf("source alias lost: %v",nodes)}
 if _,err=n.Next(map[string]float64{"pedal-bus-0":5});err!=nil {t.Fatal(err)}
}

func TestCompilePedalRLNetworkRejectsIdealSourceShorts(t *testing.T) {
 for _,wires:=range [][]TerminalEdge{
  {{From:"pedal-bus-0",To:"ground",Source:"test"}},
  {{From:"pedal-bus-0",To:"pedal-bus-1",Source:"test"}},
 } {
  topology:=PedalCircuitTopology{Wires:wires}
  if _,_,err:=CompilePedalRLNetwork(48000,topology,nil,"ground");err==nil {
   t.Fatalf("accepted conflicting ideal voltage source wires: %v",wires)
  }
 }
}
