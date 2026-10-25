package engine

import (
 "testing"

 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

func TestDualFilteredRequiresProvenance(t *testing.T) {
 p:=func(v float64,source string) electrical.Parameter {
  return electrical.Parameter{Value:v,Unit:"Hz",Provenance:electrical.ProvenanceSchematic,Source:source}
 }
 cfg:=electrical.ManualResponseConfig{UpperCutoffHz:p(1200,""),LowerCutoffHz:p(1800,"")}
 args:=func(c electrical.ManualResponseConfig) error {
  _,err:=NewDualFilteredTimedWithConfig(48000,c,registration.Routing{},registration.Routing{},registration.ContactTiming{},registration.ContactTiming{})
  return err
 }
 if err:=args(cfg);err==nil { t.Fatal("accepted undocumented schematic cutoff") }
 cfg.UpperCutoffHz.Source="test reference only"
 if err:=args(cfg);err==nil { t.Fatal("accepted undocumented lower cutoff") }
 cfg.LowerCutoffHz.Source="test reference only"
 if err:=args(cfg);err!=nil { t.Fatal(err) }
}

func TestDualFilteredExplicitExperimentalParameters(t *testing.T) {
 cfg:=electrical.ManualResponseConfig{
  UpperCutoffHz:electrical.Parameter{Value:1200,Unit:"Hz",Provenance:electrical.ProvenanceExperimental},
  LowerCutoffHz:electrical.Parameter{Value:1800,Unit:"Hz",Provenance:electrical.ProvenanceExperimental},
 }
 if _,err:=NewDualFilteredTimedWithConfig(48000,cfg,registration.Routing{},registration.Routing{},registration.ContactTiming{},registration.ContactTiming{});err!=nil {
  t.Fatal(err)
 }
}
