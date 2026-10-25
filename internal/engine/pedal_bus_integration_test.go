package engine

import (
 "testing"

 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

func TestSharedGeneratorPedalAndManualBuses(t *testing.T) {
 var upper,lower registration.Routing
 upper[0][0]=1
 var pedalRouting registration.PedalRouting
 pedalRouting[0][0]=1
 pedalRouting[0][1]=2
 mapping:=registration.PedalBusRouting{1,4}
 var pedalConfig [4]electrical.BusConfig
 for i:=range pedalConfig {
  pedalConfig[i]=electrical.BusConfig{
   SourceResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
   LoadResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
  }
 }
 e,err:=NewElectricalBusEngineWithPedals(48000,upper,lower,registration.ContactTiming{},registration.ContactTiming{},busTestConfig(),busTestConfig(),pedalRouting,mapping,pedalConfig)
 if err!=nil { t.Fatal(err) }
 manual,_:=e.Manual(0)
 pedal,_:=e.Pedals()
 _=manual.Key(0,true)
 _=pedal.Key(0,true)
 _,_,_,err=e.NextAllBuses()
 if err!=nil { t.Fatal(err) }
 u,l,p,err:=e.NextAllBuses()
 if err!=nil { t.Fatal(err) }
 if u[0]==0||p[0]==0||p[3]==0 { t.Fatalf("missing generator paths: %v %v",u,p) }
 if u[0]!=p[0] { t.Fatalf("manual and pedal must share wheel phase: %v vs %v",u[0],p[0]) }
 if l!=([9]float64{})||p[1]!=0||p[2]!=0 { t.Fatal("unexpected cross-bus leakage") }
 _=pedal.Key(0,false)
 _,_,p,err=e.NextAllBuses()
 if err!=nil||p!=([4]float64{}) { t.Fatalf("released pedal still audible: %v %v",p,err) }
}

func TestManualOnlyEngineRejectsPedalRendering(t *testing.T) {
 e,err:=NewElectricalBusEngine(48000,registration.Routing{},registration.Routing{},registration.ContactTiming{},registration.ContactTiming{},busTestConfig(),busTestConfig())
 if err!=nil { t.Fatal(err) }
 if _,err:=e.Pedals();err==nil { t.Fatal("unexpected pedal controller") }
 if _,_,_,err:=e.NextAllBuses();err==nil { t.Fatal("unexpected pedal buses") }
}
