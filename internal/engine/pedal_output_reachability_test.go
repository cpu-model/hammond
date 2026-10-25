package engine

import (
 "testing"
 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

func TestExperimentalPedalCircuitRejectsGroundOnlyOutputBeforeSample(t *testing.T) {
 var config [4]electrical.BusConfig
 for i:=range config {
  config[i]=electrical.BusConfig{
   SourceResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
   LoadResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
  }
 }
 e,err:=NewElectricalBusEngineWithPedals(48000,registration.Routing{},registration.Routing{},registration.ContactTiming{},registration.ContactTiming{},busTestConfig(),busTestConfig(),registration.PedalRouting{},registration.PedalBusRouting{},config)
 if err!=nil {t.Fatal(err)}
 if err=e.EnableExperimentalPedalCircuit([]electrical.ResistiveBranch{
  {A:"pedal-bus-0",B:"",Ohms:1000},
  {A:"isolated",B:"",Ohms:1000},
 },nil);err!=nil {t.Fatal(err)}
 if _,_,_,err=e.NextExperimentalPedalCircuit("isolated");err==nil {
  t.Fatal("ground-only connection must not count as pedal signal path")
 }
 if _,_,_,err=e.NextExperimentalPedalCircuit("pedal-bus-0");err!=nil {
  t.Fatalf("pedal source node rejected: %v",err)
 }
}
