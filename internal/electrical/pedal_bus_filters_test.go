package electrical

import (
 "math"
 "testing"
)

func TestExperimentalPedalBusFiltersIndependentAndAtomic(t *testing.T) {
 var cfg [4]PedalRLConfig
 for i:=range cfg { cfg[i]=pedalRLFixture() }
 f,err:=NewExperimentalPedalBusFilters(48000,cfg)
 if err!=nil { t.Fatal(err) }
 in:=[4]float64{1,0,0,-1}
 a,err:=f.Next(in)
 if err!=nil { t.Fatal(err) }
 if !(a[0]>0&&a[1]==0&&a[2]==0&&a[3]==-a[0]) { t.Fatalf("unexpected channel response %v",a) }
 if _,err:=f.Next([4]float64{math.NaN(),0,0,0});err==nil { t.Fatal("accepted nonfinite frame") }
 b,err:=f.Next(in)
 if err!=nil { t.Fatal(err) }
 f.Reset()
 _,_ = f.Next(in)
 c,err:=f.Next(in)
 if err!=nil { t.Fatal(err) }
 if b!=c { t.Fatalf("invalid frame advanced state: %v vs %v",b,c) }
 f.Reset()
 zero,err:=f.Next([4]float64{})
 if err!=nil||zero!=([4]float64{}) { t.Fatalf("reset failed %v %v",zero,err) }
}

func TestExperimentalPedalBusFiltersConfigValidation(t *testing.T) {
 var cfg [4]PedalRLConfig
 for i:=range cfg { cfg[i]=pedalRLFixture() }
 cfg[3].InductanceHenry.Value=0
 if _,err:=NewExperimentalPedalBusFilters(48000,cfg);err==nil { t.Fatal("accepted invalid fourth channel") }
}

func TestExperimentalPedalBusFiltersNoAllocations(t *testing.T) {
 var cfg [4]PedalRLConfig
 for i:=range cfg { cfg[i]=pedalRLFixture() }
 f,err:=NewExperimentalPedalBusFilters(48000,cfg)
 if err!=nil { t.Fatal(err) }
 if n:=testing.AllocsPerRun(100,func(){_,_ =f.Next([4]float64{1,2,3,4})});n!=0 { t.Fatalf("allocations per frame: %v",n) }
}
