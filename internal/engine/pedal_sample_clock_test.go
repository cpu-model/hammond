package engine

import (
 "testing"

 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

func TestPedalFilterSampleRateMustMatchSharedGenerator(t *testing.T) {
 var pedalConfig [4]electrical.BusConfig
 var filterConfig [4]electrical.PedalRLConfig
 for i:=range pedalConfig {
  pedalConfig[i]=electrical.BusConfig{
   SourceResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
   LoadResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
  }
  filterConfig[i]=electrical.PedalRLConfig{
   InductanceHenry:electrical.Parameter{Value:0.5,Unit:"H",Provenance:electrical.ProvenanceExperimental},
   SeriesResistanceOhm:electrical.Parameter{Value:100,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
   LoadResistanceOhm:electrical.Parameter{Value:100,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
  }
 }
 e,err:=NewElectricalBusEngineWithPedals(48000,registration.Routing{},registration.Routing{},registration.ContactTiming{},registration.ContactTiming{},busTestConfig(),busTestConfig(),registration.PedalRouting{},registration.PedalBusRouting{},pedalConfig)
 if err!=nil { t.Fatal(err) }
 if err:=e.EnableExperimentalPedalFilters(44100,filterConfig);err==nil {
  t.Fatal("accepted mismatched pedal filter clock")
 }
 if err:=e.EnableExperimentalPedalFilters(48000,filterConfig);err!=nil {
  t.Fatalf("matching sample rate rejected: %v",err)
 }
}
