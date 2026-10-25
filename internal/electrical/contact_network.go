package electrical

import (
 "errors"
 "math"
)

// ResistiveContact represents a purely resistive, static source/contact
// divider. It is NOT a calibrated Hammond B3 resistance-wire network.
// Open contacts are electrically disconnected in this simplified model.
type ResistiveContact struct {
 SourceOhms Parameter
 LoadOhms Parameter
}

func (c ResistiveContact) Validate() error {
 for _,p:=range []Parameter{c.SourceOhms,c.LoadOhms} {
  if err:=p.Validate();err!=nil { return err }
  if p.Unit!="ohm"||p.Value<=0 {
   return errors.New("resistances must be positive ohms")
  }
 }
 return nil
}

// Voltage returns the bus voltage with an ideal open/closed contact.
// It uses a simple unloaded source and single load; multiple contacts,
// pickup impedance, busbar capacitance and transformer coupling are
// deliberately excluded.
func (c ResistiveContact) Voltage(source float64,closed bool) float64 {
 if !closed { return 0 }
 return source*c.LoadOhms.Value/(c.SourceOhms.Value+c.LoadOhms.Value)
}

// Conductance is the inverse of a strictly positive resistance.
func Conductance(p Parameter) (float64,error) {
 if err:=p.Validate();err!=nil { return 0,err }
 if p.Unit!="ohm"||p.Value<=0 { return 0,errors.New("positive resistance in ohms required") }
 g:=1/p.Value
 if math.IsInf(g,0)||g==0 { return 0,errors.New("resistance outside representable conductance range") }
 return g,nil
}
