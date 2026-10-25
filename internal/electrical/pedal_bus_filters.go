package electrical

import "fmt"

// ExperimentalPedalBusFilters applies four *independent* series-RL
// stages to four bus voltages. This is a laboratory topology, not
// Hammond B3 pedal reactor wiring.
type ExperimentalPedalBusFilters struct {
 stages [4]*PedalRLStage
}

func NewExperimentalPedalBusFilters(sampleRate float64, configs [4]PedalRLConfig) (*ExperimentalPedalBusFilters,error) {
 f:=&ExperimentalPedalBusFilters{}
 for i,c:=range configs {
  stage,err:=NewPedalRLStage(sampleRate,c)
  if err!=nil { return nil,fmt.Errorf("pedal filter bus %d: %w",i,err) }
  f.stages[i]=stage
 }
 return f,nil
}

// Next checks all inputs before updating any stage, so rejected frames
// cannot advance filter history.
func (f *ExperimentalPedalBusFilters) Next(in [4]float64)(out [4]float64,err error) {
 for i,v:=range in {
  if !finite(v) { return out,fmt.Errorf("pedal filter bus %d: nonfinite input",i) }
 }
 for i,v:=range in {
  out[i],err=f.stages[i].Next(v)
  if err!=nil { return [4]float64{},err }
 }
 return out,nil
}

func (f *ExperimentalPedalBusFilters) Reset() {
 for _,s:=range f.stages { s.Reset() }
}
