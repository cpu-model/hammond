package instrument

import "testing"

func TestFilterBoundaries(t *testing.T) {
 cases := []struct{ number int; want FilterClass }{
  {1, ResistiveShunt}, {43, ResistiveShunt},
  {44, ReactorOnly}, {48, ReactorOnly},
  {49, ReactorCapacitor255nF}, {54, ReactorCapacitor255nF},
  {55, ReactorCapacitor105nF}, {91, ReactorCapacitor105nF},
 }
 for _, tc := range cases {
  got, err := FilterForFrequencyNumber(tc.number)
  if err != nil || got != tc.want { t.Fatalf("frequency %d: got %v, %v", tc.number, got, err) }
 }
 for _, n := range []int{0, 92, -1} {
  if _, err := FilterForFrequencyNumber(n); err == nil { t.Fatalf("accepted invalid frequency %d", n) }
 }
}
