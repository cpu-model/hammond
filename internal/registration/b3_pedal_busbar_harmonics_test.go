package registration

import "testing"

// This test checks the documented color/harmonic pairing independently
// of the physical contact permutation used by the implementation.
func TestB3PedalBusbarHarmonicGroups(t *testing.T) {
 order:=B3PedalContactOrderFigure7()
 buses:=B3PedalBusbarsHandyGuide()
 // Harmonic columns: fundamental, third, second, fourth,
 // sixth, eighth, tenth, twelfth.
 wantByHarmonic:=PedalBusRouting{1,1,2,2,3,3,4,4}
 seen:=[PedalContacts]bool{}
 for physical,harmonic:=range order {
  if harmonic>=PedalContacts {t.Fatalf("contact %d invalid harmonic %d",physical,harmonic)}
  if seen[harmonic] {t.Fatalf("harmonic %d mapped twice",harmonic)}
  seen[harmonic]=true
  if buses[physical]!=wantByHarmonic[harmonic] {
   t.Fatalf("physical contact %d harmonic %d on bus %d, want %d",
    physical,harmonic,buses[physical],wantByHarmonic[harmonic])
  }
 }
 for harmonic,ok:=range seen {
  if !ok {t.Fatalf("harmonic %d omitted",harmonic)}
 }
}
