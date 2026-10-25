package engine

import (
 "fmt"

 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

// NewElectricalBusEngineWithB3PedalWiring connects the documented B3
// harmonic frequency table to the shared generator and four pedal buses.
// Physical contact order and busbar assignments must be supplied from
// independently inspected wiring evidence; no factory order is assumed.
func NewElectricalBusEngineWithB3PedalWiring(
 sampleRate float64,
 upper, lower registration.Routing,
 upperTiming, lowerTiming registration.ContactTiming,
 upperBuses, lowerBuses [9]electrical.BusConfig,
 wiring registration.PedalWiringEvidence,
 pedalBuses [4]electrical.BusConfig,
) (*ElectricalBusEngine, error) {
 routing, busbars, err := registration.CompileB3PedalWiring(wiring)
 if err != nil { return nil, fmt.Errorf("B3 pedal wiring: %w",err) }
 return NewElectricalBusEngineWithPedals(
  sampleRate,upper,lower,upperTiming,lowerTiming,
  upperBuses,lowerBuses,routing,busbars,pedalBuses,
 )
}
