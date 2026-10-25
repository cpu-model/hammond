package registration

import "fmt"

const PedalBusbars = 4

// PedalBusRouting maps each of eight pedal contacts to one of four
// busbars (1..4). Zero means not connected. The contact-to-busbar
// mapping must be supplied from wiring evidence; no B3 mapping is
// assumed here.
type PedalBusRouting [PedalContacts]uint8

func (r PedalBusRouting) Validate() error {
 for i,b:=range r {
  if b>PedalBusbars { return fmt.Errorf("pedal contact %d: busbar %d outside 0..4",i,b) }
 }
 return nil
}

// ForEachClosedBusContact exposes the busbar and generator wheel
// corresponding to each active pedal contact. This is switch
// connectivity only: no resistors, reactor or bus mixing.
func (p *PedalContactsState) ForEachClosedBusContact(bus PedalBusRouting,visit func(key,contact,busbar,wheel int)) error {
 if err:=bus.Validate();err!=nil { return err }
 p.ForEachClosedContact(func(key,contact,wheel int) {
  b:=bus[contact]
  if b!=0 { visit(key,contact,int(b)-1,wheel) }
 })
 return nil
}
