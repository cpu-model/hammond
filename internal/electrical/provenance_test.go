package electrical

import (
 "math"
 "testing"
)

func TestParameterProvenance(t *testing.T) {
 cases:=[]struct{p Parameter; valid bool}{
  {Parameter{Value:1000,Unit:"Hz",Provenance:ProvenanceExperimental},true},
  {Parameter{Value:1000,Unit:"Hz",Provenance:ProvenanceSchematic,Source:"manual page 42"},true},
  {Parameter{Value:1000,Unit:"Hz",Provenance:ProvenanceMeasurement,Source:"measurement log"},true},
  {Parameter{Value:1000,Unit:"Hz",Provenance:ProvenanceSchematic},false},
  {Parameter{Value:1000,Unit:"Hz",Provenance:ProvenanceMeasurement},false},
  {Parameter{Value:1000,Unit:"Hz",Provenance:"unknown"},false},
  {Parameter{Value:math.NaN(),Unit:"Hz",Provenance:ProvenanceExperimental},false},
  {Parameter{Value:1000,Provenance:ProvenanceExperimental},false},
 }
 for i,c:=range cases {
  if got:=c.p.Validate()==nil;got!=c.valid { t.Fatalf("case %d: valid=%v expected=%v",i,got,c.valid) }
 }
}

func TestManualResponseConfig(t *testing.T) {
 p:=func(v float64) Parameter { return Parameter{Value:v,Unit:"Hz",Provenance:ProvenanceExperimental} }
 config:=ManualResponseConfig{UpperCutoffHz:p(1000),LowerCutoffHz:p(4000)}
 if err:=config.Validate(48000);err!=nil { t.Fatal(err) }
 config.LowerCutoffHz.Value=24000
 if err:=config.Validate(48000);err==nil { t.Fatal("expected Nyquist error") }
 config.LowerCutoffHz=p(4000)
 if err:=config.Validate(math.NaN());err==nil { t.Fatal("expected invalid sample rate") }
}
