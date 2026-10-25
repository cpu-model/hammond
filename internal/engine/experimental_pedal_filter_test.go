package engine

import (
 "testing"

 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

func TestExperimentalPedalFiltersOnSharedGenerator(t *testing.T) {
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
 if err!=nil { t.Fatal(err) }
 p,_:=e.Pedals()
 _=p.Key(0,true)
 var filterConfigs [4]electrical.PedalRLConfig
 for i:=range filterConfigs {
  filterConfigs[i]=electrical.PedalRLConfig{
   InductanceHenry:electrical.Parameter{Value:0.5,Unit:"H",Provenance:electrical.ProvenanceExperimental},
   SeriesResistanceOhm:electrical.Parameter{Value:100,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
   LoadResistanceOhm:electrical.Parameter{Value:100,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
  }
 }
 if err:=e.EnableExperimentalPedalFilters(48000,filterConfigs);err!=nil { t.Fatal(err) }
 _,_,first,err:=e.NextExperimentalPedalBuses()
 if err!=nil { t.Fatal(err) }
 _,_,second,err:=e.NextExperimentalPedalBuses()
 if err!=nil { t.Fatal(err) }
 if second[0]==0||second[1]!=0||second[2]!=0||second[3]!=0 { t.Fatalf("unexpected pedal filtered output: %v",second) }
 if first==second { t.Fatal("filter state did not advance") }
 e.DisableExperimentalPedalFilters()
 _,_,raw,err:=e.NextExperimentalPedalBuses()
 if err!=nil||raw[0]==0 { t.Fatalf("raw pedal path unavailable: %v %v",raw,err) }
}

func TestExperimentalPedalFilterSetupFailureIsNonDestructive(t *testing.T) {
 e,err:=NewElectricalBusEngine(48000,registration.Routing{},registration.Routing{},registration.ContactTiming{},registration.ContactTiming{},busTestConfig(),busTestConfig())
 if err!=nil { t.Fatal(err) }
 if err:=e.EnableExperimentalPedalFilters(48000,[4]electrical.PedalRLConfig{});err==nil {
  t.Fatal("accepted filter without pedal network")
 }
}
