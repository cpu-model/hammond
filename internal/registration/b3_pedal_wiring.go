package registration

import (
 "fmt"
 "strings"
)

// PedalWiringEvidence binds an explicitly transcribed physical contact
// permutation and busbar assignment to a cited wiring source. It does
// not independently verify that the citation describes factory B3 wiring.
type PedalWiringEvidence struct {
 ContactOrder PedalContactOrder
 Busbars PedalBusRouting
 Source string
 // Verified is set only after independent comparison with the original
 // B3/C3 physical contact and busbar wiring diagram.
 Verified bool
}

// CompileB3PedalWiring produces a validated contact-level routing and
// busbar mapping only when the caller supplies source-identified wiring. A nonblank source
// is an audit label, not proof of correctness.
// The frequency values are taken from B3PedalHarmonics, while the
// physical order and busbars must be established independently.
func CompileB3PedalWiring(e PedalWiringEvidence) (PedalRouting, PedalBusRouting, error) {
 if strings.TrimSpace(e.Source) == "" {
  return PedalRouting{}, PedalBusRouting{}, fmt.Errorf("pedal wiring source is required")
 }
 if !e.Verified {
  return PedalRouting{}, PedalBusRouting{}, fmt.Errorf("B3 physical wiring has not been independently verified")
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
