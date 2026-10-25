package electrical

import "fmt"

// PedalComponentKind distinguishes electrical components from plain wires.
// A resistor or reactor must never be inserted into TerminalGraph as a
// conductive edge: doing so would short-circuit the component.
type PedalComponentKind string

const (
 PedalResistor PedalComponentKind = "resistor"
 PedalReactor PedalComponentKind = "reactor"
)

// PedalComponentPlacement records the two actual component terminals,
// not a conductive connection between them. Parameter values are optional
// until supported by the schematic or measurements.
type PedalComponentPlacement struct {
 Designator string
 Kind PedalComponentKind
 TerminalA string
 TerminalB string
 Source string
}

type PedalCircuitTopology struct {
 Wires []TerminalEdge
 Components []PedalComponentPlacement
}

func (t PedalCircuitTopology) Validate() error {
 if err:=(TerminalGraph{Edges:t.Wires}).Validate();err!=nil {return err}
 names:=map[string]bool{}
 for i,c:=range t.Components {
  if c.Designator==""||c.TerminalA==""||c.TerminalB==""||c.Source=="" {
   return fmt.Errorf("component %d: missing designator, terminals or source",i)
  }
  if c.TerminalA==c.TerminalB {return fmt.Errorf("%s: identical component terminals",c.Designator)}
  if c.Kind!=PedalResistor&&c.Kind!=PedalReactor {return fmt.Errorf("%s: unknown component kind %q",c.Designator,c.Kind)}
  if names[c.Designator] {return fmt.Errorf("%s: duplicate component",c.Designator)}
  names[c.Designator]=true
  shorted,err:=(TerminalGraph{Edges:t.Wires}).Connected(c.TerminalA,c.TerminalB)
  if err!=nil {return err}
  if shorted {return fmt.Errorf("%s: component terminals shorted by documented wires",c.Designator)}
 }
 return nil
}
