package instrument

import (
 "errors"
 "fmt"
 "math"
)

const TonewheelCount = 91

// Wheel describes one generator output. FrequencyHz must come from a
// documented source, never from an equal-temperament approximation.
// Source identifies the source document/table for traceability.
type Wheel struct {
 Number int
 FrequencyHz float64
 Source string
}

// Catalog is a fully validated, immutable-by-convention wheel specification.
// The zero value is not a valid catalog.
type Catalog struct { wheels [TonewheelCount]Wheel }

// NewCatalog rejects incomplete, duplicated, and untraceable datasets.
func NewCatalog(wheels []Wheel) (Catalog, error) {
 var c Catalog
 if len(wheels) != TonewheelCount { return c, fmt.Errorf("expected %d wheels, got %d", TonewheelCount, len(wheels)) }
 var seen [TonewheelCount]bool
 for _, w := range wheels {
  if w.Number < 1 || w.Number > TonewheelCount { return Catalog{}, fmt.Errorf("invalid wheel number %d", w.Number) }
  if seen[w.Number-1] { return Catalog{}, fmt.Errorf("duplicate wheel number %d", w.Number) }
  if w.FrequencyHz <= 0 || math.IsNaN(w.FrequencyHz) || math.IsInf(w.FrequencyHz, 0) { return Catalog{}, fmt.Errorf("invalid frequency for wheel %d", w.Number) }
  if w.Source == "" { return Catalog{}, fmt.Errorf("missing source for wheel %d", w.Number) }
  seen[w.Number-1] = true
  c.wheels[w.Number-1] = w
 }
 for i, ok := range seen { if !ok { return Catalog{}, fmt.Errorf("missing wheel %d", i+1) } }
 return c, nil
}

func (c Catalog) Frequencies() ([]float64, error) {
 f := make([]float64, TonewheelCount)
 for i, w := range c.wheels {
  if w.Number != i+1 { return nil, errors.New("uninitialized catalog") }
  f[i] = w.FrequencyHz
 }
 return f, nil
}
