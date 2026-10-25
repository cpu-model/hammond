package electrical

import (
 "strings"
 "testing"
)

func TestManualOnlyCircuitCoverageCannotClaimFullT1T2(t *testing.T) {
 c:=ExperimentalManualOnlyCoverage()
 if !c.UpperManual||!c.LowerManual||c.Pedals { t.Fatalf("incorrect manual-only scope: %+v",c) }
 if err:=c.ValidateFullTransformerCoverage();err==nil||!strings.Contains(err.Error(),"pedal") {
  t.Fatalf("missing T1 pedal contribution was not detected: %v",err)
 }
 c.Pedals=true
 if err:=c.ValidateFullTransformerCoverage();err!=nil { t.Fatal(err) }
 c.UpperManual=false
 if err:=c.ValidateFullTransformerCoverage();err==nil { t.Fatal("accepted missing upper manual") }
 c.UpperManual=true
 c.LowerManual=false
 if err:=c.ValidateFullTransformerCoverage();err==nil { t.Fatal("accepted missing lower manual") }
}
