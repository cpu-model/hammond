package engine

import (
 "fmt"
 "math"
)

// ManualMix is an audible, deliberately uncalibrated output stage.
// Each gain weights one of the nine independently simulated manual buses.
// This is not a model of the B3 matching transformer or preamplifier.
type ManualMix struct {
 Upper [9]float64
 Lower [9]float64
 OutputGain float64
}

func (m ManualMix) Validate() error {
 if math.IsNaN(m.OutputGain)||math.IsInf(m.OutputGain,0)||m.OutputGain<0 {
  return fmt.Errorf("invalid output gain")
 }
 for _,bank:=range [][9]float64{m.Upper,m.Lower} {
  for _,gain:=range bank {
   if math.IsNaN(gain)||math.IsInf(gain,0)||gain<0 {
    return fmt.Errorf("invalid manual bus gain")
   }
  }
 }
 return nil
}

// NextAudio advances the shared generator exactly once and sums the
// manual bus voltages. The caller must supply explicit, validated gains.
func (e *ElectricalBusEngine) NextAudio(m ManualMix) (float64,error) {
 if err:=m.Validate();err!=nil {return 0,err}
 upper,lower:=e.NextBuses()
 var sum float64
 for i:=0;i<9;i++ {
  sum+=upper[i]*m.Upper[i]+lower[i]*m.Lower[i]
 }
 return sum*m.OutputGain,nil
}
