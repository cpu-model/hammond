package electrical

import "fmt"

// PedalCircuitEvidence describes the independent evidence needed before
// a pedal network can be labeled a B3 circuit reconstruction.
// Mere component presence in a drawing is not a conductive netlist.
type PedalCircuitEvidence struct {
 SchematicURL string
 BusbarToFilterEdges []TerminalEdge
 FilterToDrawbarEdges []TerminalEdge
 Reactor Parameter
}

// ValidateReconstruction rejects incomplete or untraceable circuit
// reconstructions. A populated parameter alone never establishes topology.
func (e PedalCircuitEvidence) ValidateReconstruction() error {
 if e.SchematicURL=="" { return fmt.Errorf("pedal schematic reference missing") }
 if len(e.BusbarToFilterEdges)==0 { return fmt.Errorf("pedal busbar-to-filter connections not transcribed") }
 if len(e.FilterToDrawbarEdges)==0 { return fmt.Errorf("pedal filter-to-drawbar connections not transcribed") }
 edges:=make([]TerminalEdge,0,len(e.BusbarToFilterEdges)+len(e.FilterToDrawbarEdges))
 edges=append(edges,e.BusbarToFilterEdges...)
 edges=append(edges,e.FilterToDrawbarEdges...)
 if err:=(TerminalGraph{Edges:edges}).Validate();err!=nil { return fmt.Errorf("pedal netlist: %w",err) }
 if err:=e.Reactor.Validate();err!=nil { return fmt.Errorf("pedal reactor: %w",err) }
 if e.Reactor.Unit!="H"||!positiveFinite(e.Reactor.Value) {
  return fmt.Errorf("pedal reactor requires positive inductance in H")
 }
 if e.Reactor.Provenance!=ProvenanceSchematic&&e.Reactor.Provenance!=ProvenanceMeasurement {
  return fmt.Errorf("pedal reactor must have schematic or measurement provenance")
 }
 return nil
}
