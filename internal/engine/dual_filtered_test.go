package engine

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
)

func TestDualFilteredPathsIndependent(t *testing.T) {
 route:=registration.B3ManualRouting()
 const rate=48000.0
 const upperCutoff=900.0
 const lowerCutoff=4200.0
 raw,err:=NewTimed(rate,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil { t.Fatal(err) }
 filtered,err:=NewDualFilteredTimed(rate,upperCutoff,lowerCutoff,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil { t.Fatal(err) }
 for _,e:=range []interface{Manual(int)(*registration.TimedManual,error)}{raw,filtered} {
  u,_:=e.Manual(0)
  l,_:=e.Manual(1)
  if err:=u.SetDrawbar(0,8);err!=nil { t.Fatal(err) }
  if err:=l.SetDrawbar(0,8);err!=nil { t.Fatal(err) }
  if err:=u.Key(36,true);err!=nil { t.Fatal(err) }
  if err:=l.Key(24,true);err!=nil { t.Fatal(err) }
 }
 au:=math.Exp(-2*math.Pi*upperCutoff/rate)
 al:=math.Exp(-2*math.Pi*lowerCutoff/rate)
 var wantU,wantL float64
 for i:=0;i<512;i++ {
  p:=raw.NextPaths()
  wantU=(1-au)*p.Upper+au*wantU
  wantL=(1-al)*p.Lower+al*wantL
  got:=filtered.NextPaths()
  if math.Abs(got.Upper-wantU)>1e-12||math.Abs(got.Lower-wantL)>1e-12 {
   t.Fatalf("sample %d: independent filter recurrence mismatch",i)
  }
 }
}

func TestDualFilteredNoAllocations(t *testing.T) {
 e,err:=NewDualFilteredTimed(48000,1000,2000,registration.Routing{},registration.Routing{},registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil { t.Fatal(err) }
 upper,lower:=make([]float64,128),make([]float64,128)
 if n:=testing.AllocsPerRun(100,func(){e.RenderPaths(upper,lower)});n!=0 {
  t.Fatalf("allocations per render: %v",n)
 }
}

func TestDualFilteredRejectsInvalidLowerCutoff(t *testing.T) {
 _,err:=NewDualFilteredTimed(48000,1000,24000,registration.Routing{},registration.Routing{},registration.ContactTiming{},registration.ContactTiming{})
 if err==nil { t.Fatal("expected cutoff error") }
}
