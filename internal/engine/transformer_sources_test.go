package engine

import (
 "testing"

 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

func TestElectricalEngineTransformerSourcesKeepsPedalSeparate(t *testing.T) {
 var route registration.Routing
 route[0][0]=1
 e,err:=NewElectricalBusEngine(48000,route,registration.Routing{},registration.ContactTiming{},registration.ContactTiming{},busTestConfig(),busTestConfig())
 if err!=nil { t.Fatal(err) }
 m,_:=e.Manual(0)
 _=m.Key(0,true)
 var taps electrical.DrawbarTapSelection
 taps[0]=8
 pair:=electrical.ManualTransformerPair{
  Upper:electrical.ManualTransformerConnection{ID:electrical.TransformerT2,Selection:taps},
  Lower:electrical.ManualTransformerConnection{ID:electrical.TransformerT1},
 }
 pedal:=electrical.PedalTransformerInputs{Connected:true,Voltage:0.625,Source:electrical.Parameter{
  Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental,
 }}
 t2,t1,err:=e.NextTransformerSources(pair,pedal)
 if err!=nil { t.Fatal(err) }
 if t2.Transformer!=electrical.TransformerT2||!t2.Connected[0]||t2.Tap[0]!=8 { t.Fatalf("incorrect upper source: %+v",t2) }
 if t1.LowerManual.Transformer!=electrical.TransformerT1||t1.Pedals.Voltage!=0.625||!t1.Pedals.Connected {
  t.Fatalf("incorrect T1 pedal boundary: %+v",t1)
 }
}

func TestInvalidPedalSourceDoesNotAdvanceGenerator(t *testing.T) {
 makeEngine:=func()*ElectricalBusEngine {
  var route registration.Routing
  route[0][0]=1
  e,err:=NewElectricalBusEngine(48000,route,registration.Routing{},registration.ContactTiming{},registration.ContactTiming{},busTestConfig(),busTestConfig())
  if err!=nil { t.Fatal(err) }
  m,_:=e.Manual(0)
  _=m.Key(0,true)
  return e
 }
 a,b:=makeEngine(),makeEngine()
 pair:=electrical.ManualTransformerPair{
  Upper:electrical.ManualTransformerConnection{ID:electrical.TransformerT2},
  Lower:electrical.ManualTransformerConnection{ID:electrical.TransformerT1},
 }
 invalid:=electrical.PedalTransformerInputs{Connected:true}
 if _,_,err:=a.NextTransformerSources(pair,invalid);err==nil { t.Fatal("accepted invalid pedal source") }
 av,_:=a.NextBuses()
 bv,_:=b.NextBuses()
 if av!=bv { t.Fatalf("invalid pedal input consumed sample: %v != %v",av,bv) }
}
