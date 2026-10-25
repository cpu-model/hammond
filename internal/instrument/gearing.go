package instrument

import "fmt"

// GearPair records the documented drive and driven tooth counts for one
// chromatic speed group. Data: Goodeve Hammond Tone Wheels reference table.
// Reference shaft speed for the US 60 Hz generator is 1200 RPM.
type GearPair struct { Drive, Driven uint16 }

var speedGears = [12]GearPair{
 {85,104}, {71,82}, {67,73}, {105,108},
 {103,100}, {84,77}, {74,64}, {98,80},
 {96,74}, {88,64}, {67,46}, {108,70},
}

// VerifiedMechanicalFrequency computes a terminal frequency for numbers
// 1..84. The highest seven terminals (192-tooth wheels) are intentionally
// excluded pending independent verification of their non-octave mapping.
func VerifiedMechanicalFrequency(number int) (float64, error) {
 if number < 1 || number > 84 {
  return 0, fmt.Errorf("frequency %d outside verified mechanical mapping 1..84", number)
 }
 octave := (number-1)/12
 semitone := (number-1)%12
 gear := speedGears[semitone]
 teeth := uint16(2 << octave)
 const driveRPM = 1200.0
 return (driveRPM/60.0)*float64(gear.Drive)/float64(gear.Driven)*float64(teeth), nil
}

// VerifiedMechanicalFrequencies returns the first 84 generator outputs.
// It is not a complete B3 specification.
func VerifiedMechanicalFrequencies() [84]float64 {
 var frequencies [84]float64
 for i := range frequencies { frequencies[i], _ = VerifiedMechanicalFrequency(i+1) }
 return frequencies
}
