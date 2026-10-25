package electrical

import "testing"

func TestTransformerPairKeepsManualsSeparate(t *testing.T) {
 var upper,lower DrawbarTapSelection
 upper[0]=8
 upper[1]=1
 lower[8]=5
 pair:=ManualTransformerPair{
  Upper:ManualTransformerConnection{ID:TransformerT2,Selection:upper},
  Lower:ManualTransformerConnection{ID:TransformerT1,Selection:lower},
 }
 var uv,lv [9]float64
 uv[0],uv[1]=2,3
 lv[8]=-4
 u,l,err:=pair.Inputs(uv,lv)
 if err!=nil { t.Fatal(err) }
 if u.Transformer!=TransformerT2||l.Transformer!=TransformerT1 { t.Fatal("incorrect transformer identity") }
 if !u.Connected[0]||u.Tap[0]!=8||u.Voltage[0]!=2||!u.Connected[1]||u.Tap[1]!=1 { t.Fatalf("upper inputs: %+v",u) }
 if u.Connected[8]||l.Connected[0]||!l.Connected[8]||l.Tap[8]!=5||l.Voltage[8]!=-4 { t.Fatalf("manual leakage: %+v %+v",u,l) }
}

func TestTransformerPairRejectsSwappedIdentity(t *testing.T) {
 pair:=ManualTransformerPair{
  Upper:ManualTransformerConnection{ID:TransformerT1},
  Lower:ManualTransformerConnection{ID:TransformerT2},
 }
 if err:=pair.Validate();err==nil { t.Fatal("accepted swapped T1/T2 identities") }
 pair.Upper.ID=TransformerT2
 pair.Lower.ID=TransformerT1
 pair.Upper.Selection[0]=9
 if err:=pair.Validate();err==nil { t.Fatal("accepted invalid tap position") }
}

func TestTransformerPairNoAllocations(t *testing.T) {
 pair:=ManualTransformerPair{
  Upper:ManualTransformerConnection{ID:TransformerT2},
  Lower:ManualTransformerConnection{ID:TransformerT1},
 }
 if n:=testing.AllocsPerRun(100,func(){ _,_,_=pair.Inputs([9]float64{},[9]float64{}) });n!=0 {
  t.Fatalf("allocations: %v",n)
 }
}
