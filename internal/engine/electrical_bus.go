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
