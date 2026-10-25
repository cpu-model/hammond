package registration

import "fmt"

// B3PedalSwitch couples a validated contact-frequency routing to its
// physical busbar assignments. Construct it only after transcribing and
// checking the B3 wiring chart. This model contains no analog filtering.
type B3PedalSwitch struct {
 contacts *PedalContactsState
 busbars PedalBusRouting
 source string
}

// NewB3PedalSwitch compiles the harmonic-frequency table into a
// switch-level routing, requiring an explicit physical wiring mapping.
func NewB3PedalSwitch(e PedalWiringEvidence) (*B3PedalSwitch, error) {
 routing, busbars, err := CompileB3PedalWiring(e)
 if err != nil { return nil, err }
 contacts, err := NewPedalContactsState(routing)
 if err != nil { return nil, err }
 return &B3PedalSwitch{contacts:contacts,busbars:busbars,source:e.Source},nil
}

func (p *B3PedalSwitch) Key(key int, down bool) error {
 if p==nil { return fmt.Errorf("nil B3 pedal switch") }
 return p.contacts.Key(key,down)
}

// ForEachClosedBusContact emits zero-based busbar indices and frequency
// numbers for connected contacts, preserving Table III disconnections.
func (p *B3PedalSwitch) ForEachClosedBusContact(visit func(key,contact,busbar,wheel int)) error {
 if p==nil { return fmt.Errorf("nil B3 pedal switch") }
 if visit==nil { return fmt.Errorf("nil pedal contact visitor") }
 return p.contacts.ForEachClosedBusContact(p.busbars,visit)
}

func (p *B3PedalSwitch) Source() string {
 if p==nil { return "" }
 return p.source
}
