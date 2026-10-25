package electrical

import "testing"

func TestB3PedalFactoryAdjustmentIsR81(t *testing.T) {
 adjustments := B3PedalFactoryAdjustments()
 if len(adjustments) != 1 {
  t.Fatalf("factory pedal adjustments: got %d, want 1",len(adjustments))
 }
 a := adjustments[0]
 if a.Designator != "R81" {
  t.Fatalf("factory adjustment refers to %q, want R81 (schematic note 7)",a.Designator)
 }
 if a.Instruction != "SELECTED AT TEST TO BALANCE THE CHANNELS" {
  t.Fatalf("unexpected adjustment instruction: %q",a.Instruction)
 }
 if a.Source == "" {
  t.Fatal("factory adjustment requires source")
 }
}
