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
 // Every downstream edge must belong to a component connected to
 // at least one upstream bus node. A single connected edge is not enough.
 graph:=TerminalGraph{Edges:edges}
 upstream:=map[string]bool{}
 for _,edge:=range e.BusbarToFilterEdges {
  upstream[edge.From]=true
  upstream[edge.To]=true
 }
 for i,edge:=range e.FilterToDrawbarEdges {
  connected:=false
  for node:=range upstream {
   linked,err:=graph.Connected(node,edge.From)
   if err!=nil { return fmt.Errorf("pedal netlist: %w",err) }
   if linked { connected=true;break }
  }
  if !connected { return fmt.Errorf("pedal filter-to-drawbar edge %d disconnected from busbars",i) }
 }
 if err:=e.Reactor.Validate();err!=nil { return fmt.Errorf("pedal reactor: %w",err) }
 if e.Reactor.Unit!="H"||!positiveFinite(e.Reactor.Value) {
  return fmt.Errorf("pedal reactor requires positive inductance in H")
 }
 if e.Reactor.Provenance!=ProvenanceSchematic&&e.Reactor.Provenance!=ProvenanceMeasurement {
  return fmt.Errorf("pedal reactor must have schematic or measurement provenance")
 }
 return nil
}
