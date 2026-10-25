package registration

import "math"

// DrawbarGain is a nominal amplitude ratio relative to stop 8.
// The 3 dB-per-step approximation is based on documented drawbar
// operation; it is not an electrical transformer simulation.
func DrawbarGain(position int) float64 {
 if position <= 0 { return 0 }
 if position >= 8 { return 1 }
 return math.Pow(10, -3*float64(8-position)/20)
}
