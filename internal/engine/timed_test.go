package engine

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
)

func TestTimedEngineStaggeredContact(t *testing.T) {
 var route registration.Routing
 route[0][0]=46
 var timing registration.ContactTiming
 timing.Make[0]=2
 e,err:=NewTimed(48000,route,registration.Routing{},timing,registration.ContactTiming{})
 if err!=nil { t.Fatal(err) }
 m,_:=e.Manual(0)
 _=m.SetDrawbar(0,8)
 _=m.Key(0,true)
 for n:=0;n<8;n++ {
  got:=e.Next()
  want:=0.0
  if n>=2 { want=math.Sin(2*math.Pi*440*float64(n)/48000) }
  if math.Abs(got-want)>1e-10 { t.Fatalf("sample %d: got %.12f want %.12f",n,got,want) }
 }
}

func TestTimedEngineSharedGenerator(t *testing.T) {
 var route registration.Routing
 route[0][0]=46
 e,err:=NewTimed(48000,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil { t.Fatal(err) }
 for i:=0;i<2;i++ {
  m,_:=e.Manual(i)
  _=m.SetDrawbar(0,8)
  _=m.Key(0,true)
 }
 for n:=0;n<100;n++ {
  want:=2*math.Sin(2*math.Pi*440*float64(n)/48000)
  if got:=e.Next();math.Abs(got-want)>1e-10 { t.Fatalf("sample %d: got %v want %v",n,got,want) }
 }
}

func TestTimedEngineNoRenderAllocations(t *testing.T) {
 e,err:=NewTimed(48000,registration.Routing{},registration.Routing{},registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil { t.Fatal(err) }
 buf:=make([]float64,64)
 if n:=testing.AllocsPerRun(100,func(){e.Render(buf)});n!=0 { t.Fatalf("allocations: %v",n) }
}
