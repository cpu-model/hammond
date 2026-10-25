package electrical

import "testing"

func TestPedalComponentCutDetectsBypass(t *testing.T) {
 circuit:=PedalCircuitTopology{
  Wires:[]TerminalEdge{
   {From:"input",To:"R80.a",Source:"synthetic"},
   {From:"R80.b",To:"output",Source:"synthetic"},
  },
  Components:[]PedalComponentPlacement{
   {Designator:"R80",Kind:PedalResistor,TerminalA:"R80.a",TerminalB:"R80.b",Source:"synthetic"},
  },
 }
 mandatory,err:=circuit.ComponentOnEverySignalPath("input","output","R80")
 if err!=nil||!mandatory {t.Fatalf("series resistor should be mandatory: %v",err)}
 circuit.Components=append(circuit.Components,PedalComponentPlacement{
  Designator:"R81",Kind:PedalResistor,TerminalA:"input",TerminalB:"output",Source:"synthetic parallel branch",
 })
 mandatory,err=circuit.ComponentOnEverySignalPath("input","output","R80")
 if err!=nil||mandatory {t.Fatalf("parallel resistor should bypass R80: %v",err)}
 mandatory,err=circuit.ComponentOnEverySignalPath("input","output","R81")
 if err!=nil||mandatory {t.Fatalf("parallel resistor should bypass R81: %v",err)}
 if _,err=circuit.ComponentOnEverySignalPath("input","output","R99");err==nil {
  t.Fatal("accepted nonexistent component")
 }
 if _,err=circuit.ComponentOnEverySignalPath("input","isolated","R80");err==nil {
  t.Fatal("accepted disconnected endpoints")
 }
}
