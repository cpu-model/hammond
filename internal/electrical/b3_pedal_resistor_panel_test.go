package electrical

import "testing"

func TestB3PedalResistorPanelTopology(t *testing.T) {
 panel:=B3PedalResistorPanelTopology()
 if err:=panel.Validate();err!=nil {t.Fatal(err)}
 if len(panel.Components)!=6 {t.Fatalf("got %d components",len(panel.Components))}
 expected:=map[string][2]string{
  "R81":{"panel-common-in","panel-row-1-mid"},
  "R84":{"panel-row-1-mid","panel-common-out"},
  "R82":{"panel-common-in","panel-row-2-mid"},
  "R85":{"panel-row-2-mid","panel-common-out"},
  "R83":{"panel-common-in","panel-row-3-mid"},
  "R86":{"panel-row-3-mid","panel-common-out"},
 }
 for _,c:=range panel.Components {
  terminals,ok:=expected[c.Designator]
  if !ok {t.Fatalf("unexpected component %s",c.Designator)}
  if c.TerminalA!=terminals[0]||c.TerminalB!=terminals[1] {
   t.Fatalf("%s terminals %s/%s, want %s/%s",c.Designator,c.TerminalA,c.TerminalB,terminals[0],terminals[1])
  }
 }
 connected,err:=panel.SignalPathExists("panel-common-in","panel-common-out")
 if err!=nil||!connected {t.Fatalf("common rails disconnected: %v",err)}
}
