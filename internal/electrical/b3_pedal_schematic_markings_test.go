package electrical

import (
 "fmt"
 "testing"
)

// Printed schematic numbers are documentary evidence, not physical
// parameters. This test guards against accidentally treating them as
// solver-ready ohms or henries.
func TestB3PedalPrintedMarkingsAreNotPhysicalValues(t *testing.T) {
 markings:=B3PedalPrintedMarkings()
 if len(markings)!=8 {t.Fatalf("got %d markings, want 8",len(markings))}
 expected:=map[string]string{
  "R80":"10","R81":"10","R82":"47","R83":"470",
  "R84":"5","R85":"5","R86":"20","L20":"(1000)",
 }
 seen:=map[string]bool{}
 for _,m:=range markings {
  want,ok:=expected[m.Designator]
  if !ok {t.Fatalf("unexpected designator %s",m.Designator)}
  if seen[m.Designator] {t.Fatalf("duplicate designator %s",m.Designator)}
  seen[m.Designator]=true
  if m.PrintedNumber!=want {t.Fatalf("%s marking %q, want %q",m.Designator,m.PrintedNumber,want)}
  if m.Source=="" {t.Fatalf("%s missing source",m.Designator)}
 }
 for name:=range expected {
  if !seen[name] {t.Fatal(fmt.Sprintf("missing marking %s",name))}
 }
 // A documented marking alone is insufficient to compile the panel.
 // Every actual physical component value must be independently supplied.
 if _,_,err:=CompilePedalRLNetwork(48000,B3PedalResistorPanelTopology(),nil,"chassis-ground");err==nil {
  t.Fatal("compiled schematic markings without physical parameters")
 }
}
