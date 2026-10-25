package engine

import (
 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

// NewElectricalBusEngineWithB3FactoryPedals installs the sourced B3
// pedal contact permutation and four colored busbar groups. The frequency
// matrix comes from B3PedalHarmonics. Electrical bus impedances remain
// explicit experimental parameters; this does not model the original
// resistor panel, pedal filter or transformer.
func NewElectricalBusEngineWithB3FactoryPedals(
 sampleRate float64,
 upper, lower registration.Routing,
 upperTiming, lowerTiming registration.ContactTiming,
 upperBuses, lowerBuses [9]electrical.BusConfig,
 pedalBuses [4]electrical.BusConfig,
) (*ElectricalBusEngine,error) {
 return NewElectricalBusEngineWithB3PedalWiring(
  sampleRate,upper,lower,upperTiming,
  lowerTiming,upperBuses,lowerBuses,
  registration.B3PedalWiringHandyGuide(),pedalBuses,
 )
}
