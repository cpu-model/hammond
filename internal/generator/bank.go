package generator

import (
 "errors"
 "math"
)

// Bank models free-running shared oscillator signals.
// Frequencies are supplied by the caller; no unverified B3 tuning is assumed.
type Bank struct {
 phase []float64
 step []float64
}

func New(sampleRate float64, frequencies []float64) (*Bank, error) {
 if sampleRate <= 0 || math.IsNaN(sampleRate) || math.IsInf(sampleRate, 0) { return nil, errors.New("invalid sample rate") }
 b := &Bank{phase: make([]float64, len(frequencies)), step: make([]float64, len(frequencies))}
 for i, f := range frequencies {
  if f <= 0 || f >= sampleRate/2 || math.IsNaN(f) || math.IsInf(f, 0) { return nil, errors.New("invalid oscillator frequency") }
  b.step[i] = f / sampleRate
 }
 return b, nil
}

func (b *Bank) Count() int { return len(b.phase) }

// Next renders and advances every oscillator exactly once. No allocations.
func (b *Bank) Next(dst []float64) {
 if len(dst) < len(b.phase) { panic("generator: destination too short") }
 for i := range b.phase {
  dst[i] = math.Sin(2 * math.Pi * b.phase[i])
  p := b.phase[i] + b.step[i]
  if p >= 1 { p -= 1 }
  b.phase[i] = p
 }
}
