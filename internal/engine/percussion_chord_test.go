package engine

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
)

func TestHarmonicPercussionFollowsHeldKeysWithoutRetrigger(t *testing.T) {
 route:=registration.B3ManualRouting()
 e,err:=NewTimed(48000,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil{t.Fatal(err)}
 if err:=e.SetHarmonicPercussion(0,3,0.8,0.5);err!=nil{t.Fatal(err)}
 manual,_:=e.Manual(0)
 if err:=manual.Key(36,true);err!=nil{t.Fatal(err)}
 first:=e.NextPaths()
 if first.Upper==0{t.Fatal("initial percussion silent")}
 // After one sample the envelope must decay, not restart when a
 // second key joins the held chord.
 if err:=manual.Key(40,true);err!=nil{t.Fatal(err)}
 e.NextPaths()
 if math.Abs(e.percussion[0].envelope-0.2)>1e-12 {
  t.Fatalf("legato note retriggered envelope: %v",e.percussion[0].envelope)
 }
 if err:=manual.Key(36,false);err!=nil{t.Fatal(err)}
 e.NextPaths()
 if math.Abs(e.percussion[0].envelope-0.1)>1e-12 {
  t.Fatalf("partial release changed envelope: %v",e.percussion[0].envelope)
 }
 if err:=manual.Key(40,false);err!=nil{t.Fatal(err)}
 e.NextPaths()
 if err:=manual.Key(40,true);err!=nil{t.Fatal(err)}
 e.NextPaths()
 if math.Abs(e.percussion[0].envelope-0.4)>1e-12 {
  t.Fatalf("new phrase did not retrigger envelope: %v",e.percussion[0].envelope)
 }
}

func TestPercussionSignalTracksKeysAndIgnoresDrawbars(t *testing.T) {
 var route registration.Routing
 route[0][4]=10
 route[1][4]=20
 m,err:=registration.NewTimedManual(route,registration.ContactTiming{})
 if err!=nil{t.Fatal(err)}
 frame:=make([]float64,91)
 frame[9]=0.3
 frame[19]=0.7
 if err:=m.Key(0,true);err!=nil{t.Fatal(err)}
 if got:=m.PercussionSignal(frame,3);math.Abs(got-0.3)>1e-12 {t.Fatalf("first key: %v",got)}
 if err:=m.Key(1,true);err!=nil{t.Fatal(err)}
 if got:=m.PercussionSignal(frame,3);math.Abs(got-1)>1e-12 {t.Fatalf("chord: %v",got)}
 if err:=m.Key(0,false);err!=nil{t.Fatal(err)}
 if got:=m.PercussionSignal(frame,3);math.Abs(got-0.7)>1e-12 {t.Fatalf("remaining key: %v",got)}
}
