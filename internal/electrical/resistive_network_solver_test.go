package electrical

import (
 "math"
 "testing"
)

func TestResistiveNetworkVoltageDivider(t *testing.T) {
 branches:=[]ResistiveBranch{{A:"source",B:"tap",Ohms:1000},{A:"tap",B:"",Ohms:1000}}
 result,err:=SolveResistiveNetwork(branches,map[string]float64{"source":10})
 if err!=nil {t.Fatal(err)}
 if math.Abs(result["tap"]-5)>1e-9 {t.Fatalf("divider: got %g, want 5",result["tap"])}
}

func TestResistiveNetworkParallelLoads(t *testing.T) {
 branches:=[]ResistiveBranch{
  {A:"source",B:"tap",Ohms:1000},
  {A:"tap",B:"",Ohms:1000},
  {A:"tap",B:"",Ohms:1000},
 }
 result,err:=SolveResistiveNetwork(branches,map[string]float64{"source":12})
 if err!=nil {t.Fatal(err)}
 if math.Abs(result["tap"]-4)>1e-9 {t.Fatalf("parallel divider: got %g, want 4",result["tap"])}
}

func TestResistiveNetworkRejectsFloatingAndInvalid(t *testing.T) {
 _,err:=SolveResistiveNetwork([]ResistiveBranch{{A:"a",B:"b",Ohms:100}},nil)
 if err==nil {t.Fatal("accepted floating network")}
 _,err=SolveResistiveNetwork([]ResistiveBranch{{A:"a",B:"",Ohms:0}},nil)
 if err==nil {t.Fatal("accepted zero ohm resistor")}
 _,err=SolveResistiveNetwork([]ResistiveBranch{{A:"a",B:"",Ohms:100}},map[string]float64{"a":math.NaN()})
 if err==nil {t.Fatal("accepted nonfinite source")}
}

func TestResistiveNetworkHighImpedanceDivider(t *testing.T) {
 result,err:=SolveResistiveNetwork([]ResistiveBranch{
  {A:"source",B:"tap",Ohms:1e12},
  {A:"tap",B:"",Ohms:1e12},
 },map[string]float64{"source":10})
 if err!=nil {t.Fatalf("valid teraohm network rejected: %v",err)}
 if math.Abs(result["tap"]-5)>1e-9 {t.Fatalf("high impedance divider: got %g, want 5",result["tap"])}
}
