package electrical

import "testing"

func TestConnectPedalPanelInputsExplicitMapping(t *testing.T) {
 original:=B3PedalResistorPanelTopology()
 mapped,err:=ConnectPedalPanelInputs(original,[3]int{0,2,3},"synthetic test mapping; not factory wiring")
 if err!=nil {t.Fatal(err)}
 if len(original.Wires)!=0 {t.Fatal("mutated original topology")}
 if len(mapped.Wires)!=3 {t.Fatalf("got %d wires",len(mapped.Wires))}
 for _,tc:=range []struct{row,bus string}{
  {"panel-row-1-in","pedal-bus-0"},
  {"panel-row-2-in","pedal-bus-2"},
  {"panel-row-3-in","pedal-bus-3"},
 } {
  ok,err:=(TerminalGraph{Edges:mapped.Wires}).Connected(tc.row,tc.bus)
  if err!=nil||!ok {t.Fatalf("%s not wired to %s: %v",tc.row,tc.bus,err)}
 }
 if _,err:=ConnectPedalPanelInputs(original,[3]int{0,1,4},"test");err==nil {
  t.Fatal("accepted invalid bus index")
 }
 if _,err:=ConnectPedalPanelInputs(original,[3]int{0,1,2},"");err==nil {
  t.Fatal("accepted mapping without provenance")
 }
 // Reusing a bus is permitted by this assembly API, but only with
 // explicit evidence. The compiler still forbids shorting different buses.
 if _,err:=ConnectPedalPanelInputs(original,[3]int{1,1,2},"synthetic test");err!=nil {
  t.Fatalf("same-bus fanout rejected: %v",err)
 }
}
