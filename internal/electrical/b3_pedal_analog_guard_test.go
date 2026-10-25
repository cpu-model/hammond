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

// The ladder is internal to the panel; factory pedal-bus assignments
// are not implied by these component connections.
func TestB3PedalAnalogNoInventedBusbarConnections(t *testing.T) {
 topology:=B3PedalAnalogComponents()
 for _,bus:=range []string{"pedal-bus-0","pedal-bus-1","pedal-bus-2","pedal-bus-3"} {
  ok,err:=topology.SignalPathExists(bus,"panel-common-out")
  if err!=nil {t.Fatal(err)}
  if ok {t.Fatalf("unverified busbar %s connected to resistor panel",bus)}
 }
}
