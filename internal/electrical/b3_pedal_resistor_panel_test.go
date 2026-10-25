package electrical

import "testing"

func TestB3PedalResistorPanelTopology(t *testing.T) {
 panel:=B3PedalResistorPanelTopology()
 if err:=panel.Validate();err!=nil {t.Fatal(err)}
 if len(panel.Components)!=6 {t.Fatalf("got %d components",len(panel.Components))}
 for _,row:=range []string{"1","2","3"} {
  ok,err:=panel.SignalPathExists("panel-row-"+row+"-in","panel-common-out")
  if err!=nil||!ok {t.Fatalf("row %s not connected: %v",row,err)}
 }
 for _,pair:=range [][2]string{{"1","2"},{"1","3"},{"2","3"}} {
  ok,err:=panel.SignalPathExists("panel-row-"+pair[0]+"-in","panel-row-"+pair[1]+"-in")
  if err!=nil {t.Fatal(err)}
  if !ok {t.Fatalf("rows %s and %s must join at common output rail",pair[0],pair[1])}
 }
}
