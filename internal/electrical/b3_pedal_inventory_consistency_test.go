package electrical

import "testing"

// Every resistor/reactor identified in the B3 pedal schematic must occur
// exactly once in the combined component inventory. Printed numbers are
// not used as physical values.
func TestB3PedalSchematicInventoryMatchesAnalogComponents(t *testing.T) {
 markings:=B3PedalPrintedMarkings()
 topology:=B3PedalAnalogComponents()
 if err:=topology.Validate();err!=nil {t.Fatal(err)}
 byName:=map[string]PedalComponentKind{}
 for _,c:=range topology.Components {
  if _,ok:=byName[c.Designator];ok {t.Fatalf("duplicate analog component %s",c.Designator)}
  byName[c.Designator]=c.Kind
 }
 if len(markings)!=len(byName) {
  t.Fatalf("markings=%d components=%d",len(markings),len(byName))
 }
 for _,m:=range markings {
  kind,ok:=byName[m.Designator]
  if !ok {t.Fatalf("marked component %s absent from analog topology",m.Designator)}
  expected:=PedalResistor
  if m.Designator=="L20" {expected=PedalReactor}
  if kind!=expected {t.Fatalf("%s kind=%s want %s",m.Designator,kind,expected)}
 }
}
