package engine

import (
 "testing"
 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

func TestB3PedalSharedGeneratorIntegration(t *testing.T) {
 var cfg [4]electrical.BusConfig
 for i:=range cfg {
  cfg[i]=electrical.BusConfig{
   SourceResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
   LoadResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
  }
 }
 wiring:=registration.PedalWiringEvidence{
  ContactOrder:registration.PedalContactOrder{0,1,2,3,4,5,6,7},
  Busbars:registration.PedalBusRouting{1,1,2,2,3,3,4,4},
  Source:"synthetic integration fixture",
 }
 e,err:=NewElectricalBusEngineWithB3PedalWiring(48000,
  registration.Routing{},registration.Routing{},
  registration.ContactTiming{},registration.ContactTiming{},
  busTestConfig(),busTestConfig(),wiring,cfg)
 if err!=nil { t.Fatal(err) }
 pedal,err:=e.Pedals()
 if err!=nil { t.Fatal(err) }
 if err:=pedal.Key(7,true);err!=nil { t.Fatal(err) }
 active:=false
 for i:=0;i<128;i++ {
  _,_,voltages,err:=e.NextAllBuses()
  if err!=nil { t.Fatal(err) }
  for _,v:=range voltages { if v!=0 { active=true } }
 }
 if !active { t.Fatal("pedal did not drive any bus") }
 if err:=pedal.Key(7,false);err!=nil { t.Fatal(err) }
 _,_,voltages,err:=e.NextAllBuses()
 if err!=nil { t.Fatal(err) }
 for _,v:=range voltages { if v!=0 { t.Fatal("released pedal still drives bus") } }
}
