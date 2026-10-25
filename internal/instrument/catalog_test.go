package instrument

import (
 "strings"
 "testing"
)

func fixture() []Wheel {
 w := make([]Wheel, TonewheelCount)
 for i := range w { w[i] = Wheel{Number:i+1, FrequencyHz:float64(i+100), Source:"TEST ONLY"} }
 return w
}

func TestCatalog(t *testing.T) {
 c, err := NewCatalog(fixture())
 if err != nil { t.Fatal(err) }
 f, err := c.Frequencies()
 if err != nil || len(f) != TonewheelCount || f[0] != 100 || f[90] != 190 { t.Fatal("incorrect catalog") }
}

func TestRejectMissing(t *testing.T) {
 if _, err := NewCatalog(fixture()[:90]); err == nil { t.Fatal("missing wheel accepted") }
}

func TestRejectDuplicate(t *testing.T) {
 w := fixture(); w[90].Number = 1
 if _, err := NewCatalog(w); err == nil { t.Fatal("duplicate accepted") }
}

func TestRejectUntraceable(t *testing.T) {
 w := fixture(); w[0].Source = ""
 if _, err := NewCatalog(w); err == nil || !strings.Contains(err.Error(),"source") { t.Fatal("untraceable wheel accepted") }
}

func TestRejectZeroValue(t *testing.T) {
 var c Catalog
 if _, err := c.Frequencies(); err == nil { t.Fatal("zero value accepted") }
}
