package electrical

import "fmt"

// NineBusNetwork is a nine-conductor resistive experiment.
// It is NOT a Hammond transformer, preset or drawbar circuit.
type NineBusNetwork struct {
 buses [9]*ResistiveBus
}

// NewNineBusNetwork requires an explicit configuration for every bus.
// No default B3 resistance or load is assumed.
func NewNineBusNetwork(config [9]BusConfig) (*NineBusNetwork,error) {
 n:=&NineBusNetwork{}
 for i:=range n.buses {
  bus,err:=NewResistiveBus(config[i])
  if err!=nil { return nil,fmt.Errorf("bus %d: %w",i+1,err) }
  n.buses[i]=bus
 }
 return n,nil
}

// BeginSample clears all nine buses and installs their ground loads.
func (n *NineBusNetwork) BeginSample() {
 for _,b:=range n.buses { b.BeginSample() }
}

// AddClosedContact connects one ideal source through its configured
// source resistance to the selected bus.
func (n *NineBusNetwork) AddClosedContact(bus int,voltage float64) error {
 if bus<0||bus>=len(n.buses) { return fmt.Errorf("bus index %d outside 0..8",bus) }
 return n.buses[bus].AddSource(voltage)
}

// Voltages returns independent nodal voltages without mixing buses.
func (n *NineBusNetwork) Voltages() [9]float64 {
 var out [9]float64
 for i,b:=range n.buses { out[i]=b.Voltage() }
 return out
}
