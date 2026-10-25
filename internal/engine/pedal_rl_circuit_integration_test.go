package engine

import (
 "testing"

 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

func TestSharedPedalBusDrivesExperimentalRLCircuit(t *testing.T) {
 var routing registration.PedalRouting
 routing[0][0]=1
 var configs [4]electrical.BusConfig
 for i:=range configs {
  configs[i]=electrical.BusConfig{
   SourceResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
   LoadResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
  }
 }
 e,err:=NewElectricalBusEngineWithPedals(48000,registration.Routing{},registration.Routing{},registration.ContactTiming{},registration.ContactTiming{},busTestConfig(),busTestConfig(),routing,registration.PedalBusRouting{1},configs)
 if err!=nil {t.Fatal(err)}
 if err=e.EnableExperimentalPedalCircuit(
  []electrical.ResistiveBranch{{A:"pedal-bus-0",B:"tap",Ohms:1000},{A:"tap",B:"",Ohms:1000}},
  []electrical.InductiveBranch{{A:"tap",B:"",Henry:1}},
 );err!=nil {t.Fatal(err)}
 p,_:=e.Pedals()
 _=p.Key(0,true)
 if _,_,_,err=e.NextExperimentalPedalCircuit("unknown");err==nil {t.Fatal("accepted unknown output")}
 _,_,_,err=e.NextExperimentalPedalCircuit("tap")
 if err!=nil {t.Fatal(err)}
 _,_,voltage,err:=e.NextExperimentalPedalCircuit("tap")
 if err!=nil {t.Fatal(err)}
 if voltage==0 {t.Fatal("pedal circuit produced no signal")}
 e.DisableExperimentalPedalCircuit()
 if _,_,_,err=e.NextExperimentalPedalCircuit("tap");err==nil {t.Fatal("disabled circuit still active")}
}

func TestCompiledPedalCircuitThroughEngine(t *testing.T) {
 var routing registration.PedalRouting
 routing[0][0]=1
 var configs [4]electrical.BusConfig
 for i:=range configs {
  configs[i]=electrical.BusConfig{
   SourceResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
   LoadResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
  }
 }
 e,err:=NewElectricalBusEngineWithPedals(48000,registration.Routing{},registration.Routing{},registration.ContactTiming{},registration.ContactTiming{},busTestConfig(),busTestConfig(),routing,registration.PedalBusRouting{1},configs)
 if err!=nil {t.Fatal(err)}
 topology:=electrical.PedalCircuitTopology{
  Wires:[]electrical.TerminalEdge{
   {From:"pedal-bus-0",To:"R1.left",Source:"synthetic test"},
   {From:"R1.right",To:"L1.left",Source:"synthetic test"},
   {From:"L1.right",To:"ground",Source:"synthetic test"},
  },
  Components:[]electrical.PedalComponentPlacement{
   {Designator:"R1",Kind:electrical.PedalResistor,TerminalA:"R1.left",TerminalB:"R1.right",Source:"synthetic test"},
   {Designator:"L1",Kind:electrical.PedalReactor,TerminalA:"L1.left",TerminalB:"L1.right",Source:"synthetic test"},
  },
 }
 values:=[]electrical.PedalComponentValue{
  {Designator:"R1",Parameter:electrical.Parameter{Value:100,Unit:"ohm",Provenance:electrical.ProvenanceExperimental}},
  {Designator:"L1",Parameter:electrical.Parameter{Value:1,Unit:"H",Provenance:electrical.ProvenanceExperimental}},
 }
 if err=e.EnableCompiledPedalCircuit(topology,values,"ground");err!=nil {t.Fatal(err)}
 p,_:=e.Pedals()
 _=p.Key(0,true)
 if _,_,_,err=e.NextCompiledPedalCircuit("missing");err==nil {t.Fatal("accepted missing terminal")}
 _,_,_,err=e.NextCompiledPedalCircuit("L1.left")
 if err!=nil {t.Fatal(err)}
 _,_,voltage,err:=e.NextCompiledPedalCircuit("R1.right")
 if err!=nil {t.Fatal(err)}
 if voltage==0 {t.Fatal("no pedal signal at compiled output")}
 if err=e.EnableCompiledPedalCircuit(topology,nil,"ground");err==nil {t.Fatal("accepted missing values")}
 if _,_,_,err=e.NextCompiledPedalCircuit("R1.right");err!=nil {t.Fatalf("failed reconfiguration destroyed existing circuit: %v",err)}
}
