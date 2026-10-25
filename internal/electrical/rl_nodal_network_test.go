package electrical

import (
 "math"
 "testing"
)

func TestRLNodalStepAndSteadyState(t *testing.T) {
 // 10 V -> 100 ohm -> L=1 H -> ground. First sample current
 // is dt*V/(L+R*dt) with zero initial current.
 n,err:=NewRLNodalNetwork(1000,[]ResistiveBranch{{A:"source",B:"mid",Ohms:100}},[]InductiveBranch{{A:"mid",B:"",Henry:1}})
 if err!=nil {t.Fatal(err)}
 out,err:=n.Next(map[string]float64{"source":10})
 if err!=nil {t.Fatal(err)}
 expected:=10/1100.0
 if math.Abs(n.currents[0]-expected)>1e-10 {t.Fatalf("first current %g, want %g",n.currents[0],expected)}
 if math.Abs(out["mid"]-(10-100*expected))>1e-9 {t.Fatalf("first node voltage %g",out["mid"])}
 for i:=0;i<500;i++ {
  if _,err=n.Next(map[string]float64{"source":10});err!=nil {t.Fatal(err)}
 }
 if math.Abs(n.currents[0]-0.1)>1e-8 {t.Fatalf("DC steady current %g",n.currents[0])}
 n.Reset()
 if n.currents[0]!=0 {t.Fatal("reset retained current")}
}

func TestRLNodalRejectsFloatingNetworkWithoutStateMutation(t *testing.T) {
 n,err:=NewRLNodalNetwork(1000,nil,[]InductiveBranch{{A:"a",B:"b",Henry:1}})
 if err!=nil {t.Fatal(err)}
 if _,err=n.Next(nil);err==nil {t.Fatal("accepted floating RL network")}
 if n.currents[0]!=0 {t.Fatal("failed sample changed state")}
}
