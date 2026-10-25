package engine

import (
 "testing"

 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

func TestElectricalBusEngineTransformerInputs(t *testing.T) {
 var upper,lower registration.Routing
 upper[0][0]=1
 lower[0][8]=2
 e,err:=NewElectricalBusEngine(48000,upper,lower,registration.ContactTiming{},registration.ContactTiming{},busTestConfig(),busTestConfig())
 if err!=nil { t.Fatal(err) }
 u,_:=e.Manual(0)
 l,_:=e.Manual(1)
 _=u.Key(0,true)
 _=l.Key(0,true)
 var ut,lt electrical.DrawbarTapSelection
 ut[0]=8
 lt[8]=3
 pair:=electrical.ManualTransformerPair{
  Upper:electrical.ManualTransformerConnection{ID:electrical.TransformerT2,Selection:ut},
  Lower:electrical.ManualTransformerConnection{ID:electrical.TransformerT1,Selection:lt},
 }
 _,_,err=e.NextTransformerInputs(pair)
 if err!=nil { t.Fatal(err) }
 a,b,err:=e.NextTransformerInputs(pair)
 if err!=nil { t.Fatal(err) }
 if a.Transformer!=electrical.TransformerT2||b.Transformer!=electrical.TransformerT1 {
  t.Fatal("transformer identities swapped")
 }
 if !a.Connected[0]||a.Tap[0]!=8||a.Voltage[0]==0 {
  t.Fatalf("upper transformer input: %+v",a)
 }
 if !b.Connected[8]||b.Tap[8]!=3||b.Voltage[8]==0 {
  t.Fatalf("lower transformer input: %+v",b)
 }
 if a.Connected[8]||b.Connected[0] { t.Fatal("manual paths mixed") }
}

func TestInvalidTransformerPairDoesNotAdvanceElectricalEngine(t *testing.T) {
 var route registration.Routing
 route[0][0]=1
 makeEngine:=func()*ElectricalBusEngine {
  e,err:=NewElectricalBusEngine(48000,route,registration.Routing{},registration.ContactTiming{},registration.ContactTiming{},busTestConfig(),busTestConfig())
  if err!=nil { t.Fatal(err) }
  m,_:=e.Manual(0)
  _=m.Key(0,true)
  return e
 }
 a,b:=makeEngine(),makeEngine()
 invalid:=electrical.ManualTransformerPair{}
 if _,_,err:=a.NextTransformerInputs(invalid);err==nil { t.Fatal("accepted invalid pair") }
 av,_:=a.NextBuses()
 bv,_:=b.NextBuses()
 if av!=bv { t.Fatalf("invalid configuration advanced state: %v vs %v",av,bv) }
}
