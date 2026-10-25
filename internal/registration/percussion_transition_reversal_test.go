package registration

import (
 "math"
 "testing"
)

func TestNinthDrawbarFastToggleHasBoundedStep(t *testing.T) {
 var route Routing
 route[0][8]=1
 m,err:=NewTimedManual(route,ContactTiming{})
 if err!=nil{t.Fatal(err)}
 _=m.SetDrawbar(8,8)
 _=m.Key(0,true)
 m.SetNinthDrawbarTransition(144)
 frame:=make([]float64,91)
 frame[0]=1
 previous:=m.Mix(frame)
 for i:=0;i<400;i++ {
  if i%17==0 {m.SetPercussionDrawbarCut((i/17)%2==0)}
  current:=m.Mix(frame)
  if math.Abs(current-previous)>1.0/144+1e-12 {
   t.Fatalf("abrupt gain step at sample %d: %.9f",i,math.Abs(current-previous))
  }
  if current<0||current>1 {t.Fatalf("gain out of bounds: %v",current)}
  previous=current
 }
}
