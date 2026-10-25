package electrical

// ManualPaths preserves independent upper and lower electrical signals.
// It is a topology boundary, not a calibrated transformer or preamp.
type ManualPaths struct {
 Upper float64
 Lower float64
}

// Sum is intended only for diagnostic/offline listening. A real B3
// uses separate upper and lower transformer circuits before its preamp.
func (p ManualPaths) Sum() float64 { return p.Upper+p.Lower }
