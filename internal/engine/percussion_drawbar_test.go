package engine

import (
 "testing"
 "github.com/cpu-model/hammond/internal/registration"
)

func TestUpperNinthDrawbarCancellation(t *testing.T) {
 var route registration.Routing
 route[0][8]=1
 e,err:=NewTimed(48000,route,route,registration.ContactTiming{},registration.ContactTiming{})
 if err!=nil {t.Fatal(err)}
 upper,_:=e.Manual(0)
 lower,_:=e.Manual(1)
 if err:=upper.SetDrawbar(8,8);err!=nil {t.Fatal(err)}
 if err:=lower.SetDrawbar(8,8);err!=nil {t.Fatal(err)}
 if err:=upper.Key(0,true);err!=nil {t.Fatal(err)}
 if err:=lower.Key(0,true);err!=nil {t.Fatal(err)}
 if err:=e.SetHarmonicPercussion(0,3,0.5,0.9);err!=nil {t.Fatal(err)}
 for i:=0;i<150;i++ {e.NextPaths()}
 if !upperPercussionCut(upper) {t.Fatal("upper drawbar not cancelled")}
 if upperPercussionCut(lower) {t.Fatal("lower drawbar cancelled")}
 if err:=e.SetHarmonicPercussion(0,3,0,0.9);err!=nil {t.Fatal(err)}
 for i:=0;i<150;i++ {e.NextPaths()}
 if upperPercussionCut(upper) {t.Fatal("drawbar did not return")}
}

// Verify cancellation through rendered sound rather than private state.
func upperPercussionCut(m *registration.TimedManual) bool {
 frame:=make([]float64,91)
 frame[0]=1
 return m.Mix(frame)==0
}
