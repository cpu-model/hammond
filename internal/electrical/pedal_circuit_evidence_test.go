package electrical

import "testing"

func TestPedalCircuitEvidenceRequiresActualNetlist(t *testing.T) {
 e:=PedalCircuitEvidence{
  SchematicURL:"https://synthmania.com/wp-content/uploads/2020/05/Hammond-Organ-B3-C3-Service-Manual.pdf",
  Reactor:Parameter{Value:1,Unit:"H",Provenance:ProvenanceExperimental},
 }
 if err:=e.ValidateReconstruction();err==nil { t.Fatal("accepted schematic URL without wiring") }
 e.BusbarToFilterEdges=[]TerminalEdge{{From:"bus-1",To:"L20-a",Source:"schematic trace A"}}
 e.FilterToDrawbarEdges=[]TerminalEdge{{From:"L20-a",To:"drawbar-1",Source:"synthetic trace B"}}
 if err:=e.ValidateReconstruction();err==nil { t.Fatal("accepted experimental reactor value") }
 e.Reactor.Provenance=ProvenanceSchematic
 e.Reactor.Source="schematic measurement example"
 if err:=e.ValidateReconstruction();err!=nil { t.Fatalf("rejected structurally complete synthetic evidence: %v",err) }
 e.FilterToDrawbarEdges[0].Source=""
 if err:=e.ValidateReconstruction();err==nil { t.Fatal("accepted unsourced netlist edge") }
}

func TestPedalCircuitEvidenceRejectsDuplicateEdges(t *testing.T) {
 e:=PedalCircuitEvidence{
  SchematicURL:"schematic",
  BusbarToFilterEdges:[]TerminalEdge{{From:"a",To:"b",Source:"trace"}},
  FilterToDrawbarEdges:[]TerminalEdge{{From:"b",To:"a",Source:"trace"}},
  Reactor:Parameter{Value:1,Unit:"H",Provenance:ProvenanceMeasurement,Source:"measurement"},
 }
 if err:=e.ValidateReconstruction();err==nil { t.Fatal("accepted duplicate conductive edges") }
}

func TestPedalCircuitEvidenceRejectsDisconnectedSubgraphs(t *testing.T) {
 e:=PedalCircuitEvidence{
  SchematicURL:"synthetic schematic",
  BusbarToFilterEdges:[]TerminalEdge{{From:"bus-1",To:"reactor-a",Source:"synthetic edge 1"}},
  FilterToDrawbarEdges:[]TerminalEdge{{From:"reactor-b",To:"drawbar",Source:"synthetic edge 2"}},
  Reactor:Parameter{Value:1,Unit:"H",Provenance:ProvenanceMeasurement,Source:"synthetic measurement"},
 }
 if err:=e.ValidateReconstruction();err==nil { t.Fatal("accepted disconnected reactor circuit") }
 e.FilterToDrawbarEdges[0].From="reactor-a"
 if err:=e.ValidateReconstruction();err!=nil { t.Fatalf("rejected connected synthetic graph: %v",err) }
}

func TestPedalCircuitEvidenceRejectsPartiallyFloatingOutputs(t *testing.T) {
 e:=PedalCircuitEvidence{
  SchematicURL:"synthetic schematic",
  BusbarToFilterEdges:[]TerminalEdge{{From:"bus-1",To:"filter-a",Source:"fixture"}},
  FilterToDrawbarEdges:[]TerminalEdge{
   {From:"filter-a",To:"drawbar-1",Source:"fixture"},
   {From:"floating",To:"drawbar-2",Source:"fixture"},
  },
  Reactor:Parameter{Value:1,Unit:"H",Provenance:ProvenanceMeasurement,Source:"synthetic measurement"},
 }
 if err:=e.ValidateReconstruction();err==nil { t.Fatal("accepted one disconnected output among connected outputs") }
 e.FilterToDrawbarEdges[1].From="filter-a"
 if err:=e.ValidateReconstruction();err!=nil { t.Fatalf("rejected fully connected fixture: %v",err) }
}
