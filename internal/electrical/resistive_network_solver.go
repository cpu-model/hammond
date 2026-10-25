package electrical

import (
 "fmt"
 "math"
)

// ResistiveBranch is a two-terminal resistor. An empty node denotes ground.
// Values must be explicitly supplied; this solver has no B3 defaults.
type ResistiveBranch struct {
 A, B string
 Ohms float64
}

// SolveResistiveNetwork solves a linear resistive network using nodal
// analysis. Fixed potentials model ideal voltage sources, including ground
// as the empty node. Floating or singular networks are rejected.
func SolveResistiveNetwork(branches []ResistiveBranch, fixed map[string]float64) (map[string]float64,error) {
 known:=map[string]float64{"":0}
 for n,v:=range fixed {
  if n==""&&v!=0 {return nil,fmt.Errorf("ground potential must be zero")}
  if math.IsNaN(v)||math.IsInf(v,0) {return nil,fmt.Errorf("nonfinite fixed potential at %q",n)}
  known[n]=v
 }
 nodes:=map[string]bool{}
 for i,b:=range branches {
  if b.A==b.B||b.Ohms<=0||math.IsNaN(b.Ohms)||math.IsInf(b.Ohms,0) {
   return nil,fmt.Errorf("invalid resistor %d",i)
  }
  if _,ok:=known[b.A];!ok {nodes[b.A]=true}
  if _,ok:=known[b.B];!ok {nodes[b.B]=true}
 }
 order:=make([]string,0,len(nodes))
 for n:=range nodes {order=append(order,n)}
 // Sort for deterministic elimination and error reporting.
 sortStrings(order)
 index:=map[string]int{}
 for i,n:=range order {index[n]=i}
 size:=len(order)
 matrix:=make([][]float64,size)
 for i:=range matrix {matrix[i]=make([]float64,size+1)}
 for _,b:=range branches {
  g:=1/b.Ohms
  for _,pair:=range [][2]string{{b.A,b.B},{b.B,b.A}} {
   i,ok:=index[pair[0]]
   if !ok {continue}
   matrix[i][i]+=g
   if j,ok:=index[pair[1]];ok {matrix[i][j]-=g} else {
    matrix[i][size]+=g*known[pair[1]]
   }
  }
 }
 for k:=0;k<size;k++ {
  pivot:=k
  for i:=k+1;i<size;i++ {if math.Abs(matrix[i][k])>math.Abs(matrix[pivot][k]) {pivot=i}}
  if math.Abs(matrix[pivot][k])<1e-15 {return nil,fmt.Errorf("floating or singular resistive network")}
  matrix[k],matrix[pivot]=matrix[pivot],matrix[k]
  divisor:=matrix[k][k]
  for j:=k;j<=size;j++ {matrix[k][j]/=divisor}
  for i:=0;i<size;i++ {
   if i==k {continue}
   factor:=matrix[i][k]
   for j:=k;j<=size;j++ {matrix[i][j]-=factor*matrix[k][j]}
  }
 }
 result:=make(map[string]float64,len(known)+size)
 for n,v:=range known {result[n]=v}
 for i,n:=range order {result[n]=matrix[i][size]}
 return result,nil
}

func sortStrings(a []string) {
 for i:=1;i<len(a);i++ {
  for j:=i;j>0&&a[j]<a[j-1];j-- {a[j],a[j-1]=a[j-1],a[j]}
 }
}
