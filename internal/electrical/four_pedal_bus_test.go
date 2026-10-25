package electrical

import "testing"

func TestFourPedalBusIndependentLoading(t *testing.T) {
 var config [4]BusConfig
 for i:=range config {
  config[i]=BusConfig{
   SourceResistanceOhm:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental},
   LoadResistanceOhm:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental},
  }
 }
 n,err:=NewFourPedalBusNetwork(config)
 if err!=nil { t.Fatal(err) }
 n.BeginSample()
 if err:=n.AddClosedContact(0,2);err!=nil { t.Fatal(err) }
 if err:=n.AddClosedContact(0,4);err!=nil { t.Fatal(err) }
 if err:=n.AddClosedContact(3,-2);err!=nil { t.Fatal(err) }
 v:=n.Voltages()
 if v[0]!=2||v[1]!=0||v[2]!=0||v[3]!=-1 { t.Fatalf("unexpected four bus voltages: %v",v) }
 n.BeginSample()
 if v:=n.Voltages();v!=([4]float64{}) { t.Fatalf("sample state not reset: %v",v) }
}

func TestFourPedalBusRejectsBadInput(t *testing.T) {
 var config [4]BusConfig
 for i:=range config {
  config[i]=BusConfig{
   SourceResistanceOhm:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental},
   LoadResistanceOhm:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceExperimental},
  }
 }
 config[2].LoadResistanceOhm.Value=-1
 if _,err:=NewFourPedalBusNetwork(config);err==nil { t.Fatal("accepted negative load") }
 config[2].LoadResistanceOhm.Value=1000
 n,err:=NewFourPedalBusNetwork(config)
 if err!=nil { t.Fatal(err) }
 n.BeginSample()
 for _,i:=range []int{-1,4} {
  if err:=n.AddClosedContact(i,1);err==nil { t.Fatalf("accepted bus %d",i) }
 }
}
