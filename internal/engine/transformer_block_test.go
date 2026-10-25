package engine

import (
 "testing"

 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

func TestRenderTransformerSourcesBlock(t *testing.T) {
 var routing registration.Routing
 routing[0][0]=1
 e,err:=NewElectricalBusEngine(48000,routing,registration.Routing{},registration.ContactTiming{},registration.ContactTiming{},busTestConfig(),busTestConfig())
 if err!=nil { t.Fatal(err) }
 m,_:=e.Manual(0)
 _=m.Key(0,true)
 var selection electrical.DrawbarTapSelection
 selection[0]=8
 pair:=electrical.ManualTransformerPair{
  Upper:electrical.ManualTransformerConnection{ID:electrical.TransformerT2,Selection:selection},
  Lower:electrical.ManualTransformerConnection{ID:electrical.TransformerT1},
 }
 pedals:=[]electrical.PedalTransformerInputs{
  {Connected:false},
  {Connected:true,Voltage:0.3,Source:electrical.Parameter{Value:1000,Unit:"ohm",Provenance:electrical.ProvenanceExperimental}},
 }
 upper:=make([]electrical.TransformerInputs,len(pedals))
 lower:=make([]electrical.T1InputBoundary,len(pedals))
 if err:=e.RenderTransformerSources(pair,pedals,upper,lower);err!=nil { t.Fatal(err) }
 if upper[0].Transformer!=electrical.TransformerT2||!upper[0].Connected[0] { t.Fatalf("upper routing lost: %+v",upper[0]) }
 if lower[0].Pedals.Connected||!lower[1].Pedals.Connected||lower[1].Pedals.Voltage!=0.3 { t.Fatalf("pedal block lost: %+v",lower) }
}

func TestRenderTransformerSourcesRejectsWholeInvalidBlock(t *testing.T) {
 makeEngine:=func()*ElectricalBusEngine {
  var routing registration.Routing
  routing[0][0]=1
  e,err:=NewElectricalBusEngine(48000,routing,registration.Routing{},registration.ContactTiming{},registration.ContactTiming{},busTestConfig(),busTestConfig())
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
 pedals:=[]electrical.PedalTransformerInputs{{},{Connected:true}}
 upper:=make([]electrical.TransformerInputs,2)
 lower:=make([]electrical.T1InputBoundary,2)
 if err:=a.RenderTransformerSources(pair,pedals,upper,lower);err==nil { t.Fatal("accepted invalid later pedal sample") }
 if upper[0]!=(electrical.TransformerInputs{})||lower[0]!=(electrical.T1InputBoundary{}) { t.Fatal("partial output written on invalid block") }
 av,_:=a.NextBuses()
 bv,_:=b.NextBuses()
 if av!=bv { t.Fatal("invalid block advanced engine") }
 if err:=a.RenderTransformerSources(pair,pedals,upper[:1],lower);err==nil { t.Fatal("accepted mismatched buffers") }
}
