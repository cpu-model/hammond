package engine

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
)

func TestFilteredTimedMatchesReferenceRecurrence(t *testing.T) {
 var route registration.Routing
 route[0][0]=46
 const rate=48000.0
 const cutoff=1200.0
 var timing registration.ContactTiming
 timing.Make[0]=4
 raw,err:=NewTimed(rate,route,registration.Routing{},timing,registration.ContactTiming{})
 if err!=nil { t.Fatal(err) }
 filtered,err:=NewFilteredTimed(rate,cutoff,route,registration.Routing{},timing,registration.ContactTiming{})
 if err!=nil { t.Fatal(err) }
 rm,_:=raw.Manual(0)
 fm,_:=filtered.Manual(0)
 for _,m:=range []*registration.TimedManual{rm,fm} {
  _=m.SetDrawbar(0,8)
  _=m.Key(0,true)
 }
 a:=math.Exp(-2*math.Pi*cutoff/rate)
 var expected float64
 for i:=0;i<300;i++ {
  expected=(1-a)*raw.Next()+a*expected
  if got:=filtered.Next();math.Abs(got-expected)>1e-12 {
   t.Fatalf("sample %d: got %.15f want %.15f",i,got,expected)
  }
 }
}

func TestFilteredTimedNoAllocations(t *testing.T) {
 e,err:=NewFilteredTimed(48000,1000,registration.Routing{},registration.Routing{},registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil { t.Fatal(err) }
 dst:=make([]float64,128)
 if n:=testing.AllocsPerRun(100,func(){e.Render(dst)});n!=0 {
  t.Fatalf("allocations: %v",n)
 }
}

func TestFilteredTimedRejectsBadCutoff(t *testing.T) {
 if _,err:=NewFilteredTimed(48000,24000,registration.Routing{},registration.Routing{},registration.ContactTiming{},registration.ContactTiming{});err==nil {
  t.Fatal("expected invalid cutoff")
 }
}
