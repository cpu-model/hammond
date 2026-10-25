package engine

import (
 "errors"
 "fmt"

 "github.com/cpu-model/hammond/internal/generator"
 "github.com/cpu-model/hammond/internal/registration"
)

// Engine owns one continuously running generator and two manuals.
// Both manuals receive the same generator frame at each sample.
type Engine struct {
 generator *generator.MechanicalBank
 manuals [2]*registration.Manual
 frame [91]float64
}

func New(sampleRate float64, upper, lower registration.Routing) (*Engine,error) {
 bank,err:=generator.NewMechanicalBank(sampleRate)
 if err!=nil { return nil,err }
 u,err:=registration.NewManual(upper)
 if err!=nil { return nil,fmt.Errorf("upper manual: %w",err) }
 l,err:=registration.NewManual(lower)
 if err!=nil { return nil,fmt.Errorf("lower manual: %w",err) }
 return &Engine{generator:bank,manuals:[2]*registration.Manual{u,l}},nil
}

func (e *Engine) Manual(index int) (*registration.Manual,error) {
 if index<0 || index>=len(e.manuals) { return nil,errors.New("manual index must be 0 (upper) or 1 (lower)") }
 return e.manuals[index],nil
}

// Next advances the mechanical generator exactly once per audio sample.
// Mixing and output scaling are provisional, pending electrical modelling.
func (e *Engine) Next() float64 {
 e.generator.Next(e.frame[:])
 return e.manuals[0].Mix(e.frame[:])+e.manuals[1].Mix(e.frame[:])
}

// Render writes one sample at a time without allocating.
func (e *Engine) Render(dst []float64) {
 for i:=range dst { dst[i]=e.Next() }
}
