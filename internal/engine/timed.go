package engine

import (
 "fmt"
 "math"

 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/generator"
 "github.com/cpu-model/hammond/internal/registration"
)

// TimedEngine renders two independently switched manuals against the same
// continuously rotating 91-output mechanical generator.
type TimedEngine struct {
 generator *generator.MechanicalBank
 manuals [2]*registration.TimedManual
 frame [91]float64
 percussion [2]Percussion
 sampleRate float64
}

// NewTimed accepts explicit routing and contact timings for both manuals.
// No physical contact delays are assumed.
func NewTimed(sampleRate float64, upper, lower registration.Routing, upperTiming, lowerTiming registration.ContactTiming) (*TimedEngine,error) {
 bank,err:=generator.NewMechanicalBank(sampleRate)
 if err!=nil { return nil,err }
 u,err:=registration.NewTimedManual(upper,upperTiming)
 if err!=nil { return nil,fmt.Errorf("upper manual: %w",err) }
 l,err:=registration.NewTimedManual(lower,lowerTiming)
 if err!=nil { return nil,fmt.Errorf("lower manual: %w",err) }
 u.SetNinthDrawbarTransition(int(math.Round(sampleRate*0.003)))
 return &TimedEngine{generator:bank,manuals:[2]*registration.TimedManual{u,l},sampleRate:sampleRate},nil
}

func (e *TimedEngine) Manual(index int) (*registration.TimedManual,error) {
 if index<0 || index>1 { return nil,fmt.Errorf("manual index %d outside 0..1",index) }
 return e.manuals[index],nil
}

// SetWheelSecondHarmonic configures the experimental pickup waveform
// of a specific shared tonewheel, numbered 1..91.
func (e *TimedEngine) SetWheelSecondHarmonic(wheel int,level float64) error {
 return e.generator.SetWheelSecondHarmonic(wheel,level)
}

// SetWheelPhaseOffset configures the experimental phase of one shared wheel.
func (e *TimedEngine) SetWheelPhaseOffset(wheel int,cycles float64) error {
 return e.generator.SetWheelPhaseOffset(wheel,cycles)
}

// SetPhaseSeed assigns deterministic, distinct phase offsets to all 91 wheels.
// The seed is an experimental comparison parameter, not a B3 calibration.
func (e *TimedEngine) SetPhaseSeed(seed uint64) error {
 for wheel:=1;wheel<=91;wheel++ {
  x:=seed+uint64(wheel)*0x9e3779b97f4a7c15
  x=(x^(x>>30))*0xbf58476d1ce4e5b9
  x=(x^(x>>27))*0x94d049bb133111eb
  x^=x>>31
  cycles:=float64(x>>11)/9007199254740992.0
  if err:=e.generator.SetWheelPhaseOffset(wheel,cycles);err!=nil{return err}
 }
 return nil
}

// SetSecondHarmonic applies the same experimental shape to all wheels.
func (e *TimedEngine) SetSecondHarmonic(level float64) error {
 return e.generator.SetSecondHarmonic(level)
}

// NextPaths advances the shared generator once and preserves manual outputs.
// Upper and lower are not electrically combined here.
func (e *TimedEngine) NextPaths() electrical.ManualPaths {
 e.generator.Next(e.frame[:])
 upper:=e.manuals[0].Mix(e.frame[:])
 lower:=e.manuals[1].Mix(e.frame[:])
 if e.manuals[0].TakePercussionTrigger() {e.percussion[0].Trigger()}
 if e.manuals[1].TakePercussionTrigger() {e.percussion[1].Trigger()}
 if e.percussion[0].harmonic==0 {
  upper+=e.percussion[0].Next(e.frame[:])
 } else {
  signal:=e.manuals[0].PercussionSignal(e.frame[:],e.percussion[0].harmonic)
  upper+=e.percussion[0].NextHarmonic(signal)
 }
 if e.percussion[1].harmonic==0 {
  lower+=e.percussion[1].Next(e.frame[:])
 } else {
  signal:=e.manuals[1].PercussionSignal(e.frame[:],e.percussion[1].harmonic)
  lower+=e.percussion[1].NextHarmonic(signal)
 }
 return electrical.ManualPaths{Upper:upper,Lower:lower}
}

// Next retains the legacy diagnostic mono mix.
func (e *TimedEngine) Next() float64 { return e.NextPaths().Sum() }

// RenderPaths writes separate manual samples without allocating.
func (e *TimedEngine) RenderPaths(upper,lower []float64) {
 if len(upper)!=len(lower) { panic("manual output buffers must have equal length") }
 for i:=range upper {
  paths:=e.NextPaths()
  upper[i]=paths.Upper
  lower[i]=paths.Lower
 }
}

func (e *TimedEngine) Render(dst []float64) {
 for i:=range dst { dst[i]=e.Next() }
}

// SetPercussion configures one manual's experimental tonewheel percussion.
func (e *TimedEngine) SetPercussion(manual int,wheel int,level,decay float64) error {
 if manual<0||manual>=len(e.percussion) {return fmt.Errorf("manual index %d outside 0..1",manual)}
 err:=e.percussion[manual].Configure(wheel,level,decay)
 if err==nil && manual==0 {e.manuals[0].SetPercussionDrawbarCut(level>0)}
 return err
}

func (e *TimedEngine) SetHarmonicPercussion(manual,harmonic int,level,decay float64) error {
 if manual<0||manual>1 {return fmt.Errorf("manual index outside 0..1")}
 err:=e.percussion[manual].ConfigureHarmonic(harmonic,level,decay)
 if err==nil && manual==0 {e.manuals[0].SetPercussionDrawbarCut(level>0)}
 return err
}

// SetTimedPercussion configures a key-dependent harmonic with an
// exponential time constant in seconds, independent of sample rate.
func (e *TimedEngine) SetTimedPercussion(manual,harmonic int,level,seconds float64) error {
 if math.IsNaN(seconds)||math.IsInf(seconds,0)||seconds<=0 {
  return fmt.Errorf("percussion time constant must be positive and finite")
 }
 decay:=math.Exp(-1/(e.sampleRate*seconds))
 return e.SetHarmonicPercussion(manual,harmonic,level,decay)
}
