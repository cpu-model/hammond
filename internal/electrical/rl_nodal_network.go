package electrical

import (
 "fmt"
 "math"
)

// InductiveBranch is a series inductor between two named circuit nodes.
// Its state is branch current in amperes, positive from A to B.
type InductiveBranch struct {
 A, B string
 Henry float64
}

// RLNodalNetwork uses backward Euler companion conductances to solve
// time-domain resistor/inductor circuits with fixed voltage source nodes.
// Component values are caller-supplied; no historical B3 values are assumed.
type RLNodalNetwork struct {
 SampleRate float64
 Resistors []ResistiveBranch
 Inductors []InductiveBranch
 currents []float64
}

func NewRLNodalNetwork(sampleRate float64,resistors []ResistiveBranch,inductors []InductiveBranch) (*RLNodalNetwork,error) {
 if !positiveFinite(sampleRate) {return nil,fmt.Errorf("invalid sample rate")}
 for i,b:=range inductors {
  if b.A==b.B||!positiveFinite(b.Henry) {return nil,fmt.Errorf("invalid inductor %d",i)}
 }
 for i,b:=range resistors {
  if b.A==b.B||!positiveFinite(b.Ohms) {return nil,fmt.Errorf("invalid resistor %d",i)}
 }
 return &RLNodalNetwork{SampleRate:sampleRate,Resistors:append([]ResistiveBranch(nil),resistors...),Inductors:append([]InductiveBranch(nil),inductors...),currents:make([]float64,len(inductors))},nil
}

// Next solves one sample and commits branch current only on success.
func (n *RLNodalNetwork) Next(fixed map[string]float64) (map[string]float64,error) {
 if n==nil {return nil,fmt.Errorf("nil RL network")}
 dt:=1/n.SampleRate
 branches:=make([]ResistiveBranch,0,len(n.Resistors)+len(n.Inductors))
 branches=append(branches,n.Resistors...)
 // BE inductor companion: i[n] = i[n-1] + dt/L*(Va[n]-Vb[n]).
 // The equivalent current source is handled via an augmented nodal RHS.
 for _,b:=range n.Inductors {
  branches=append(branches,ResistiveBranch{A:b.A,B:b.B,Ohms:b.Henry/dt})
 }
 result,err:=solveResistiveWithCurrentSources(branches,fixed,n.Inductors,n.currents)
 if err!=nil {return nil,err}
 updated:=make([]float64,len(n.currents))
 for i,b:=range n.Inductors {
  updated[i]=n.currents[i]+dt/b.Henry*(result[b.A]-result[b.B])
  if math.IsNaN(updated[i])||math.IsInf(updated[i],0) {return nil,fmt.Errorf("nonfinite inductor current")}
 }
 copy(n.currents,updated)
 return result,nil
}

func (n *RLNodalNetwork) Reset() {
 if n!=nil {for i:=range n.currents {n.currents[i]=0}}
}

func solveResistiveWithCurrentSources(branches []ResistiveBranch,fixed map[string]float64,inductors []InductiveBranch,previous []float64) (map[string]float64,error) {
 // Current source injections are represented using a superposition of
 // nodal right-hand sides. Delegate to the shared linear solver below.
 injections:=map[string]float64{}
 for i,b:=range inductors {
  injections[b.A]-=previous[i]
  injections[b.B]+=previous[i]
 }
 return solveResistiveNodal(branches,fixed,injections)
}
