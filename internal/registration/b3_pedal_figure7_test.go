package registration

import "testing"

func TestB3PedalContactOrderFigure7(t *testing.T) {
 order:=B3PedalContactOrderFigure7()
 if err:=order.Validate();err!=nil { t.Fatal(err) }
 want:=PedalContactOrder{6,4,2,0,7,5,3,1}
 if order!=want { t.Fatalf("physical order = %v, want %v",order,want) }
 harmonics:=PedalRouting{}
 harmonics[0]=[PedalContacts]uint8{1,2,3,4,5,6,7,8}
 physical,err:=PedalRoutingFromHarmonics(harmonics,order)
 if err!=nil { t.Fatal(err) }
 for i,column:=range want {
  if physical[0][i]!=harmonics[0][column] {
   t.Fatalf("contact %d = %d, expected harmonic column %d",i,physical[0][i],column)
  }
 }
}
