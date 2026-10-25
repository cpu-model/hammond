package engine

import (
 "testing"

 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

func TestB3PedalWiringEngineRejectsUnknownWiring(t *testing.T) {
 _,err:=NewElectricalBusEngineWithB3PedalWiring(
  44100,registration.Routing{},registration.Routing{},
  registration.ContactTiming{},registration.ContactTiming{},
  [9]electrical.BusConfig{},[9]electrical.BusConfig{},
  registration.PedalWiringEvidence{},
  [4]electrical.BusConfig{},
 )
 if err==nil { t.Fatal("accepted missing B3 physical wiring evidence") }
}
