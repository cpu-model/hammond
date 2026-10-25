package electrical

import (
 "math"
 "testing"
)

// The backward-Euler RL solver should converge toward the analytical
// frequency response as sample rate rises. This is a synthetic circuit
// validation, not a claim about the B3 L20 filter.
func TestSyntheticRLFrequencyResponseAgainstTheory(t *testing.T) {
 const rate=48000.0
 const r=1000.0
 const l=0.1
 for _,freq:=range []float64{50,200,1000,5000} {
  n,err:=NewRLNodalNetwork(rate,
   []ResistiveBranch{{A:"out",B:"",Ohms:r}},
   []InductiveBranch{{A:"in",B:"out",Henry:l}})
  if err!=nil {t.Fatal(err)}
  // Integer-period frequencies ensure the measurement spans full cycles.
  const samples=48000
  var sum float64
  for i:=0;i<samples;i++ {
   v:=math.Sin(2*math.Pi*freq*float64(i)/rate)
   result,err:=n.Next(map[string]float64{"in":v,"":0})
   if err!=nil {t.Fatal(err)}
   if i>=samples/2 {sum+=result["out"]*result["out"]}
  }
  got:=math.Sqrt(2*sum/float64(samples/2))
  theoretical:=r/math.Hypot(r,2*math.Pi*freq*l)
  // Backward Euler has frequency-dependent damping at high frequencies.
  tolerance:=0.025
  if math.Abs(got-theoretical)>tolerance {
   t.Fatalf("%.0f Hz: got %.5f, theory %.5f",freq,got,theoretical)
  }
 }
}
