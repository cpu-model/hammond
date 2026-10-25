package engine

import (
 "fmt"

 "github.com/cpu-model/hammond/internal/generator"
 "github.com/cpu-model/hammond/internal/registration"
)

// TimedEngine renders two independently switched manuals against the same
// continuously rotating 91-output mechanical generator.
type TimedEngine struct {
 generator *generator.MechanicalBank
 manuals [2]*registration.TimedManual
 frame [91]float64
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
 return &TimedEngine{generator:bank,manuals:[2]*registration.TimedManual{u,l}},nil
}

func (e *TimedEngine) Manual(index int) (*registration.TimedManual,error) {
 if index<0 || index>1 { return nil,fmt.Errorf("manual index %d outside 0..1",index) }
 return e.manuals[index],nil
}

func (e *TimedEngine) Next() float64 {
 e.generator.Next(e.frame[:])
 return e.manuals[0].Mix(e.frame[:])+e.manuals[1].Mix(e.frame[:])
}

func (e *TimedEngine) Render(dst []float64) {
 for i:=range dst { dst[i]=e.Next() }
}
