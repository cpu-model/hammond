package engine

import (
 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

// FilteredTimedEngine is an experimental signal-chain variant.
// It filters the sum of both manuals after contact mixing. This
// insertion point is NOT claimed to match B3 circuitry.
type FilteredTimedEngine struct {
 source *TimedEngine
 filter *electrical.Lowpass
}

func NewFilteredTimed(sampleRate, cutoffHz float64, upper,lower registration.Routing, upperTiming,lowerTiming registration.ContactTiming) (*FilteredTimedEngine,error) {
 source,err:=NewTimed(sampleRate,upper,lower,upperTiming,lowerTiming)
 if err!=nil { return nil,err }
 filter,err:=electrical.NewLowpass(sampleRate,cutoffHz)
 if err!=nil { return nil,err }
 return &FilteredTimedEngine{source:source,filter:filter},nil
}

func (e *FilteredTimedEngine) Manual(index int) (*registration.TimedManual,error) {
 return e.source.Manual(index)
}

func (e *FilteredTimedEngine) Next() float64 {
 return e.filter.Next(e.source.Next())
}

func (e *FilteredTimedEngine) Render(dst []float64) {
 for i:=range dst { dst[i]=e.Next() }
}

// ResetFilter clears the experimental filter memory but does not
// reset generator phases or key-contact state.
func (e *FilteredTimedEngine) ResetFilter() { e.filter.Reset() }
