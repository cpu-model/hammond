package electrical

import (
 "errors"
 "math"
)

// ResistiveNode models a single linear, memoryless electrical node
// driven by voltage sources through series resistances and loaded to
// ground. It is a circuit-analysis primitive, NOT a calibrated B3 bus.
type ResistiveNode struct {
 conductance float64
 weightedVoltage float64
}

func (n *ResistiveNode) AddSource(voltage,resistanceOhm float64) error {
 if !finite(voltage)||!positiveFinite(resistanceOhm) {
  return errors.New("source voltage must be finite and resistance positive")
 }
 g:=1/resistanceOhm
 n.conductance+=g
 n.weightedVoltage+=voltage*g
 return nil
}

func (n *ResistiveNode) AddGroundLoad(resistanceOhm float64) error {
 if !positiveFinite(resistanceOhm) {
  return errors.New("load resistance must be positive and finite")
 }
 n.conductance+=1/resistanceOhm
 return nil
}

// Voltage returns the nodal voltage. A floating node is an error.
func (n ResistiveNode) Voltage() (float64,error) {
 if n.conductance==0 { return 0,errors.New("floating node") }
 v:=n.weightedVoltage/n.conductance
 if !finite(v) { return 0,errors.New("nonfinite nodal voltage") }
 return v,nil
}

func (n *ResistiveNode) Reset() {
 n.conductance=0
 n.weightedVoltage=0
}

func finite(v float64) bool { return !math.IsNaN(v)&&!math.IsInf(v,0) }
func positiveFinite(v float64) bool { return finite(v)&&v>0 }
