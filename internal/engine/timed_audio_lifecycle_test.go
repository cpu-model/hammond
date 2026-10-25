package engine

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
)

// The timed audio path must respond to a key press and return to silence
// after release, without stopping the continuously running generator.
func TestTimedManualAudiblePressAndRelease(t *testing.T) {
 route:=registration.B3ManualRouting()
 e,err:=NewTimed(48000,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil {t.Fatal(err)}
 upper,err:=e.Manual(0)
 if err!=nil {t.Fatal(err)}
 if err:=upper.SetDrawbar(0,8);err!=nil {t.Fatal(err)}
 for i:=0;i<128;i++ {
  if v:=e.Next();v!=0 {t.Fatalf("silent engine sample %d = %g",i,v)}
 }
 if err:=upper.Key(36,true);err!=nil {t.Fatal(err)}
 peak:=0.0
 for i:=0;i<2048;i++ {
  v:=e.Next()
  if math.IsNaN(v)||math.IsInf(v,0) {t.Fatalf("nonfinite pressed sample %d",i)}
  if math.Abs(v)>peak {peak=math.Abs(v)}
 }
 if peak==0 {t.Fatal("pressed key produced no audio")}
 if err:=upper.Key(36,false);err!=nil {t.Fatal(err)}
 for i:=0;i<128;i++ {
  if v:=e.Next();v!=0 {t.Fatalf("released key sample %d = %g",i,v)}
 }
}
