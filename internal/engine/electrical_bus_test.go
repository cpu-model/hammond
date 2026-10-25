package engine

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

func busTestConfig() [9]electrical.BusConfig {
 var c [9]electrical.BusConfig
 for i:=range c {
  c[i]=electrical.BusConfig{
   SourceResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
   LoadResistanceOhm:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental},
  }
 }
 return c
}

func TestElectricalBusEngineIndependentManualsAndTiming(t *testing.T) {
 var upper,lower registration.Routing
 upper[0][0]=1
 lower[0][8]=2
 var timing registration.ContactTiming
 timing.Make[0]=2
 e,err:=NewElectricalBusEngine(48000,upper,lower,timing,registration.ContactTiming{},busTestConfig(),busTestConfig())
 if err!=nil { t.Fatal(err) }
 u,_:=e.Manual(0)
 l,_:=e.Manual(1)
 _=u.Key(0,true)
 _=l.Key(0,true)
 for sample:=0;sample<4;sample++ {
  a,b:=e.NextBuses()
  if sample<2&&a[0]!=0 { t.Fatalf("upper closed too early at %d",sample) }
  for i:=1;i<9;i++ { if a[i]!=0 { t.Fatalf("unexpected upper bus %d",i) } }
  for i:=0;i<8;i++ { if b[i]!=0 { t.Fatalf("unexpected lower bus %d",i) } }
 }
 a,b:=e.NextBuses()
 if math.Abs(a[0])<1e-12||math.Abs(b[8])<1e-12 {
  t.Fatalf("expected both independent buses active: %v %v",a,b)
 }
}

func TestElectricalBusEngineRejectsInvalidConfig(t *testing.T) {
 cfg:=busTestConfig()
 cfg[4].LoadResistanceOhm.Value=0
 if _,err:=NewElectricalBusEngine(48000,registration.Routing{},registration.Routing{},registration.ContactTiming{},registration.ContactTiming{},cfg,busTestConfig());err==nil {
  t.Fatal("accepted invalid bus config")
 }
}

func TestElectricalBusEngineNoAllocations(t *testing.T) {
 var route registration.Routing
 route[0][0]=1
 e,err:=NewElectricalBusEngine(48000,route,registration.Routing{},registration.ContactTiming{},registration.ContactTiming{},busTestConfig(),busTestConfig())
 if err!=nil { t.Fatal(err) }
 m,_:=e.Manual(0)
 _=m.Key(0,true)
 if n:=testing.AllocsPerRun(100,func(){ _,_=e.NextBuses() });n!=0 {
  t.Fatalf("allocations per sample: %v",n)
 }
}
