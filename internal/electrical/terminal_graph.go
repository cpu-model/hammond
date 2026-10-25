package electrical

import "fmt"

// TerminalEdge represents a documented electrical connection between
// two named terminals. It does not imply a resistance or transformer
// coupling coefficient.
type TerminalEdge struct {
 From string
 To string
 Source string
}

// TerminalGraph holds only directly documented conductive links.
// Magnetic coupling and winding impedance are intentionally absent.
type TerminalGraph struct {
 Edges []TerminalEdge
}

func (g TerminalGraph) Validate() error {
 seen:=make(map[string]bool,len(g.Edges))
 for i,e:=range g.Edges {
  if e.From==""||e.To=="" { return fmt.Errorf("edge %d: terminal name missing",i) }
  if e.From==e.To { return fmt.Errorf("edge %d: self connection",i) }
  if e.Source=="" { return fmt.Errorf("edge %d: source reference missing",i) }
  a,b:=e.From,e.To
  if a>b { a,b=b,a }
  key:=fmt.Sprintf("%d:%s|%d:%s",len(a),a,len(b),b)
  if seen[key] { return fmt.Errorf("edge %d: duplicate connection",i) }
  seen[key]=true
 }
 return nil
}

// Connected reports whether two terminals share a path of recorded
// conductive edges. It does not infer a link through magnetic coupling.
func (g TerminalGraph) Connected(from,to string) (bool,error) {
 if err:=g.Validate();err!=nil { return false,err }
 if from==""||to=="" { return false,fmt.Errorf("terminal name missing") }
 if from==to { return true,nil }
 visited:=map[string]bool{from:true}
 queue:=[]string{from}
 for len(queue)>0 {
  n:=queue[0]
  queue=queue[1:]
  for _,e:=range g.Edges {
   next:=""
   if e.From==n { next=e.To } else if e.To==n { next=e.From }
   if next==""||visited[next] { continue }
   if next==to { return true,nil }
   visited[next]=true
   queue=append(queue,next)
  }
 }
 return false,nil
}
