package registration

import (
 "math"
 "testing"
)

func TestNinthDrawbarRampMonotonicAndReversible(t *testing.T) {
 var route Routing
 route[0][8]=1
 m,err:=NewTimedManual(route,ContactTiming{})
 if err!=nil{t.Fatal(err)}
 if err:=m.SetDrawbar(8,8);err!=nil{t.Fatal(err)}
 if err:=m.Key(0,true);err!=nil{t.Fatal(err)}
 m.SetNinthDrawbarTransition(10)
 frame:=make([]float64,91)
 frame[0]=1
 if got:=m.Mix(frame);math.Abs(got-1)>1e-12 {t.Fatalf("initial gain %v",got)}
 m.SetPercussionDrawbarCut(true)
 previous:=1.0
 for i:=0;i<4;i++ {
  got:=m.Mix(frame)
  if got>previous||got<0 {t.Fatalf("non-monotonic fade out at %d: %v",i,got)}
  previous=got
 }
 if math.Abs(previous-0.6)>1e-12 {t.Fatalf("four samples should leave 0.6: %v",previous)}
 m.SetPercussionDrawbarCut(false)
 for i:=0;i<4;i++ {
  got:=m.Mix(frame)
  if got<previous||got>1 {t.Fatalf("non-monotonic fade in at %d: %v",i,got)}
  previous=got
 }
 if math.Abs(previous-1)>1e-12 {t.Fatalf("restoration not complete: %v",previous)}
 m.SetPercussionDrawbarCut(true)
 for i:=0;i<10;i++ {m.Mix(frame)}
 if got:=m.Mix(frame);math.Abs(got)>1e-12 {t.Fatalf("fade out did not finish: %v",got)}
}

func TestNinthDrawbarRampIndependentOfOtherHarmonics(t *testing.T) {
 var route Routing
 route[0][2]=2
 route[0][8]=1
 m,err:=NewTimedManual(route,ContactTiming{})
 if err!=nil{t.Fatal(err)}
 _=m.SetDrawbar(2,8)
 _=m.SetDrawbar(8,8)
 _=m.Key(0,true)
 frame:=make([]float64,91)
 frame[0]=1
 frame[1]=2
 m.SetNinthDrawbarTransition(2)
 m.SetPercussionDrawbarCut(true)
 m.Mix(frame)
 if got:=m.Mix(frame);math.Abs(got-2)>1e-12 {t.Fatalf("other harmonic affected: %v",got)}
}
