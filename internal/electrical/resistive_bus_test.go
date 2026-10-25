package electrical

import (
 "math"
 "testing"
)

func busExperimentalOhms(v float64) Parameter {
 return Parameter{Value:v,Unit:"ohm",Provenance:ProvenanceExperimental}
}

func TestResistiveBusMultipleClosedContacts(t *testing.T) {
 b,err:=NewResistiveBus(BusConfig{
  SourceResistanceOhm:busExperimentalOhms(1000),
  LoadResistanceOhm:busExperimentalOhms(1000),
 })
 if err!=nil { t.Fatal(err) }
 b.BeginSample()
 if b.Voltage()!=0 { t.Fatal("open bus should be grounded") }
 if err:=b.AddSource(6);err!=nil { t.Fatal(err) }
 if got:=b.Voltage();math.Abs(got-3)>1e-12 { t.Fatalf("one contact: %g",got) }
 if err:=b.AddSource(0);err!=nil { t.Fatal(err) }
 if got:=b.Voltage();math.Abs(got-2)>1e-12 { t.Fatalf("two contacts: %g",got) }
 b.BeginSample()
 if b.Voltage()!=0 { t.Fatal("bus contributions leaked into next sample") }
}

func TestResistiveBusRejectsUnprovenSchematicValue(t *testing.T) {
 c:=BusConfig{
  SourceResistanceOhm:Parameter{Value:1000,Unit:"ohm",Provenance:ProvenanceSchematic},
  LoadResistanceOhm:busExperimentalOhms(1000),
 }
 if _,err:=NewResistiveBus(c);err==nil { t.Fatal("accepted unreferenced schematic resistance") }
 c.SourceResistanceOhm=busExperimentalOhms(0)
 if _,err:=NewResistiveBus(c);err==nil { t.Fatal("accepted zero resistance") }
}

func TestResistiveBusNoAllocations(t *testing.T) {
 b,err:=NewResistiveBus(BusConfig{busExperimentalOhms(1000),busExperimentalOhms(1000)})
 if err!=nil { t.Fatal(err) }
 if n:=testing.AllocsPerRun(100,func(){
  b.BeginSample()
  _=b.AddSource(1)
  _=b.AddSource(2)
  _=b.Voltage()
 });n!=0 { t.Fatalf("allocations: %v",n) }
}
