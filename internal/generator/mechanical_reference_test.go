package generator

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/instrument"
)

// Compare every mechanical output with the independently calculated
// nominal frequency model. This catches wrong speed-group mapping,
// particularly the special 192-tooth wheels 85..91.
func TestMechanicalBankAll91Frequencies(t *testing.T) {
 const sampleRate=48000.0
 bank,err:=NewMechanicalBank(sampleRate)
 if err!=nil { t.Fatal(err) }
 frequencies:=instrument.MechanicalFrequencies()
 frame:=make([]float64,instrument.TonewheelCount)
 for sample:=0;sample<2400;sample++ {
  bank.Next(frame)
  for wheel,f:=range frequencies {
   want:=math.Sin(2*math.Pi*f*float64(sample)/sampleRate)
   if math.Abs(frame[wheel]-want)>1e-8 {
    t.Fatalf("sample %d wheel %d: got %.12f want %.12f (%.6f Hz)",sample,wheel+1,frame[wheel],want,f)
   }
  }
 }
}
