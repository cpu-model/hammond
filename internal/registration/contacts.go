package registration

import (
 "errors"
)

// ContactTiming is an explicitly configurable research model, not a
// measurement of B3 key-contact mechanics. Delays are in audio samples.
// A contact switches instantaneously at its scheduled sample.
type ContactTiming struct {
 Make [Harmonics]uint32
 Break [Harmonics]uint32
}

// TimedManual shares a Manual's routing and drawbar gains but independently
// schedules all nine contact closures and openings per key.
type TimedManual struct {
 manual *Manual
 timing ContactTiming
 requested [Keys]bool
 active [Keys][Harmonics]bool
 pending [Keys][Harmonics]bool
 target [Keys][Harmonics]bool
 due [Keys][Harmonics]uint64
 clock uint64
}

func NewTimedManual(route Routing, timing ContactTiming) (*TimedManual,error) {
 m,err:=NewManual(route)
 if err!=nil { return nil,err }
 return &TimedManual{manual:m,timing:timing},nil
}

func (m *TimedManual) SetDrawbar(index,position int) error {
 return m.manual.SetDrawbar(index,position)
}

// Key schedules transitions relative to the next rendered sample.
// Repeated identical requests have no effect. An opposite request
// supersedes any pending transition.
func (m *TimedManual) Key(key int,down bool) error {
 if key<0 || key>=Keys { return errors.New("key index outside 0..60") }
 if m.requested[key]==down { return nil }
 m.requested[key]=down
 for contact:=0;contact<Harmonics;contact++ {
  delay:=m.timing.Break[contact]
  if down { delay=m.timing.Make[contact] }
  m.target[key][contact]=down
  m.pending[key][contact]=true
  m.due[key][contact]=m.clock+uint64(delay)
 }
 return nil
}

// AdvanceContacts applies transitions due at this sample and advances
// the contact clock exactly once. It does not inspect generator signals.
func (m *TimedManual) AdvanceContacts() {
 for key:=0;key<Keys;key++ {
  for contact:=0;contact<Harmonics;contact++ {
   if m.pending[key][contact] && m.clock>=m.due[key][contact] {
    m.active[key][contact]=m.target[key][contact]
    m.pending[key][contact]=false
   }
  }
 }
 if m.clock < ^uint64(0) { m.clock++ }
}

// Mix advances contact events once, then mixes the shared generator
// frame. It retains the original sample timing and allocates nothing.
func (m *TimedManual) Mix(frame []float64) float64 {
 if len(frame)<91 { panic("registration: generator frame too short") }
 m.AdvanceContacts()
 var sum float64
 for key:=0;key<Keys;key++ {
  for contact:=0;contact<Harmonics;contact++ {
   if !m.active[key][contact] { continue }
   wheel:=m.manual.route[key][contact]
   if wheel!=0 { sum+=m.manual.level[contact]*frame[wheel-1] }
  }
 }
 return sum
}

// ForEachActiveContact visits active, connected key contacts without
// changing the contact clock. Call AdvanceContacts or Mix once per sample before visiting.
// wheel is a one-based generator index and bus is a zero-based harmonic
// conductor index. Drawbar gain is deliberately not applied.
func (m *TimedManual) ForEachActiveContact(visit func(bus int,wheel int)) {
 for key:=0;key<Keys;key++ {
  for bus:=0;bus<Harmonics;bus++ {
   if !m.active[key][bus] { continue }
   wheel:=m.manual.route[key][bus]
   if wheel!=0 { visit(bus,int(wheel)) }
  }
 }
}
