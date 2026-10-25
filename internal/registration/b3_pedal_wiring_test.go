package registration

import "testing"

func TestCompileB3PedalWiringRequiresSourceAndBusbars(t *testing.T) {
 e := PedalWiringEvidence{
  ContactOrder: PedalContactOrder{0,1,2,3,4,5,6,7},
  Busbars: PedalBusRouting{1,1,2,2,3,3,4,4},
 }
 if _,_,err:=CompileB3PedalWiring(e); err==nil { t.Fatal("accepted missing source") }
 e.Source=" \t\n "
 if _,_,err:=CompileB3PedalWiring(e); err==nil { t.Fatal("accepted whitespace-only source") }
 e.Source="synthetic test fixture, not historical B3 wiring"
 if _,_,err:=CompileB3PedalWiring(e);err==nil {t.Fatal("accepted unverified physical wiring")}
 e.Verified=true
 routing,bus,err:=CompileB3PedalWiring(e)
 if err!=nil { t.Fatal(err) }
 if bus!=e.Busbars { t.Fatal("busbar assignment changed") }
 if routing[7][1]!=0 || routing[11][1]!=0 { t.Fatal("disconnected contacts filled") }
 e.Busbars[4]=0
 if _,_,err:=CompileB3PedalWiring(e); err==nil { t.Fatal("accepted missing busbar") }
 e.Busbars[4]=5
 if _,_,err:=CompileB3PedalWiring(e); err==nil { t.Fatal("accepted invalid busbar") }
 e.Busbars[4]=3
 e.ContactOrder[7]=6
 if _,_,err:=CompileB3PedalWiring(e); err==nil { t.Fatal("accepted duplicate harmonic") }
}
