package engine

import (
 "testing"
 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

func TestB3FactoryPedalsSharedEngine(t *testing.T) {
 var pedalCfg [4]electrical.BusConfig
 for i:=range pedalCfg {
  pedalCfg[i]=electrical.BusConfig{
   SourceResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
   LoadResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
  }
 }
 e,err:=NewElectricalBusEngineWithB3FactoryPedals(
  48000,registration.Routing{},registration.Routing{},
  registration.ContactTiming{},registration.ContactTiming{},
  busTestConfig(),busTestConfig(),pedalCfg)
 if err!=nil {t.Fatal(err)}
 p,err:=e.Pedals()
 if err!=nil {t.Fatal(err)}
 if err:=p.Key(0,true);err!=nil {t.Fatal(err)}
 active:=false
 for i:=0;i<128;i++ {
  upper,lower,pedal,err:=e.NextAllBuses()
  if err!=nil {t.Fatal(err)}
  for _,v:=range upper {if v!=0 {t.Fatal("unexpected upper manual voltage")}}
  for _,v:=range lower {if v!=0 {t.Fatal("unexpected lower manual voltage")}}
  for _,v:=range pedal {if v!=0 {active=true}}
 }
 if !active {t.Fatal("factory pedal mapping produced no signal")}
 if err:=p.Key(0,false);err!=nil {t.Fatal(err)}
 _,_,pedal,err:=e.NextAllBuses()
 if err!=nil {t.Fatal(err)}
 for _,v:=range pedal {if v!=0 {t.Fatal("released pedal still active")}}
}
