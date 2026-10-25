package engine

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
)

func TestTimedSeparatePathsMatchDiagnosticMix(t *testing.T) {
 route:=registration.B3ManualRouting()
 var timing registration.ContactTiming
 separate,err:=NewTimed(48000,route,route,timing,timing)
 if err!=nil { t.Fatal(err) }
 mixed,err:=NewTimed(48000,route,route,timing,timing)
 if err!=nil { t.Fatal(err) }
 for _,e:=range []*TimedEngine{separate,mixed} {
  upper,_:=e.Manual(0)
  lower,_:=e.Manual(1)
  if err:=upper.SetDrawbar(0,8);err!=nil { t.Fatal(err) }
  if err:=lower.SetDrawbar(1,8);err!=nil { t.Fatal(err) }
  if err:=upper.Key(36,true);err!=nil { t.Fatal(err) }
  if err:=lower.Key(24,true);err!=nil { t.Fatal(err) }
 }
 for i:=0;i<512;i++ {
  paths:=separate.NextPaths()
  got:=paths.Sum()
  want:=mixed.Next()
  if math.Abs(got-want)>1e-12 { t.Fatalf("sample %d: got %g want %g",i,got,want) }
 }
}

func TestTimedRenderPathsNoAllocations(t *testing.T) {
 e,err:=NewTimed(48000,registration.Routing{},registration.Routing{},registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil { t.Fatal(err) }
 upper:=make([]float64,128)
 lower:=make([]float64,128)
 if n:=testing.AllocsPerRun(100,func(){e.RenderPaths(upper,lower)});n!=0 {
  t.Fatalf("unexpected allocations: %v",n)
 }
}
