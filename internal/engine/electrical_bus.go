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
 sampleRate float64
 manuals [2]*registration.TimedManual
 buses [2]*electrical.NineBusNetwork
 frame [91]float64
 pedal *registration.PedalContactsState
 pedalMapping registration.PedalBusRouting
 pedalBuses *electrical.FourPedalBusNetwork
 experimentalPedalFilters *electrical.ExperimentalPedalBusFilters
 experimentalPedalCircuit *electrical.RLNodalNetwork
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
 return &ElectricalBusEngine{generator:bank,sampleRate:sampleRate,manuals:[2]*registration.TimedManual{u,l},buses:[2]*electrical.NineBusNetwork{ub,lb}},nil
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
 err=e.pedal.ForEachClosedBusContact(e.pedalMapping,func(_,_,bus,wheel int) {
  // Routing and bus configuration were validated at construction.
  _=e.pedalBuses.AddClosedContact(bus,e.frame[wheel-1])
 })
 if err!=nil { return upper,lower,pedals,err }
 return upper,lower,e.pedalBuses.Voltages(),nil
}

// EnableExperimentalPedalFilters installs four independent RL stages.
// This is an explicitly synthetic experiment, not a B3 circuit model.
// Validation completes before changing the engine's filter configuration.
func (e *ElectricalBusEngine) EnableExperimentalPedalFilters(sampleRate float64,configs [4]electrical.PedalRLConfig) error {
 if e.pedal==nil { return fmt.Errorf("pedal network not configured") }
 if sampleRate!=e.sampleRate { return fmt.Errorf("pedal filter sample rate %g does not match generator rate %g",sampleRate,e.sampleRate) }
 filters,err:=electrical.NewExperimentalPedalBusFilters(sampleRate,configs)
 if err!=nil { return err }
 e.experimentalPedalFilters=filters
 return nil
}

// DisableExperimentalPedalFilters restores raw pedal bus output.
func (e *ElectricalBusEngine) DisableExperimentalPedalFilters() {
 e.experimentalPedalFilters=nil
}

// NextExperimentalPedalBuses advances the shared tonewheel generator
// exactly once and optionally processes the four pedal bus voltages.
// Manual bus voltages remain unfiltered.
func (e *ElectricalBusEngine) NextExperimentalPedalBuses() (upper,lower [9]float64,pedals [4]float64,err error) {
 upper,lower,pedals,err=e.NextAllBuses()
 if err!=nil { return upper,lower,pedals,err }
 if e.experimentalPedalFilters!=nil {
  pedals,err=e.experimentalPedalFilters.Next(pedals)
 }
 return upper,lower,pedals,err
}

// EnableExperimentalPedalCircuit installs a caller-defined, shared RL
// network. The four ideal source nodes are pedal-bus-0 through pedal-bus-3;
// outputNode names the observed node. This is not a verified B3 netlist.
func (e *ElectricalBusEngine) EnableExperimentalPedalCircuit(resistors []electrical.ResistiveBranch,inductors []electrical.InductiveBranch) error {
 if e.pedal==nil {return fmt.Errorf("pedal network not configured")}
 network,err:=electrical.NewRLNodalNetwork(e.sampleRate,resistors,inductors)
 if err!=nil {return err}
 e.experimentalPedalCircuit=network
 return nil
}

// DisableExperimentalPedalCircuit removes the optional circuit and its
// inductor history. The original pedal bus voltages remain available.
func (e *ElectricalBusEngine) DisableExperimentalPedalCircuit() {
 e.experimentalPedalCircuit=nil
}

// NextExperimentalPedalCircuit advances the mechanical bank once and
// drives the optional RL network with four independent pedal bus voltages.
// It does not infer drawbar connections or transformer coupling.
func (e *ElectricalBusEngine) NextExperimentalPedalCircuit(outputNode string) (upper,lower [9]float64,output float64,err error) {
 if e.experimentalPedalCircuit==nil {return upper,lower,0,fmt.Errorf("experimental pedal circuit not configured")}
 if outputNode=="" {return upper,lower,0,fmt.Errorf("output node required")}
 found:=outputNode=="pedal-bus-0"||outputNode=="pedal-bus-1"||outputNode=="pedal-bus-2"||outputNode=="pedal-bus-3"
 for _,b:=range e.experimentalPedalCircuit.Resistors {
  if b.A==outputNode||b.B==outputNode {found=true}
 }
 for _,b:=range e.experimentalPedalCircuit.Inductors {
  if b.A==outputNode||b.B==outputNode {found=true}
 }
 if !found {return upper,lower,0,fmt.Errorf("output node %q absent from circuit",outputNode)}
 upper,lower,pedals,err:=e.NextAllBuses()
 if err!=nil {return upper,lower,0,err}
 fixed:=map[string]float64{
  "pedal-bus-0":pedals[0],
  "pedal-bus-1":pedals[1],
  "pedal-bus-2":pedals[2],
  "pedal-bus-3":pedals[3],
 }
 nodes,err:=e.experimentalPedalCircuit.Next(fixed)
 if err!=nil {return upper,lower,0,err}
 v,ok:=nodes[outputNode]
 if !ok {return upper,lower,0,fmt.Errorf("output node %q absent from circuit",outputNode)}
 return upper,lower,v,nil
}
