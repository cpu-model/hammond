package instrument

import "fmt"

// GearPair records the drive and driven gear tooth counts for a speed group.
// Source: https://www.goodeveca.net/RotorOrgan/ToneWheelSpec.html
type GearPair struct { Drive, Driven uint16 }

var speedGears = [12]GearPair{
 {85,104}, {71,82}, {67,73}, {105,108},
 {103,100}, {84,77}, {74,64}, {98,80},
 {96,74}, {88,64}, {67,46}, {108,70},
}

// MechanicalFrequency returns the documented nominal US 60 Hz frequency
// for a generator terminal number 1..91. It is a pitch-only model, not
// a complete mechanical, pickup or electrical simulation.
func MechanicalFrequency(number int) (float64, error) {
 if number < 1 || number > TonewheelCount {
  return 0, fmt.Errorf("invalid generator frequency number %d", number)
 }
 var teeth uint16
 var group int
 if number <= 84 {
  octave := (number-1)/12
  group = (number-1)%12
  teeth = uint16(2 << octave)
 } else {
  // 192-tooth upper wheels C7..F#7 use the gear ratios F..B.
  // See HammondWiki Gear Ratio and Goodeve's tonewheel table.
  teeth = 192
  group = number - 85 + 5
 }
 gear := speedGears[group]
 const driveRPM = 1200.0
 return (driveRPM/60.0)*float64(gear.Drive)/float64(gear.Driven)*float64(teeth), nil
}

// MechanicalFrequencies returns the complete nominal 91-frequency array.
func MechanicalFrequencies() [TonewheelCount]float64 {
 var frequencies [TonewheelCount]float64
 for i := range frequencies {
  frequencies[i], _ = MechanicalFrequency(i+1)
 }
 return frequencies
}

// Deprecated: use MechanicalFrequency. Retained for existing callers.
func VerifiedMechanicalFrequency(number int) (float64, error) {
 return MechanicalFrequency(number)
}

// Deprecated: use MechanicalFrequencies. Retained for existing callers.
func VerifiedMechanicalFrequencies() [84]float64 {
 var frequencies [84]float64
 for i := range frequencies {
  frequencies[i], _ = MechanicalFrequency(i+1)
 }
 return frequencies
}
