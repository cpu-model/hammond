package electrical

import (
 "math"
 "testing"
)

func TestResistiveNodeVoltageDivider(t *testing.T) {
 var n ResistiveNode
 if err:=n.AddSource(10,1000);err!=nil { t.Fatal(err) }
 if err:=n.AddGroundLoad(1000);err!=nil { t.Fatal(err) }
 v,err:=n.Voltage()
 if err!=nil||math.Abs(v-5)>1e-12 { t.Fatalf("voltage=%g err=%v",v,err) }
}

func TestResistiveNodeMultipleSources(t *testing.T) {
 var n ResistiveNode
 if err:=n.AddSource(6,1000);err!=nil { t.Fatal(err) }
 if err:=n.AddSource(0,2000);err!=nil { t.Fatal(err) }
 v,err:=n.Voltage()
 if err!=nil||math.Abs(v-4)>1e-12 { t.Fatalf("voltage=%g err=%v",v,err) }
 n.Reset()
 if _,err:=n.Voltage();err==nil { t.Fatal("expected floating node") }
}

func TestResistiveNodeRejectsInvalidValues(t *testing.T) {
 var n ResistiveNode
 if err:=n.AddSource(math.NaN(),1000);err==nil { t.Fatal("accepted NaN voltage") }
 if err:=n.AddSource(1,0);err==nil { t.Fatal("accepted zero resistance") }
 if err:=n.AddGroundLoad(math.Inf(1));err==nil { t.Fatal("accepted infinite load") }
}

func TestResistiveNodeNoAllocations(t *testing.T) {
 var n ResistiveNode
 if count:=testing.AllocsPerRun(100,func(){
  n.Reset()
  _=n.AddSource(2,1000)
  _=n.AddGroundLoad(2000)
  _,_=n.Voltage()
 });count!=0 { t.Fatalf("allocations: %v",count) }
}
