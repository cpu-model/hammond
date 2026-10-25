package electrical

import "testing"

func TestT1PedalBoundaryPreservesIndependentSources(t *testing.T) {
 lower:=TransformerInputs{Transformer:TransformerT1}
 lower.Connected[0]=true
 lower.Tap[0]=7
 lower.Voltage[0]=0.75
 pedal:=PedalTransformerInputs{Connected:true,Voltage:-0.25,Source:Parameter{
  Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental,
 }}
 boundary,err:=NewT1InputBoundary(lower,pedal)
 if err!=nil { t.Fatal(err) }
 if boundary.LowerManual.Voltage[0]!=0.75||boundary.Pedals.Voltage!=-0.25 {
  t.Fatalf("sources altered: %+v",boundary)
 }
}

func TestT1PedalBoundaryRejectsInvalidSources(t *testing.T) {
 lower:=TransformerInputs{Transformer:TransformerT2}
 if _,err:=NewT1InputBoundary(lower,PedalTransformerInputs{});err==nil { t.Fatal("accepted T2 as T1") }
 lower.Transformer=TransformerT1
 for _,pedal:=range []PedalTransformerInputs{
  {Connected:true,Voltage:1},
  {Connected:true,Voltage:1,Source:Parameter{Value:-1,Unit:"ohm",Provenance:ProvenanceExperimental}},
  {Connected:true,Voltage:1,Source:Parameter{Value:1,Unit:"V",Provenance:ProvenanceExperimental}},
 } {
  if _,err:=NewT1InputBoundary(lower,pedal);err==nil { t.Fatalf("accepted invalid pedal source: %+v",pedal) }
 }
 if _,err:=NewT1InputBoundary(lower,PedalTransformerInputs{});err!=nil { t.Fatalf("disconnected pedal must be permitted: %v",err) }
}
