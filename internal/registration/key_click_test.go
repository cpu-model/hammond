package registration

import (
 "math"
 "testing"
)

func TestKeyClickOnContactMakeAndBreak(t *testing.T) {
 var route Routing
 route[0][0]=1
 var timing ContactTiming
 timing.Make[0]=2
 m,err:=NewTimedManual(route,timing)
 if err!=nil{t.Fatal(err)}
 if err:=m.SetDrawbar(0,8);err!=nil{t.Fatal(err)}
 if err:=m.SetKeyClick(0.9);err!=nil{t.Fatal(err)}
 if err:=m.Key(0,true);err!=nil{t.Fatal(err)}
 frame:=make([]float64,91)
 for i:=0;i<2;i++ {
  if got:=m.Mix(frame);got!=0 {t.Fatalf("click before contact closure at sample %d: %v",i,got)}
 }
 if got:=m.Mix(frame);got<=0 {t.Fatalf("missing make click: %v",got)}
 if got:=m.Mix(frame);got>=0 {t.Fatalf("missing bipolar click decay: %v",got)}
 if err:=m.Key(0,false);err!=nil{t.Fatal(err)}
 if got:=m.Mix(frame);got>=0 {t.Fatalf("missing break click: %v",got)}
}

func TestKeyClickOffAndValidation(t *testing.T) {
 var route Routing
 route[0][0]=1
 m,err:=NewTimedManual(route,ContactTiming{})
 if err!=nil{t.Fatal(err)}
 if err:=m.SetDrawbar(0,8);err!=nil{t.Fatal(err)}
 for _,v:=range []float64{-1,1.01,math.NaN(),math.Inf(1)} {
  if err:=m.SetKeyClick(v);err==nil {t.Fatalf("accepted invalid click level %v",v)}
 }
 if err:=m.Key(0,true);err!=nil{t.Fatal(err)}
 frame:=make([]float64,91)
 if got:=m.Mix(frame);got!=0 {t.Fatalf("default click should be silent: %v",got)}
 if err:=m.SetKeyClick(1);err!=nil{t.Fatal(err)}
 if err:=m.Key(0,false);err!=nil{t.Fatal(err)}
 if got:=m.Mix(frame);got==0 {t.Fatal("expected transient")}
 if err:=m.SetKeyClick(0);err!=nil{t.Fatal(err)}
 if got:=m.Mix(frame);got!=0 {t.Fatalf("disabling click left residual: %v",got)}
}
