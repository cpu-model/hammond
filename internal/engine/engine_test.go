package engine

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
)

func TestTwoManualsShareOneGenerator(t *testing.T) {
 var upper,lower registration.Routing
 upper[0][0]=46
 lower[0][0]=46
 e,err:=New(48000,upper,lower)
 if err!=nil { t.Fatal(err) }
 for i:=0;i<2;i++ {
  m,_:=e.Manual(i)
  if err:=m.Key(0,true);err!=nil { t.Fatal(err) }
  if err:=m.SetDrawbar(0,8);err!=nil { t.Fatal(err) }
 }
 for n:=0;n<1000;n++ {
  got:=e.Next()
  want:=2*math.Sin(2*math.Pi*440*float64(n)/48000)
  if math.Abs(got-want)>1e-9 { t.Fatalf("sample %d: %.12f != %.12f",n,got,want) }
 }
}

func TestKeyReleaseDoesNotRestartWheels(t *testing.T) {
 var route registration.Routing
 route[0][0]=46
 e,_:=New(48000,route,registration.Routing{})
 m,_:=e.Manual(0)
 _=m.SetDrawbar(0,8)
 _=m.Key(0,true)
 for n:=0;n<100;n++ { e.Next() }
 _=m.Key(0,false)
 for n:=100;n<200;n++ {
  if got:=e.Next();got!=0 { t.Fatalf("released key at %d: %v",n,got) }
 }
 _=m.Key(0,true)
 got:=e.Next()
 want:=math.Sin(2*math.Pi*440*200/48000)
 if math.Abs(got-want)>1e-10 { t.Fatalf("phase restarted: %.12f vs %.12f",got,want) }
}

func TestEngineRenderNoAllocations(t *testing.T) {
 e,err:=New(48000,registration.Routing{},registration.Routing{})
 if err!=nil { t.Fatal(err) }
 dst:=make([]float64,64)
 if n:=testing.AllocsPerRun(100,func(){e.Render(dst)});n!=0 { t.Fatalf("render allocations: %v",n) }
}

func TestEngineInvalidManual(t *testing.T) {
 e,_:=New(48000,registration.Routing{},registration.Routing{})
 if _,err:=e.Manual(2);err==nil { t.Fatal("invalid manual accepted") }
}
