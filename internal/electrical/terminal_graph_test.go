package electrical

import "testing"

func TestTerminalGraphConductivePaths(t *testing.T) {
 g:=TerminalGraph{Edges:[]TerminalEdge{
  {From:"T1.A",To:"T1.B",Source:"fixture winding segment"},
  {From:"T1.B",To:"T1.C",Source:"fixture winding segment"},
  {From:"T2.A",To:"T2.B",Source:"fixture separate winding"},
 }}
 for _,tc:=range []struct{a,b string;want bool}{
  {"T1.A","T1.C",true},
  {"T1.C","T1.A",true},
  {"T1.A","T2.A",false},
  {"T1.A","unknown",false},
 } {
  got,err:=g.Connected(tc.a,tc.b)
  if err!=nil||got!=tc.want { t.Fatalf("%s to %s: %v, %v",tc.a,tc.b,got,err) }
 }
}

func TestTerminalGraphRequiresEvidenceAndNoDuplicateEdges(t *testing.T) {
 cases:=[]TerminalGraph{
  {Edges:[]TerminalEdge{{From:"A",To:"B"}}},
  {Edges:[]TerminalEdge{{From:"A",To:"A",Source:"fixture"}}},
  {Edges:[]TerminalEdge{{From:"A",To:"B",Source:"fixture"},{From:"B",To:"A",Source:"fixture"}}},
  {Edges:[]TerminalEdge{{From:"",To:"B",Source:"fixture"}}},
 }
 for i,g:=range cases {
  if err:=g.Validate();err==nil { t.Fatalf("case %d accepted",i) }
 }
}
