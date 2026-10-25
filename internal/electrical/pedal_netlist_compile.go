package electrical

import "fmt"

// PedalComponentValue binds a topology designator to a sourced physical
// value. Resistors use ohm; reactors use H. No schematic defaults exist.
type PedalComponentValue struct {
 Designator string
 Parameter Parameter
}

// CompilePedalRLNetwork converts an evidence-bearing terminal topology into
// a solvable network. Wires collapse into common nodes; components retain
// their impedance. GroundTerminal must be a named schematic terminal.
func CompilePedalRLNetwork(sampleRate float64,topology PedalCircuitTopology,values []PedalComponentValue,groundTerminal string) (*RLNodalNetwork,map[string]string,error) {
 if err:=topology.Validate();err!=nil {return nil,nil,err}
 if groundTerminal=="" {return nil,nil,fmt.Errorf("ground terminal required")}
 parent:=map[string]string{}
 var root func(string)string
 root=func(x string)string {
  p,ok:=parent[x];if !ok {parent[x]=x;return x}
  if p!=x {parent[x]=root(p)}
  return parent[x]
 }
 for _,w:=range topology.Wires {
  a,b:=root(w.From),root(w.To)
  if a!=b {parent[b]=a}
 }
 for _,c:=range topology.Components {root(c.TerminalA);root(c.TerminalB)}
 root(groundTerminal)
 parameters:=map[string]Parameter{}
 for _,v:=range values {
  if v.Designator=="" {return nil,nil,fmt.Errorf("empty value designator")}
  if _,ok:=parameters[v.Designator];ok {return nil,nil,fmt.Errorf("duplicate value for %s",v.Designator)}
  if err:=v.Parameter.Validate();err!=nil {return nil,nil,fmt.Errorf("%s: %w",v.Designator,err)}
  parameters[v.Designator]=v.Parameter
 }
 resistors:=[]ResistiveBranch{}
 inductors:=[]InductiveBranch{}
 used:=map[string]bool{}
 ground:=root(groundTerminal)
 nodes:=map[string]string{}
 node:=func(terminal string)string {
  r:=root(terminal)
  if r==ground {return ""}
  return r
 }
 for terminal:=range parent {nodes[terminal]=node(terminal)}
 for _,c:=range topology.Components {
  p,ok:=parameters[c.Designator]
  if !ok {return nil,nil,fmt.Errorf("missing physical value for %s",c.Designator)}
  used[c.Designator]=true
  a,b:=node(c.TerminalA),node(c.TerminalB)
  if a==b {return nil,nil,fmt.Errorf("%s: electrically shorted",c.Designator)}
  switch c.Kind {
  case PedalResistor:
   if p.Unit!="ohm" {return nil,nil,fmt.Errorf("%s: expected ohm, got %q",c.Designator,p.Unit)}
   resistors=append(resistors,ResistiveBranch{A:a,B:b,Ohms:p.Value})
  case PedalReactor:
   if p.Unit!="H" {return nil,nil,fmt.Errorf("%s: expected H, got %q",c.Designator,p.Unit)}
   inductors=append(inductors,InductiveBranch{A:a,B:b,Henry:p.Value})
  }
 }
 for name:=range parameters {if !used[name] {return nil,nil,fmt.Errorf("value for unknown component %s",name)}}
 network,err:=NewRLNodalNetwork(sampleRate,resistors,inductors)
 if err!=nil {return nil,nil,err}
 return network,nodes,nil
}
