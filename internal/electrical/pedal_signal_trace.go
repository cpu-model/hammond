package electrical

import "fmt"

// PedalSignalPath is an ordered topological trace. Components are traversed
// as components, never reclassified as conductive wires.
type PedalSignalPath struct {
 Terminals []string
 Components []string
}

// TraceSignalPath returns one shortest-hop path through wires and
// components. It establishes connectivity only, not gain or impedance.
func (t PedalCircuitTopology) TraceSignalPath(from,to string) (PedalSignalPath,bool,error) {
 if err:=t.Validate();err!=nil {return PedalSignalPath{},false,err}
 if from==""||to=="" {return PedalSignalPath{},false,fmt.Errorf("signal path terminal missing")}
 type link struct {to,component string}
 adj:=map[string][]link{}
 add:=func(a,b,component string) {
  adj[a]=append(adj[a],link{b,component})
  adj[b]=append(adj[b],link{a,component})
 }
 for _,w:=range t.Wires {add(w.From,w.To,"")}
 for _,c:=range t.Components {add(c.TerminalA,c.TerminalB,c.Designator)}
 type step struct {prev,component string}
 visited:=map[string]step{from:{}}
 queue:=[]string{from}
 found:=false
 for len(queue)>0 {
  node:=queue[0];queue=queue[1:]
  if node==to {found=true;break}
  for _,next:=range adj[node] {
   if _,ok:=visited[next.to];ok {continue}
   visited[next.to]=step{node,next.component}
   queue=append(queue,next.to)
  }
 }
 if !found {return PedalSignalPath{},false,nil}
 path:=PedalSignalPath{}
 for n:=to;; {
  path.Terminals=append(path.Terminals,n)
  if n==from {break}
  s:=visited[n]
  if s.component!="" {path.Components=append(path.Components,s.component)}
  n=s.prev
 }
 for i,j:=0,len(path.Terminals)-1;i<j;i,j=i+1,j-1 {
  path.Terminals[i],path.Terminals[j]=path.Terminals[j],path.Terminals[i]
 }
 for i,j:=0,len(path.Components)-1;i<j;i,j=i+1,j-1 {
  path.Components[i],path.Components[j]=path.Components[j],path.Components[i]
 }
 return path,true,nil
}
