package engine

import (
 "math"
 "testing"
)

func TestPercussionSharedWheelEnvelope(t *testing.T) {
 var p Percussion
 if err:=p.Configure(12,0.8,0.5);err!=nil{t.Fatal(err)}
 frame:=make([]float64,91)
 frame[11]=0.75
 if got:=p.Next(frame);got!=0 {t.Fatalf("inactive percussion: %v",got)}
 p.Trigger()
 for i,want:=range []float64{0.6,0.3,0.15} {
  if got:=p.Next(frame);math.Abs(got-want)>1e-12 {t.Fatalf("sample %d: got %v want %v",i,got,want)}
 }
 frame[11]=0
 if got:=p.Next(frame);got!=0 {t.Fatalf("percussion not using wheel: %v",got)}
}

func TestPercussionRejectsInvalidSettings(t *testing.T) {
 var p Percussion
 for _,tc:=range []struct{wheel int;level,decay float64}{
  {0,0.5,0.9},{92,0.5,0.9},{1,-1,0.9},{1,0.5,1},{1,math.NaN(),0.9},
 } {
  if err:=p.Configure(tc.wheel,tc.level,tc.decay);err==nil {t.Fatalf("accepted %+v",tc)}
 }
}
