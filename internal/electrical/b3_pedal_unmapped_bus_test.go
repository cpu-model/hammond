package electrical

import "testing"

// A fourth pedal bus that is not connected to any of the three
// synthetic panel inputs must not affect the resistor-panel output.
// This is a compiler/solver isolation check, not factory B3 wiring.
func TestPedalPanelUnmappedBusDoesNotDrivePanel(t *testing.T) {
 topology,err:=ConnectPedalPanelInputs(B3PedalResistorPanelTopology(),[3]int{0,2,3},"synthetic unmapped-bus regression")
 if err!=nil {t.Fatal(err)}
 values:=make([]PedalComponentValue,0,len(topology.Components))
 for _,c:=range topology.Components {
  values=append(values,PedalComponentValue{
   Designator:c.Designator,
   Parameter:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental},
  })
 }
 n,nodes,err:=CompilePedalRLNetwork(48000,topology,values,"chassis-ground")
 if err!=nil {t.Fatal(err)}
 fixed:=map[string]float64{
  "pedal-bus-0":0,"pedal-bus-1":1,"pedal-bus-2":0,"pedal-bus-3":0,
  nodes["panel-common-out"]:0,
 }
 result,err:=n.Next(fixed)
 if err!=nil {t.Fatal(err)}
 for _,name:=range []string{"panel-row-1-mid","panel-row-2-mid","panel-row-3-mid"} {
  if got:=result[nodes[name]];got!=0 {
   t.Fatalf("unmapped bus drove %s: %g",name,got)
  }
 }
}
