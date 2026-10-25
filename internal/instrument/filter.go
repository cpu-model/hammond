package instrument

import "fmt"

// FilterClass describes the generator's documented electrical output network.
// This is descriptive metadata, not yet a circuit/DSP implementation.
type FilterClass uint8

const (
 ResistiveShunt FilterClass = iota + 1
 ReactorOnly
 ReactorCapacitor255nF
 ReactorCapacitor105nF
)

// FilterForFrequencyNumber maps original generator terminal numbers 1..91.
// Source: Hammond B3/C3 Service Manual, section 5, paragraph 5-10.
// The manual prints capacitor values as 0.255 and 0.105 microfarads.
func FilterForFrequencyNumber(number int) (FilterClass, error) {
 if number < 1 || number > TonewheelCount {
  return 0, fmt.Errorf("invalid frequency number %d", number)
 }
 switch {
 case number <= 43:
  return ResistiveShunt, nil
 case number <= 48:
  return ReactorOnly, nil
 case number <= 54:
  return ReactorCapacitor255nF, nil
 default:
  return ReactorCapacitor105nF, nil
 }
}
