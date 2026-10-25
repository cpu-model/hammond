package registration

import "testing"

func TestContactStaggerAndRelease(t *testing.T) {
 var route Routing
 route[0][0]=46
 route[0][1]=47
 timing:=ContactTiming{}
 timing.Make[1]=2
 timing.Break[0]=1
 timing.Break[1]=3
 m,err:=NewTimedManual(route,timing)
 if err!=nil { t.Fatal(err) }
 _=m.SetDrawbar(0,8)
 _=m.SetDrawbar(1,8)
 frame:=make([]float64,91)
 frame[45]=1
 frame[46]=2
 _=m.Key(0,true)
 for i,want:=range []float64{1,1,3} {
  if got:=m.Mix(frame);got!=want { t.Fatalf("make sample %d: got %v want %v",i,got,want) }
 }
 _=m.Key(0,false)
 for i,want:=range []float64{3,2,2,0} {
  if got:=m.Mix(frame);got!=want { t.Fatalf("break sample %d: got %v want %v",i,got,want) }
 }
}

func TestContactReversalSupersedesPending(t *testing.T) {
 var route Routing
 route[0][0]=46
 timing:=ContactTiming{}
 timing.Make[0]=10
 m,_:=NewTimedManual(route,timing)
 _=m.SetDrawbar(0,8)
 frame:=make([]float64,91)
 frame[45]=1
 _=m.Key(0,true)
 for i:=0;i<3;i++ { if m.Mix(frame)!=0 { t.Fatal("early make") } }
 _=m.Key(0,false)
 for i:=0;i<20;i++ { if m.Mix(frame)!=0 { t.Fatal("stale make event") } }
}

func TestTimedManualNoAllocations(t *testing.T) {
 m,_:=NewTimedManual(Routing{},ContactTiming{})
 frame:=make([]float64,91)
 if n:=testing.AllocsPerRun(100,func(){m.Mix(frame)});n!=0 { t.Fatalf("allocations: %v",n) }
}
