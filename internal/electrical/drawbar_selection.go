package electrical

import "fmt"

// DrawbarTapSelection describes nine independent user selections.
// Positions are 0 (disconnected) through 8. No electrical gain,
// transformer turns ratio or loading is implied by these positions.
type DrawbarTapSelection [9]uint8

func (s DrawbarTapSelection) Validate() error {
 for i,p:=range s {
  if p>8 { return fmt.Errorf("drawbar %d: position %d outside 0..8",i+1,p) }
 }
 return nil
}

// SelectedBusVoltages masks disconnected drawbar conductors. Connected
// buses retain their raw voltage; this is a routing diagnostic only,
// NOT an audio mix or B3 drawbar attenuation model.
func (s DrawbarTapSelection) SelectedBusVoltages(bus [9]float64) ([9]float64,error) {
 if err:=s.Validate();err!=nil { return [9]float64{},err }
 for i,p:=range s {
  if p==0 { bus[i]=0 }
 }
 return bus,nil
}
