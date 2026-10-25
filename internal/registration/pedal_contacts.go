package registration

import "fmt"

const (
 PedalKeys=25
 PedalContacts=8
)

// PedalRouting maps each of the eight physical contacts under each
// pedal key to a tonewheel (1..91); zero denotes no connection.
// The actual B3 wiring chart must be transcribed before populating it.
type PedalRouting [PedalKeys][PedalContacts]uint8

func (r PedalRouting) Validate() error {
 for key,contacts:=range r {
  for contact,wheel:=range contacts {
   if wheel>91 { return fmt.Errorf("pedal key %d contact %d: wheel %d outside 0..91",key,contact,wheel) }
  }
 }
 return nil
}

// PedalContactsState is a switch-only model. A key actuates eight
// contacts; no contact bounce or make/break staggering is implied.
type PedalContactsState struct {
 routing PedalRouting
 pressed [PedalKeys]bool
}

func NewPedalContactsState(r PedalRouting) (*PedalContactsState,error) {
 if err:=r.Validate();err!=nil { return nil,err }
 return &PedalContactsState{routing:r},nil
}

func (p *PedalContactsState) Key(key int,down bool) error {
 if key<0||key>=PedalKeys { return fmt.Errorf("pedal key %d outside 0..24",key) }
 p.pressed[key]=down
 return nil
}

// ForEachClosedContact visits the connected tonewheel references.
// Contacts with routing zero are electrically unspecified and skipped.
func (p *PedalContactsState) ForEachClosedContact(visit func(key,contact,wheel int)) {
 for key,down:=range p.pressed {
  if !down { continue }
  for contact,wheel:=range p.routing[key] {
   if wheel!=0 { visit(key,contact,int(wheel)) }
  }
 }
}
