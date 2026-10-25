package electrical

import (
 "math"
 "testing"
)

func TestNineBusNetworkIsolationAndMultipleContacts(t *testing.T) {
 var cfg [9]BusConfig
 for i:=range cfg {
  cfg[i]=BusConfig{SourceResistanceOhm:busExperimentalOhms(1000),LoadResistanceOhm:busExperimentalOhms(1000)}
 }
 n,err:=NewNineBusNetwork(cfg)
 if err!=nil { t.Fatal(err) }
 n.BeginSample()
 if err:=n.AddClosedContact(0,6);err!=nil { t.Fatal(err) }
 if err:=n.AddClosedContact(0,0);err!=nil { t.Fatal(err) }
 if err:=n.AddClosedContact(8,4);err!=nil { t.Fatal(err) }
 v:=n.Voltages()
 if math.Abs(v[0]-2)>1e-12||math.Abs(v[8]-2)>1e-12 {
  t.Fatalf("unexpected bus voltages: %v",v)
 }
 for i:=1;i<8;i++ { if v[i]!=0 { t.Fatalf("bus %d leaked: %g",i,v[i]) } }
 n.BeginSample()
 for i,x:=range n.Voltages() { if x!=0 { t.Fatalf("bus %d not reset",i) } }
}

func TestNineBusNetworkRejectsInvalidConfigurationAndBus(t *testing.T) {
 var cfg [9]BusConfig
 for i:=range cfg {
  cfg[i]=BusConfig{SourceResistanceOhm:busExperimentalOhms(1000),LoadResistanceOhm:busExperimentalOhms(1000)}
 }
 cfg[4].LoadResistanceOhm.Value=0
 if _,err:=NewNineBusNetwork(cfg);err==nil { t.Fatal("accepted invalid bus") }
 cfg[4].LoadResistanceOhm.Value=1000
 n,err:=NewNineBusNetwork(cfg)
 if err!=nil { t.Fatal(err) }
 n.BeginSample()
 if err:=n.AddClosedContact(9,1);err==nil { t.Fatal("accepted invalid index") }
}

func TestNineBusNetworkNoAllocations(t *testing.T) {
 var cfg [9]BusConfig
 for i:=range cfg {
  cfg[i]=BusConfig{SourceResistanceOhm:busExperimentalOhms(1000),LoadResistanceOhm:busExperimentalOhms(1000)}
 }
 n,err:=NewNineBusNetwork(cfg)
 if err!=nil { t.Fatal(err) }
 if count:=testing.AllocsPerRun(100,func(){
  n.BeginSample()
  _=n.AddClosedContact(0,1)
  _=n.Voltages()
 });count!=0 { t.Fatalf("allocations: %v",count) }
}
