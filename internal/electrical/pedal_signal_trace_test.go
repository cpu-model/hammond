package electrical

import "testing"

func TestTracePedalSignalPathListsComponentsInOrder(t *testing.T) {
 topology:=PedalCircuitTopology{
  Wires:[]TerminalEdge{
   {From:"bus",To:"R80.a",Source:"synthetic fixture"},
   {From:"R80.b",To:"L20.a",Source:"synthetic fixture"},
   {From:"L20.b",To:"drawbar",Source:"synthetic fixture"},
  },
  Components:[]PedalComponentPlacement{
   {Designator:"R80",Kind:PedalResistor,TerminalA:"R80.a",TerminalB:"R80.b",Source:"synthetic fixture"},
   {Designator:"L20",Kind:PedalReactor,TerminalA:"L20.a",TerminalB:"L20.b",Source:"synthetic fixture"},
  },
 }
 path,ok,err:=topology.TraceSignalPath("bus","drawbar")
 if err!=nil||!ok {t.Fatalf("trace missing: %v",err)}
 if len(path.Components)!=2||path.Components[0]!="R80"||path.Components[1]!="L20" {
  t.Fatalf("wrong ordered components: %v",path.Components)
 }
 if path.Terminals[0]!="bus"||path.Terminals[len(path.Terminals)-1]!="drawbar" {
  t.Fatalf("wrong endpoints: %v",path.Terminals)
 }
 path,ok,err=topology.TraceSignalPath("bus","unconnected")
 if err!=nil||ok||len(path.Components)!=0 {t.Fatalf("unexpected trace: %+v, %v, %v",path,ok,err)}
}
