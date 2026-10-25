package electrical

import (
 "math"
 "testing"
)

func pedalRLFixture() PedalRLConfig {
 return PedalRLConfig{
  InductanceHenry:Parameter{Value:0.5,Unit:"H",Provenance:ProvenanceExperimental},
  SeriesResistanceOhm:Parameter{Value:100,Unit:"ohm",Provenance:ProvenanceExperimental},
  LoadResistanceOhm:Parameter{Value:100,Unit:"ohm",Provenance:ProvenanceExperimental},
 }
}

func TestPedalRLStepAndReset(t *testing.T) {
 s,err:=NewPedalRLStage(48000,pedalRLFixture())
 if err!=nil { t.Fatal(err) }
 prev:=0.0
 for i:=0;i<10000;i++ {
  v,err:=s.Next(1)
  if err!=nil||v<prev||v>0.5 { t.Fatalf("unstable step at %d: %v %v",i,v,err) }
  prev=v
 }
 if math.Abs(prev-0.5)>1e-6 { t.Fatalf("incorrect DC gain: %v",prev) }
 s.Reset()
 v,err:=s.Next(0)
 if err!=nil||v!=0 { t.Fatalf("reset failed: %v %v",v,err) }
}

func TestPedalRLInvalidParametersAndInput(t *testing.T) {
 c:=pedalRLFixture()
 c.InductanceHenry.Unit="mH"
 if _,err:=NewPedalRLStage(48000,c);err==nil { t.Fatal("accepted wrong inductance unit") }
 c=pedalRLFixture()
 c.LoadResistanceOhm.Value=0
 if _,err:=NewPedalRLStage(48000,c);err==nil { t.Fatal("accepted zero load") }
 s,err:=NewPedalRLStage(48000,pedalRLFixture())
 if err!=nil { t.Fatal(err) }
 if _,err:=s.Next(math.NaN());err==nil { t.Fatal("accepted NaN") }
}

func TestPedalRLNoAllocations(t *testing.T) {
 s,err:=NewPedalRLStage(48000,pedalRLFixture())
 if err!=nil { t.Fatal(err) }
 if n:=testing.AllocsPerRun(100,func(){_,_ =s.Next(1)});n!=0 { t.Fatalf("allocations per sample: %v",n) }
}
