package electrical

import "fmt"

// ComponentOnEverySignalPath reports whether removing a component breaks
// all paths between two terminals. This is a topological cut test, not an
// electrical gain or impedance calculation.
func (t PedalCircuitTopology) ComponentOnEverySignalPath(from,to,designator string) (bool,error) {
 if err:=t.Validate();err!=nil {return false,err}
 if from==""||to==""||designator=="" {return false,fmt.Errorf("signal path endpoints and component designator required")}
 found:=false
 for _,c:=range t.Components {
  if c.Designator==designator {found=true;break}
 }
 if !found {return false,fmt.Errorf("unknown component %q",designator)}
 original,err:=t.SignalPathExists(from,to)
 if err!=nil {return false,err}
 if !original {return false,fmt.Errorf("no original signal path from %q to %q",from,to)}
 filtered:=PedalCircuitTopology{Wires:t.Wires}
 for _,c:=range t.Components {
  if c.Designator!=designator {filtered.Components=append(filtered.Components,c)}
 }
 bypass,err:=filtered.SignalPathExists(from,to)
 if err!=nil {return false,err}
 return !bypass,nil
}
