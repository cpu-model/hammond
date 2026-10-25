package electrical

import (
 "math"
 "testing"
)

// Synthetic component values test the three-branch topology and solver,
// not the historical resistance values of a Hammond B3.
func TestB3PedalResistorPanelSyntheticVoltageDivider(t *testing.T) {
 topology:=B3PedalResistorPanelTopology()
 values:=[]PedalComponentValue{
  {Designator:"R81",Parameter:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental}},
  {Designator:"R84",Parameter:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental}},
  {Designator:"R82",Parameter:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental}},
  {Designator:"R85",Parameter:Parameter{Value:3000,Unit:"ohm",Provenance:ProvenanceExperimental}},
  {Designator:"R83",Parameter:Parameter{Value:3000,Unit:"ohm",Provenance:ProvenanceExperimental}},
  {Designator:"R86",Parameter:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental}},
 }
 network,nodes,err:=CompilePedalRLNetwork(48000,topology,values,"chassis-ground")
 if err!=nil {t.Fatal(err)}
 result,err:=network.Next(map[string]float64{
  nodes["panel-common-in"]:1,
  nodes["panel-common-out"]:0,
  "pedal-bus-0":0,"pedal-bus-1":0,"pedal-bus-2":0,"pedal-bus-3":0,
 })
 if err!=nil {t.Fatal(err)}
 for _,tc:=range []struct{node string;want float64}{
  {"panel-row-1-mid",0.5},
  {"panel-row-2-mid",0.75},
  {"panel-row-3-mid",0.25},
 } {
  got:=result[nodes[tc.node]]
  if math.Abs(got-tc.want)>1e-9 {
   t.Fatalf("%s: got %.12f want %.12f",tc.node,got,tc.want)
  }
 }
}
