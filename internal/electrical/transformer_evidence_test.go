package electrical

import (
 "strings"
 "testing"
)

func TestTransformerEvidenceRequiresEverySource(t *testing.T) {
 e:=TransformerEvidence{ID:TransformerT2}
 for _,expected:=range []string{"winding topology","tap topology","impedance","transfer"} {
  err:=e.ReadyForCircuitModel()
  if err==nil||!strings.Contains(err.Error(),expected) { t.Fatalf("expected %s: %v",expected,err) }
  switch expected {
  case "winding topology":e.WindingTopologySource="test winding reference"
  case "tap topology":e.TapTopologySource="test tap reference"
  case "impedance":e.ImpedanceSource="test impedance reference"
  case "transfer":e.TransferSource="test transfer reference"
  }
 }
 if err:=e.ReadyForCircuitModel();err!=nil { t.Fatal(err) }
}

func TestTransformerEvidencePairRejectsSwappedOrIncompleteSources(t *testing.T) {
 p:=TransformerEvidencePair{UpperT2:TransformerEvidence{ID:TransformerT1},LowerT1:TransformerEvidence{ID:TransformerT2}}
 if err:=p.ReadyForCircuitModel();err==nil { t.Fatal("accepted swapped transformers") }
 p.UpperT2.ID=TransformerT2
 p.LowerT1.ID=TransformerT1
 if err:=p.ReadyForCircuitModel();err==nil { t.Fatal("accepted missing physical sources") }
}
