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
 connected,err:=topology.SignalPathExists("filter-R80-a","panel-common-out")
 if err!=nil {t.Fatal(err)}
 if connected {t.Fatal("R80-to-panel connection is unverified")}
 connected,err=topology.SignalPathExists("filter-R80-a","filter-L20-b")
 if err!=nil {t.Fatal(err)}
 if connected {t.Fatal("R80-to-L20 connection is unverified")}
}
