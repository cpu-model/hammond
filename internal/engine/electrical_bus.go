package engine

import (
 "fmt"

 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/generator"
 "github.com/cpu-model/hammond/internal/registration"
)

// ElectricalBusEngine is an isolated, uncalibrated experiment that
// drives nine independent resistive buses per manual from one shared
// mechanical tonewheel bank. No drawbar or transformer stage is implied.
type ElectricalBusEngine struct {
 generator *generator.MechanicalBank
 manuals [2]*registration.TimedManual
 buses [2]*electrical.NineBusNetwork
 frame [91]float64
 pedal *registration.PedalContactsState
 pedalMapping registration.PedalBusRouting
 pedalBuses *electrical.FourPedalBusNetwork
}

func NewElectricalBusEngine(sampleRate float64,upper,lower registration.Routing,upperTiming,lowerTiming registration.ContactTiming,upperBuses,lowerBuses [9]electrical.BusConfig) (*ElectricalBusEngine,error) {
 bank,err:=generator.NewMechanicalBank(sampleRate)
 if err!=nil { return nil,err }
 u,err:=registration.NewTimedManual(upper,upperTiming)
 if err!=nil { return nil,fmt.Errorf("upper manual: %w",err) }
 l,err:=registration.NewTimedManual(lower,lowerTiming)
 if err!=nil { return nil,fmt.Errorf("lower manual: %w",err) }
 ub,err:=electrical.NewNineBusNetwork(upperBuses)
 if err!=nil { return nil,fmt.Errorf("upper buses: %w",err) }
 lb,err:=electrical.NewNineBusNetwork(lowerBuses)
 if err!=nil { return nil,fmt.Errorf("lower buses: %w",err) }
 return &ElectricalBusEngine{generator:bank,manuals:[2]*registration.TimedManual{u,l},buses:[2]*electrical.NineBusNetwork{ub,lb}},nil
}

func (e *ElectricalBusEngine) Manual(index int) (*registration.TimedManual,error) {
 if index<0||index>=len(e.manuals) { return nil,fmt.Errorf("manual index %d outside 0..1",index) }
 return e.manuals[index],nil
}

// NextBuses advances the generator and contact clocks exactly once.
// Results are nine independent bus voltages for each manual. No legacy digital mix is computed.
func (e *ElectricalBusEngine) NextBuses() (upper,lower [9]float64) {
 e.generator.Next(e.frame[:])
 for i,m:=range e.manuals {
  m.AdvanceContacts()
  b:=e.buses[i]
  b.BeginSample()
  // Closure does not escape this call or allocate on the heap.
  m.ForEachActiveContact(func(bus,wheel int) {
   _=b.AddClosedContact(bus,e.frame[wheel-1])
  })
 }
 return e.buses[0].Voltages(),e.buses[1].Voltages()
}

// NextTransformerInputs advances the electrical bus model once and
// records each manual's selected transformer tap connections.
// It does not calculate transformer voltage transfer or loading.
func (e *ElectricalBusEngine) NextTransformerInputs(pair electrical.ManualTransformerPair) (electrical.TransformerInputs,electrical.TransformerInputs,error) {
 if err:=pair.Validate();err!=nil {
  return electrical.TransformerInputs{},electrical.TransformerInputs{},err
 }
 upper,lower:=e.NextBuses()
 return pair.Inputs(upper,lower)
}

// NextTransformerSources advances the manual electrical path once,
// then associates an externally supplied pedal source with T1.
// It does not generate pedal tones or solve the transformer circuit.
// Invalid tap/pedal configuration leaves engine sample state unchanged.
func (e *ElectricalBusEngine) NextTransformerSources(pair electrical.ManualTransformerPair,pedals electrical.PedalTransformerInputs) (electrical.TransformerInputs,electrical.T1InputBoundary,error) {
 if err:=pair.Validate();err!=nil {
  return electrical.TransformerInputs{},electrical.T1InputBoundary{},err
 }
 if err:=pedals.Validate();err!=nil {
  return electrical.TransformerInputs{},electrical.T1InputBoundary{},err
 }
 upper,lower:=e.NextBuses()
 t2,t1,err:=pair.Inputs(upper,lower)
 if err!=nil { return electrical.TransformerInputs{},electrical.T1InputBoundary{},err }
 boundary,err:=electrical.NewT1InputBoundary(t1,pedals)
 return t2,boundary,err
}

// RenderTransformerSources renders one electrical source record per
// pedal input sample. The entire input block is validated before any
// generator/contact state advances. This is not an audio mix.
func (e *ElectricalBusEngine) RenderTransformerSources(pair electrical.ManualTransformerPair,pedals []electrical.PedalTransformerInputs,upper []electrical.TransformerInputs,lower []electrical.T1InputBoundary) error {
 if len(upper)!=len(pedals)||len(lower)!=len(pedals) {
  return fmt.Errorf("transformer output lengths must match pedal input length")
 }
 if err:=pair.Validate();err!=nil { return err }
 for i,p:=range pedals {
  if err:=p.Validate();err!=nil { return fmt.Errorf("pedal sample %d: %w",i,err) }
 }
 for i,p:=range pedals {
  t2,t1,_:=e.NextTransformerSources(pair,p)
  upper[i],lower[i]=t2,t1
 }
 return nil
}

// NewElectricalBusEngineWithPedals adds an experimental four-bus pedal
// switch network to the same generator used by both manuals.
func NewElectricalBusEngineWithPedals(sampleRate float64,upper,lower registration.Routing,upperTiming,lowerTiming registration.ContactTiming,upperBuses,lowerBuses [9]electrical.BusConfig,pedalRouting registration.PedalRouting,pedalMapping registration.PedalBusRouting,pedalBuses [4]electrical.BusConfig) (*ElectricalBusEngine,error) {
 if err:=pedalMapping.Validate();err!=nil { return nil,err }
 p,err:=registration.NewPedalContactsState(pedalRouting)
 if err!=nil { return nil,err }
 pb,err:=electrical.NewFourPedalBusNetwork(pedalBuses)
 if err!=nil { return nil,err }
 e,err:=NewElectricalBusEngine(sampleRate,upper,lower,upperTiming,lowerTiming,upperBuses,lowerBuses)
 if err!=nil { return nil,err }
 e.pedal,e.pedalMapping,e.pedalBuses=p,pedalMapping,pb
 return e,nil
}

// Pedals returns the switch controller, or an error for manual-only engines.
func (e *ElectricalBusEngine) Pedals() (*registration.PedalContactsState,error) {
 if e.pedal==nil { return nil,fmt.Errorf("pedal network not configured") }
 return e.pedal,nil
}

// NextAllBuses advances the shared generator exactly once and returns
// independent manual and pedal voltages. No pedal reactor, drawbar
// or T1 coupling is simulated.
func (e *ElectricalBusEngine) NextAllBuses() (upper,lower [9]float64,pedals [4]float64,err error) {
 if e.pedal==nil { return upper,lower,pedals,fmt.Errorf("pedal network not configured") }
 upper,lower=e.NextBuses()
 e.pedalBuses.BeginSample()
 err:=e.pedal.ForEachClosedBusContact(e.pedalMapping,func(_,_,bus,wheel int) {
  // Routing and bus configuration were validated at construction.
  _=e.pedalBuses.AddClosedContact(bus,e.frame[wheel-1])
 })
 if err!=nil { return upper,lower,pedals,err }
 return upper,lower,e.pedalBuses.Voltages(),nil
}
