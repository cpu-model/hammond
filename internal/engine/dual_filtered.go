package engine

import (
 "github.com/cpu-model/hammond/internal/electrical"
 "github.com/cpu-model/hammond/internal/registration"
)

// DualFilteredTimedEngine independently processes the two manual paths.
// This is a controlled experimental response, NOT a B3 T1/T2 model.
type DualFilteredTimedEngine struct {
 source *TimedEngine
 upper *electrical.Lowpass
 lower *electrical.Lowpass
}

// Both cutoff values are explicit experimental parameters. No
// transformer, bus loading, or AO-28 characteristics are inferred.
func NewDualFilteredTimed(sampleRate,upperCutoff,lowerCutoff float64,upperRoute,lowerRoute registration.Routing,upperTiming,lowerTiming registration.ContactTiming) (*DualFilteredTimedEngine,error) {
 source,err:=NewTimed(sampleRate,upperRoute,lowerRoute,upperTiming,lowerTiming)
 if err!=nil { return nil,err }
 u,err:=electrical.NewLowpass(sampleRate,upperCutoff)
 if err!=nil { return nil,err }
 l,err:=electrical.NewLowpass(sampleRate,lowerCutoff)
 if err!=nil { return nil,err }
 return &DualFilteredTimedEngine{source:source,upper:u,lower:l},nil
}

func (e *DualFilteredTimedEngine) Manual(index int) (*registration.TimedManual,error) {
 return e.source.Manual(index)
}

func (e *DualFilteredTimedEngine) NextPaths() electrical.ManualPaths {
 p:=e.source.NextPaths()
 return electrical.ManualPaths{Upper:e.upper.Next(p.Upper),Lower:e.lower.Next(p.Lower)}
}

func (e *DualFilteredTimedEngine) RenderPaths(upper,lower []float64) {
 if len(upper)!=len(lower) { panic("manual output buffers must have equal length") }
 for i:=range upper {
  p:=e.NextPaths()
  upper[i],lower[i]=p.Upper,p.Lower
 }
}

// Next is a diagnostic sum only; no physical B3 mixing is implied.
func (e *DualFilteredTimedEngine) Next() float64 { return e.NextPaths().Sum() }

func (e *DualFilteredTimedEngine) ResetFilters() {
 e.upper.Reset()
 e.lower.Reset()
}
