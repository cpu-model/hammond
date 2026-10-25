package electrical

import "testing"

// The three resistor-panel inputs are distinct voltage nodes. This
// regression detects an accidental reintroduction of a common input rail.
func TestB3PedalResistorPanelIndependentInputs(t *testing.T) {
 topology:=B3PedalResistorPanelTopology()
 values:=[]PedalComponentValue{
  {Designator:"R81",Parameter:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental}},
  {Designator:"R84",Parameter:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental}},
  {Designator:"R82",Parameter:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental}},
  {Designator:"R85",Parameter:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental}},
  {Designator:"R83",Parameter:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental}},
  {Designator:"R86",Parameter:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental}},
 }
 network,nodes,err:=CompilePedalRLNetwork(48000,topology,values,"chassis-ground")
 if err!=nil {t.Fatal(err)}
 result,err:=network.Next(map[string]float64{
  nodes["panel-row-1-in"]:1,
  nodes["panel-row-2-in"]:0,
  nodes["panel-row-3-in"]:0,
  nodes["panel-common-out"]:0,
  "pedal-bus-0":0,"pedal-bus-1":0,"pedal-bus-2":0,"pedal-bus-3":0,
 })
 if err!=nil {t.Fatal(err)}
 if result[nodes["panel-row-1-mid"]]!=0.5 {t.Fatalf("driven row midpoint: %v",result[nodes["panel-row-1-mid"]])}
 for _,row:=range []string{"panel-row-2-mid","panel-row-3-mid"} {
  if result[nodes[row]]!=0 {t.Fatalf("%s leaked from another input: %v",row,result[nodes[row]])}
 }
}
