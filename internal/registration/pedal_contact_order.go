package registration

import "fmt"

// PedalContactOrder maps physical contact index to a harmonic column
// in B3PedalHarmonics. This order must come from an independently
// verified pedal switch wiring chart, not from Table III alone.
// A permutation is required: no harmonic may be silently duplicated.
type PedalContactOrder [PedalContacts]uint8

func (order PedalContactOrder) Validate() error {
 var seen [PedalContacts]bool
 for contact, harmonic := range order {
  if harmonic >= PedalContacts {
   return fmt.Errorf("contact %d: harmonic column %d outside 0..7", contact, harmonic)
  }
  if seen[harmonic] {
   return fmt.Errorf("contact %d: duplicate harmonic column %d", contact, harmonic)
  }
  seen[harmonic] = true
 }
 return nil
}

// PedalRoutingFromHarmonics converts frequency-table column order into
// physical contact order. Callers must explicitly supply the latter;
// the published harmonic table does not establish switch wiring order.
func PedalRoutingFromHarmonics(harmonics PedalRouting, order PedalContactOrder) (PedalRouting, error) {
 if err := harmonics.Validate(); err != nil { return PedalRouting{}, err }
 if err := order.Validate(); err != nil { return PedalRouting{}, err }
 var routing PedalRouting
 for key, row := range harmonics {
  for contact, harmonic := range order {
   routing[key][contact] = row[harmonic]
  }
 }
 return routing, nil
}
