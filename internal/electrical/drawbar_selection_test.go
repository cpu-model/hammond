package electrical

import "testing"

func TestDrawbarTapSelectionMasksOnlyDisconnectedBuses(t *testing.T) {
 var selection DrawbarTapSelection
 selection[0]=8
 selection[2]=1
 var buses [9]float64
 buses[0],buses[1],buses[2]=1,2,3
 out,err:=selection.SelectedBusVoltages(buses)
 if err!=nil { t.Fatal(err) }
 if out[0]!=1||out[1]!=0||out[2]!=3 { t.Fatalf("unexpected routing: %v",out) }
 if buses[1]!=2 { t.Fatal("input was mutated") }
}

func TestDrawbarTapSelectionRejectsInvalidPosition(t *testing.T) {
 var selection DrawbarTapSelection
 selection[5]=9
 if err:=selection.Validate();err==nil { t.Fatal("accepted position 9") }
 if _,err:=selection.SelectedBusVoltages([9]float64{});err==nil { t.Fatal("accepted invalid selection") }
}

func TestDrawbarTapSelectionNoAllocations(t *testing.T) {
 var selection DrawbarTapSelection
 selection[0]=8
 if n:=testing.AllocsPerRun(100,func(){
  _,_=selection.SelectedBusVoltages([9]float64{1})
 });n!=0 { t.Fatalf("allocations: %v",n) }
}
