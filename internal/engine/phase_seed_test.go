package engine

import (
 "testing"
 "github.com/cpu-model/hammond/internal/registration"
)

func TestSeededWheelPhasesReproducible(t *testing.T) {
 newEngine:=func()*TimedEngine{
  e,err:=NewTimed(48000,registration.B3ManualRouting(),registration.B3ManualRouting(),registration.ContactTiming{},registration.ContactTiming{})
  if err!=nil{t.Fatal(err)}
  m,_:=e.Manual(0)
  if err:=m.Key(36,true);err!=nil{t.Fatal(err)}
  if err:=m.SetDrawbar(0,8);err!=nil{t.Fatal(err)}
  return e
 }
 a,b,c:=newEngine(),newEngine(),newEngine()
 if err:=a.SetPhaseSeed(17);err!=nil{t.Fatal(err)}
 if err:=b.SetPhaseSeed(17);err!=nil{t.Fatal(err)}
 if err:=c.SetPhaseSeed(18);err!=nil{t.Fatal(err)}
 different:=false
 for i:=0;i<100;i++ {
  x,y,z:=a.Next(),b.Next(),c.Next()
  if x!=y{t.Fatalf("same seed diverged at sample %d",i)}
  if x!=z{different=true}
 }
 if !different{t.Fatal("different seeds yielded identical audio")}
}
