package electrical

import (
 "math"
 "testing"
)

func experimentalOhms(v float64) Parameter {
 return Parameter{Value:v,Unit:"ohm",Provenance:ProvenanceExperimental}
}

func TestResistiveContactDivider(t *testing.T) {
 c:=ResistiveContact{SourceOhms:experimentalOhms(1000),LoadOhms:experimentalOhms(3000)}
 if err:=c.Validate();err!=nil { t.Fatal(err) }
 if got:=c.Voltage(2,false);got!=0 { t.Fatalf("open contact: %g",got) }
 if got:=c.Voltage(2,true);math.Abs(got-1.5)>1e-12 { t.Fatalf("closed contact: %g",got) }
 if got:=c.Voltage(-2,true);math.Abs(got+1.5)>1e-12 { t.Fatalf("negative source: %g",got) }
}

func TestResistiveContactRequiresEvidenceAndUnits(t *testing.T) {
 c:=ResistiveContact{SourceOhms:experimentalOhms(1000),LoadOhms:experimentalOhms(3000)}
 c.SourceOhms.Provenance=ProvenanceSchematic
 if err:=c.Validate();err==nil { t.Fatal("accepted undocumented schematic value") }
 c.SourceOhms.Source="test fixture"
 c.LoadOhms.Unit="Hz"
 if err:=c.Validate();err==nil { t.Fatal("accepted wrong unit") }
 c.LoadOhms.Unit="ohm"
 c.LoadOhms.Value=0
 if err:=c.Validate();err==nil { t.Fatal("accepted zero load") }
}

func TestConductance(t *testing.T) {
 g,err:=Conductance(experimentalOhms(250))
 if err!=nil { t.Fatal(err) }
 if math.Abs(g-0.004)>1e-15 { t.Fatalf("conductance %g",g) }
 if _,err:=Conductance(experimentalOhms(0));err==nil { t.Fatal("accepted zero resistance") }
}
