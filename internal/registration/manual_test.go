package registration

import (
 "math"
 "testing"
)

func TestManualSharedWheelAndKeyRelease(t *testing.T) {
 var route Routing
 route[0][0]=46
 route[1][0]=46 // two contacts use the same physical signal
 m,err:=NewManual(route)
 if err!=nil { t.Fatal(err) }
 if err=m.SetDrawbar(0,8);err!=nil { t.Fatal(err) }
 frame:=make([]float64,91)
 frame[45]=0.25
 if err=m.Key(0,true);err!=nil { t.Fatal(err) }
 if got:=m.Mix(frame);got!=0.25 { t.Fatalf("one key: %v",got) }
 if err=m.Key(1,true);err!=nil { t.Fatal(err) }
 if got:=m.Mix(frame);got!=0.5 { t.Fatalf("two contacts: %v",got) }
 if err=m.Key(0,false);err!=nil { t.Fatal(err) }
 if got:=m.Mix(frame);got!=0.25 { t.Fatalf("release: %v",got) }
}

func TestManualNineDrawbars(t *testing.T) {
 var route Routing
 for d:=0;d<Harmonics;d++ { route[0][d]=uint8(d+1) }
 m,_:=NewManual(route)
 frame:=make([]float64,91)
 for i:=0;i<9;i++ { frame[i]=1; if err:=m.SetDrawbar(i,8);err!=nil { t.Fatal(err) } }
 _=m.Key(0,true)
 if got:=m.Mix(frame);math.Abs(got-9)>1e-12 { t.Fatalf("nine contacts: %v",got) }
}

func TestManualRejectsInvalid(t *testing.T) {
 var route Routing
 route[0][0]=92
 if _,err:=NewManual(route);err==nil { t.Fatal("invalid wheel accepted") }
 m,_:=NewManual(Routing{})
 if err:=m.Key(61,true);err==nil { t.Fatal("invalid key accepted") }
 if err:=m.SetDrawbar(0,9);err==nil { t.Fatal("invalid stop accepted") }
}

func TestManualMixNoAllocations(t *testing.T) {
 var route Routing
 route[0][0]=46
 m,_:=NewManual(route)
 _=m.Key(0,true)
 _=m.SetDrawbar(0,8)
 frame:=make([]float64,91)
 if n:=testing.AllocsPerRun(100,func(){ _=m.Mix(frame) });n!=0 { t.Fatalf("allocations: %v",n) }
}
