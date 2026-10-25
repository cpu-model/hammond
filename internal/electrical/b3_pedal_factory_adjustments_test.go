package electrical

import "testing"

// Note 7 concerns R61 in the percussion circuitry, not pedal R81.
func TestB3PedalFactoryAdjustmentsExcludeNonPedalR61(t *testing.T) {
 if got:=B3PedalFactoryAdjustments(); len(got)!=0 {
  t.Fatalf("unexpected pedal-specific factory adjustments: %+v",got)
 }
}
