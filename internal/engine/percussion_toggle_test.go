package engine

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
)

func TestPercussionToggleHeldKeyDoesNotTriggerEnvelope(t *testing.T) {
 route:=registration.B3ManualRouting()
 e,err:=NewTimed(48000,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil{t.Fatal(err)}
 m,_:=e.Manual(0)
 if err:=m.Key(36,true);err!=nil{t.Fatal(err)}
 e.NextPaths()
 if err:=e.SetHarmonicPercussion(0,3,0.5,0.9);err!=nil{t.Fatal(err)}
 if got:=e.percussion[0].envelope;got!=0{t.Fatalf("envelope started during toggle: %v",got)}
 e.NextPaths()
 if got:=e.percussion[0].envelope;got!=0{t.Fatalf("held key triggered percussion: %v",got)}
 if err:=m.Key(36,false);err!=nil{t.Fatal(err)}
 e.NextPaths()
 if err:=m.Key(36,true);err!=nil{t.Fatal(err)}
 e.NextPaths()
 if got:=e.percussion[0].envelope;math.Abs(got-0.45)>1e-12 {t.Fatalf("fresh key press not triggered: %v",got)}
}

func TestPercussionTogglePreservesLowerManualSignal(t *testing.T) {
 var route registration.Routing
 route[0][8]=1
 e,err:=NewTimed(48000,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil{t.Fatal(err)}
 upper,_:=e.Manual(0)
 lower,_:=e.Manual(1)
 _=upper.SetDrawbar(8,8)
 _=lower.SetDrawbar(8,8)
 _=upper.Key(0,true)
 _=lower.Key(0,true)
 e.NextPaths()
 if err:=e.SetHarmonicPercussion(0,3,0.5,0.9);err!=nil{t.Fatal(err)}
 for i:=0;i<150;i++ {e.NextPaths()}
 paths:=e.NextPaths()
 if paths.Upper!=0 {t.Fatalf("upper 1-foot drawbar still sounding: %v",paths.Upper)}
 if paths.Lower==0 {t.Fatal("lower manual unexpectedly silent")}
 if err:=e.SetHarmonicPercussion(0,3,0,0.9);err!=nil{t.Fatal(err)}
 for i:=0;i<150;i++ {e.NextPaths()}
 paths=e.NextPaths()
 if math.Abs(paths.Upper-paths.Lower)>1e-12 {t.Fatalf("upper drawbar not restored: %+v",paths)}
}
