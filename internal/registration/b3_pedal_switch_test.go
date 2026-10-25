package registration

import "testing"

func TestB3PedalSwitchRoutesThroughBusbars(t *testing.T) {
 // Deliberately synthetic wiring fixture, not a claim about B3 factory order.
 e:=PedalWiringEvidence{
  ContactOrder:PedalContactOrder{2,0,1,3,4,5,6,7},
  Busbars:PedalBusRouting{1,1,2,2,3,3,4,4},
  Source:"synthetic integration fixture",
 }
 p,err:=NewB3PedalSwitch(e)
 if err!=nil { t.Fatal(err) }
 if err:=p.Key(7,true);err!=nil { t.Fatal(err) }
 count:=0
 err=p.ForEachClosedBusContact(func(key,contact,busbar,wheel int){
  count++
  if key!=7 { t.Fatalf("unexpected key %d",key) }
  if contact==2 { t.Fatal("disconnected third harmonic emitted") }
  if busbar<0||busbar>=PedalBusbars { t.Fatalf("invalid busbar %d",busbar) }
  if wheel==0 { t.Fatal("disconnected frequency emitted") }
 })
 if err!=nil { t.Fatal(err) }
 if count!=7 { t.Fatalf("got %d connected contacts, want 7",count) }
 if err:=p.Key(7,false);err!=nil { t.Fatal(err) }
 count=0
 if err:=p.ForEachClosedBusContact(func(_,_,_,_ int){count++});err!=nil { t.Fatal(err) }
 if count!=0 { t.Fatalf("released pedal still has %d contacts",count) }
 if _,err:=NewB3PedalSwitch(PedalWiringEvidence{});err==nil { t.Fatal("accepted unknown wiring") }
}
