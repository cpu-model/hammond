package registration

import "fmt"

const (
 Keys = 61
 Harmonics = 9
)

// Routing maps each of the 61 keys and nine drawbar contacts to a
// generator frequency number (1..91). Zero represents an unconnected
// contact. Values must be transcribed from a documented wiring source;
// no equal-temperament or modulo foldback is assumed.
type Routing [Keys][Harmonics]uint8

// Manual is an explicitly configured, allocation-free contact mixer.
// This is a digital linear mix, not a simulation of resistive drawbar
// networks, contact bounce, crosstalk or expression.
type Manual struct {
 route Routing
 down [Keys]bool
 level [Harmonics]float64
}

func NewManual(route Routing) (*Manual,error) {
 for key, contacts := range route {
  for drawbar, wheel := range contacts {
   if wheel > 91 {
    return nil,fmt.Errorf("key %d drawbar %d: wheel %d outside 1..91",key+1,drawbar+1,wheel)
   }
  }
 }
 return &Manual{route:route},nil
}

func (m *Manual) Key(key int, down bool) error {
 if key<0 || key>=Keys { return fmt.Errorf("key index %d outside 0..60",key) }
 m.down[key]=down
 return nil
}

// SetDrawbar accepts an organ stop position from 0 (off) to 8 (full).
// The current linear gain is an explicitly provisional placeholder.
func (m *Manual) SetDrawbar(index, position int) error {
 if index<0 || index>=Harmonics { return fmt.Errorf("drawbar index %d outside 0..8",index) }
 if position<0 || position>8 { return fmt.Errorf("drawbar position %d outside 0..8",position) }
 m.level[index]=float64(position)/8
 return nil
}

// Mix adds the selected generator signals. It never advances generator
// phases; callers advance their single shared generator once per sample.
func (m *Manual) Mix(frame []float64) float64 {
 if len(frame)<91 { panic("registration: generator frame too short") }
 var sum float64
 for k, held:=range m.down {
  if !held { continue }
  for d,wheel:=range m.route[k] {
   if wheel!=0 && m.level[d]!=0 { sum+=m.level[d]*frame[wheel-1] }
  }
 }
 return sum
}
