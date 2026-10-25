package electrical

import "testing"

func TestB3PedalAnalogComponentInventory(t *testing.T) {
 topology:=B3PedalAnalogComponents()
 if err:=topology.Validate();err!=nil {t.Fatal(err)}
 if len(topology.Components)!=8 {t.Fatalf("got %d components; want 8",len(topology.Components))}
 seen:=map[string]bool{}
 for _,c:=range topology.Components {
  if seen[c.Designator] {t.Fatalf("duplicate %s",c.Designator)}
  seen[c.Designator]=true
 }
 for _,name:=range []string{"R80","R81","R82","R83","R84","R85","R86","L20"} {
  if !seen[name] {t.Fatalf("missing %s",name)}
 }
 for _,term:=range []string{"panel-row-1-in","panel-row-2-in","panel-row-3-in"} {
  connected,err:=topology.SignalPathExists(term,"panel-common-out")
  if err!=nil {t.Fatal(err)}
  if !connected {t.Fatalf("row %s has no path to panel output",term)}
 }
}
