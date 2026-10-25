package generator

import (
 "math"
 "testing"
)

func TestContinuousPhase(t *testing.T) {
 b, err := New(96000, []float64{440})
 if err != nil { t.Fatal(err) }
 sample := make([]float64, 1)
 for n := 0; n < 96000; n++ {
  b.Next(sample)
  want := math.Sin(2 * math.Pi * 440 * float64(n) / 96000)
  if math.Abs(sample[0]-want) > 1e-9 { t.Fatalf("sample %d: got %.12f want %.12f", n, sample[0], want) }
 }
}

func TestNoAllocations(t *testing.T) {
 b, err := New(48000, []float64{220, 440})
 if err != nil { t.Fatal(err) }
 sample := make([]float64, 2)
 if n := testing.AllocsPerRun(100, func() { b.Next(sample) }); n != 0 {
  t.Fatalf("allocations per render: %v", n)
 }
}

func TestInvalidFrequency(t *testing.T) {
 if _, err := New(48000, []float64{24000}); err == nil { t.Fatal("expected Nyquist validation") }
}
