package registration

import "fmt"

// PedalWiringEvidence binds an explicitly transcribed physical contact
// permutation and busbar assignment to a cited wiring source. It does
// not assert that a particular permutation is the B3 factory wiring.
type PedalWiringEvidence struct {
 ContactOrder PedalContactOrder
 Busbars PedalBusRouting
 Source string
}

// CompileB3PedalWiring produces a validated contact-level routing and
// busbar mapping only when the caller supplies source-identified wiring.
// The frequency values are taken from B3PedalHarmonics, while the
// physical order and busbars must be established independently.
func CompileB3PedalWiring(e PedalWiringEvidence) (PedalRouting, PedalBusRouting, error) {
 if e.Source == "" {
  return PedalRouting{}, PedalBusRouting{}, fmt.Errorf("pedal wiring source is required")
 }
 if err := e.Busbars.Validate(); err != nil {
  return PedalRouting{}, PedalBusRouting{}, err
 }
 routing, err := PedalRoutingFromHarmonics(B3PedalHarmonics(), e.ContactOrder)
 if err != nil { return PedalRouting{}, PedalBusRouting{}, err }
 for contact, bus := range e.Busbars {
  if bus == 0 {
   return PedalRouting{}, PedalBusRouting{}, fmt.Errorf("contact %d has no verified busbar assignment", contact)
  }
 }
 return routing, e.Busbars, nil
}
