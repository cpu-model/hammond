package electrical

import (
 "math"
 "testing"
)

// Synthetic frequencies and inductance exercise the time-domain reactor
// solver. These numbers are NOT Hammond L20 component specifications.
func TestB3PedalFilterSyntheticFrequencyResponse(t *testing.T) {
 const rate=48000.0
 const resistance=1000.0
 const inductance=0.1
 // Series RL branch with output across resistor: high frequencies attenuate.
 measure:=func(frequency float64)float64 {
  t.Helper()
  n,err:=NewRLNodalNetwork(rate,
   []ResistiveBranch{{A:"out",B:"",Ohms:resistance}},
   []InductiveBranch{{A:"in",B:"out",Henry:inductance}})
  if err!=nil {t.Fatal(err)}
  const samples=48000
  sum:=0.0
  count:=0
  for i:=0;i<samples;i++ {
   in:=math.Sin(2*math.Pi*frequency*float64(i)/rate)
   result,err:=n.Next(map[string]float64{"in":in,"":0})
   if err!=nil {t.Fatal(err)}
   if i>=samples/2 {sum+=result["out"]*result["out"];count++}
  }
  return math.Sqrt(2*sum/float64(count))
 }
 low:=measure(50)
 high:=measure(5000)
 if !(low>0.9 && high<0.5 && low>high) {
  t.Fatalf("synthetic RL response unexpected: 50Hz=%.4f, 5kHz=%.4f",low,high)
 }
}
