package electrical

import "fmt"

// FourPedalBusNetwork is an experimental, independently loaded
// four-node pedal switch bus model. It does not include the
// B3 resistor panel, filter reactor, drawbars or T1 loading.
type FourPedalBusNetwork struct {
 buses [4]*ResistiveBus
}

func NewFourPedalBusNetwork(config [4]BusConfig) (*FourPedalBusNetwork,error) {
 n:=&FourPedalBusNetwork{}
 for i,c:=range config {
  b,err:=NewResistiveBus(c)
  if err!=nil { return nil,fmt.Errorf("pedal bus %d: %w",i,err) }
  n.buses[i]=b
 }
 return n,nil
}

func (n *FourPedalBusNetwork) BeginSample() {
 for _,b:=range n.buses { b.BeginSample() }
}

func (n *FourPedalBusNetwork) AddClosedContact(bus int,voltage float64) error {
 if bus<0||bus>=len(n.buses) { return fmt.Errorf("pedal bus %d outside 0..3",bus) }
 return n.buses[bus].AddSource(voltage)
}

func (n *FourPedalBusNetwork) Voltages() [4]float64 {
 var v [4]float64
 for i,b:=range n.buses { v[i]=b.Voltage() }
 return v
}
