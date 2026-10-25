package registration

import "testing"

func TestPedalContactOrderRejectsDuplicatesAndOutOfRange(t *testing.T) {
 for _, order := range []PedalContactOrder{
  {0,1,2,3,4,5,6,6},
  {0,1,2,3,4,5,6,8},
 } {
  if _, err := PedalRoutingFromHarmonics(B3PedalHarmonics(), order); err == nil {
   t.Fatalf("accepted invalid contact order %v", order)
  }
 }
}

func TestPedalContactOrderPermutesWithoutInventingWheels(t *testing.T) {
 // Synthetic permutation: a test of the conversion, NOT B3 wiring evidence.
 order := PedalContactOrder{2,0,1,3,4,5,6,7}
 got, err := PedalRoutingFromHarmonics(B3PedalHarmonics(), order)
 if err != nil { t.Fatal(err) }
 source := B3PedalHarmonics()
 for key := 0; key < PedalKeys; key++ {
  for contact, harmonic := range order {
   if got[key][contact] != source[key][harmonic] {
    t.Fatalf("key %d contact %d: got %d want %d", key, contact, got[key][contact], source[key][harmonic])
   }
  }
 }
 if got[7][2] != 0 || got[11][2] != 0 {
  t.Fatal("missing third-harmonic contacts were filled")
 }
 state, err := NewPedalContactsState(got)
 if err != nil { t.Fatal(err) }
 if err := state.Key(7, true); err != nil { t.Fatal(err) }
 count := 0
 state.ForEachClosedContact(func(key, contact, wheel int) { count++ })
 if count != 7 { t.Fatalf("expected 7 connected contacts on pedal 8, got %d", count) }
}
