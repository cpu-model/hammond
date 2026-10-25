package electrical

import "testing"

// An explicit, synthetic mapping must compile into the same nodes as
// the engine's four pedal-bus sources. No factory mapping is asserted.
func TestPedalPanelMappedBusCompilerIntegration(t *testing.T) {
 topology,err:=ConnectPedalPanelInputs(B3PedalResistorPanelTopology(),[3]int{0,2,3},"synthetic integration test")
 if err!=nil {t.Fatal(err)}
 values:=[]PedalComponentValue{}
 for _,c:=range topology.Components {
  values=append(values,PedalComponentValue{
   Designator:c.Designator,
   Parameter:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental},
  })
 }
 n,nodes,err:=CompilePedalRLNetwork(48000,topology,values,"chassis-ground")
 if err!=nil {t.Fatal(err)}
 for _,tc:=range []struct{terminal,bus string}{
  {"panel-row-1-in","pedal-bus-0"},
  {"panel-row-2-in","pedal-bus-2"},
  {"panel-row-3-in","pedal-bus-3"},
 } {
  if got:=nodes[tc.terminal];got!=tc.bus {
   t.Fatalf("%s compiled to %q, want %q",tc.terminal,got,tc.bus)
  }
 }
 result,err:=n.Next(map[string]float64{
  "pedal-bus-0":1,"pedal-bus-1":0,"pedal-bus-2":0,"pedal-bus-3":0,
  nodes["panel-common-out"]:0,
 })
 if err!=nil {t.Fatal(err)}
 if got:=result[nodes["panel-row-1-mid"]];got!=0.5 {
  t.Fatalf("driven midpoint got %g want 0.5",got)
 }
 for _,row:=range []string{"panel-row-2-mid","panel-row-3-mid"} {
  if got:=result[nodes[row]];got!=0 {
   t.Fatalf("inactive row %s leaked: %g",row,got)
  }
 }
}
