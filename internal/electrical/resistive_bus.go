package electrical

import "errors"

// BusConfig defines one experimental resistive bus.
// All electrical values are explicit and provenance-tagged.
type BusConfig struct {
 SourceResistanceOhm Parameter
 LoadResistanceOhm Parameter
}

func (c BusConfig) Validate() error {
 for _,p:=range []Parameter{c.SourceResistanceOhm,c.LoadResistanceOhm} {
  if err:=p.Validate();err!=nil { return err }
  if p.Unit!="ohm"||!positiveFinite(p.Value) {
   return errors.New("bus resistance must be positive and in ohm")
  }
 }
 return nil
}

// ResistiveBus is an uncalibrated one-node electrical experiment.
// Its source impedance and ground load are not B3 specifications.
type ResistiveBus struct {
 sourceResistance float64
 loadResistance float64
 node ResistiveNode
}

func NewResistiveBus(c BusConfig) (*ResistiveBus,error) {
 if err:=c.Validate();err!=nil { return nil,err }
 return &ResistiveBus{
  sourceResistance:c.SourceResistanceOhm.Value,
  loadResistance:c.LoadResistanceOhm.Value,
 },nil
}

// BeginSample clears prior contributions and adds the ground load.
func (b *ResistiveBus) BeginSample() {
 b.node.Reset()
 // Validated at construction; cannot fail.
 _=b.node.AddGroundLoad(b.loadResistance)
}

// AddSource represents one closed contact driven by an ideal source.
// A caller must only add sources whose contacts are actually closed.
func (b *ResistiveBus) AddSource(voltage float64) error {
 return b.node.AddSource(voltage,b.sourceResistance)
}

func (b *ResistiveBus) Voltage() float64 {
 v,_:=b.node.Voltage() // ground load guarantees a nonfloating node
 return v
}
