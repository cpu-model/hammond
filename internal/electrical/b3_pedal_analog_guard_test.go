package electrical

import "testing"

// Factory B3 pedal components must not silently become an operational
// circuit while the filter terminals and measured component values
// remain unverified. This test protects that distinction.
func TestB3PedalAnalogRequiresExplicitValues(t *testing.T) {
 topology:=B3PedalAnalogComponents()
 if err:=topology.Validate();err!=nil {t.Fatal(err)}
 _,_,err:=CompilePedalRLNetwork(48000,topology,nil,"chassis-ground")
 if err==nil {t.Fatal("B3 analog components compiled without physical values")}
}

func TestB3PedalAnalogFilterRemainsIsolated(t *testing.T) {
 topology:=B3PedalAnalogComponents()
 for _,term:=range []string{"filter-R80-a","filter-R80-b","filter-L20-a","filter-L20-b"} {
  ok,err:=topology.SignalPathExists(term,"panel-common-out")
  if err!=nil {t.Fatal(err)}
  if ok {t.Fatalf("unverified connection from %s to resistor panel",term)}
 }
}
