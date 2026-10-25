package registration

import "testing"

func TestB3PedalFactoryWiring(t *testing.T) {
 e:=B3PedalWiringHandyGuide()
 if !e.Verified {t.Fatal("factory wiring must carry source attestation")}
 want:=PedalBusRouting{4,3,2,1,4,3,2,1}
 if e.Busbars!=want {t.Fatalf("busbars %v, want %v",e.Busbars,want)}
 routing,busbars,err:=CompileB3PedalWiring(e)
 if err!=nil {t.Fatal(err)}
 if busbars!=want {t.Fatal("busbar mapping changed")}
 harmonics:=B3PedalHarmonics()
 for key:=0;key<PedalKeys;key++ {
  for physical,column:=range e.ContactOrder {
   if routing[key][physical]!=harmonics[key][column] {
    t.Fatalf("pedal %d contact %d: got %d want %d",key+1,physical,routing[key][physical],harmonics[key][column])
   }
  }
 }
 // Pedal 8 has no 3rd harmonic contact.
 if routing[7][7]!=0 {t.Fatal("pedal 8 third harmonic must be disconnected")}
}
