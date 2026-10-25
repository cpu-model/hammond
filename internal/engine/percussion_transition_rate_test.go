package engine

import (
 "math"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
)

func TestDefaultNinthDrawbarFadeDurationAcrossSampleRates(t *testing.T) {
 for _,rate:=range []float64{44100,48000,96000} {
  t.Run(fmtRate(rate),func(t *testing.T) {
   var route registration.Routing
   route[0][8]=1
   e,err:=NewTimed(rate,route,route,registration.ContactTiming{},registration.ContactTiming{})
   if err!=nil{t.Fatal(err)}
   upper,_:=e.Manual(0)
   if err:=upper.SetDrawbar(8,8);err!=nil{t.Fatal(err)}
   if err:=upper.Key(0,true);err!=nil{t.Fatal(err)}
   frame:=make([]float64,91)
   frame[0]=1
   upper.Mix(frame)
   if err:=e.SetHarmonicPercussion(0,3,0.5,0.9);err!=nil{t.Fatal(err)}
   count:=0
   for count<int(rate/10) {
    count++
    if math.Abs(upper.Mix(frame))<1e-12 {break}
   }
   expected:=int(math.Round(rate*0.003))
   if count!=expected {t.Fatalf("fade lasted %d samples, want %d",count,expected)}
   if math.Abs(float64(count)/rate-0.003)>1/rate {t.Fatalf("fade duration %g seconds",float64(count)/rate)}
  })
 }
}

func fmtRate(rate float64)string {
 if rate==44100{return "44100"}
 if rate==48000{return "48000"}
 return "96000"
}
