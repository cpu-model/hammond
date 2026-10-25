package electrical

import "testing"

func TestB3PedalResistorPanelTopology(t *testing.T) {
 panel:=B3PedalResistorPanelTopology()
 if err:=panel.Validate();err!=nil {t.Fatal(err)}
 if len(panel.Components)!=6 {t.Fatalf("got %d components",len(panel.Components))}
 expected:=map[string][2]string{
  "R81":{"panel-row-1-in","panel-row-1-mid"},
  "R84":{"panel-row-1-mid","panel-common-out"},
  "R82":{"panel-row-2-in","panel-row-2-mid"},
  "R85":{"panel-row-2-mid","panel-common-out"},
  "R83":{"panel-row-3-in","panel-row-3-mid"},
  "R86":{"panel-row-3-mid","panel-common-out"},
 }
 for _,c:=range panel.Components {
  terminals,ok:=expected[c.Designator]
  if !ok {t.Fatalf("unexpected component %s",c.Designator)}
  if c.TerminalA!=terminals[0]||c.TerminalB!=terminals[1] {
   t.Fatalf("%s terminals %s/%s, want %s/%s",c.Designator,c.TerminalA,c.TerminalB,terminals[0],terminals[1])
  }
 }
 for i:=1;i<=3;i++ {
  in:="panel-row-"+string(rune('0'+i))+"-in"
  connected,err:=panel.SignalPathExists(in,"panel-common-out")
  if err!=nil||!connected {t.Fatalf("row %d disconnected: %v",i,err)}
 }
}
