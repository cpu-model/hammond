package registration

import "testing"

func TestPedalContactsEightPerKeyAndRelease(t *testing.T) {
 var routing PedalRouting
 for i:=0;i<PedalContacts;i++ { routing[0][i]=uint8(i+1) }
 pedal,err:=NewPedalContactsState(routing)
 if err!=nil { t.Fatal(err) }
 if err:=pedal.Key(0,true);err!=nil { t.Fatal(err) }
 count:=0
 pedal.ForEachClosedContact(func(key,contact,wheel int){
  if key!=0||contact!=count||wheel!=count+1 { t.Fatalf("unexpected contact %d/%d/%d",key,contact,wheel) }
  count++
 })
 if count!=8 { t.Fatalf("expected eight contacts, got %d",count) }
 if err:=pedal.Key(0,false);err!=nil { t.Fatal(err) }
 pedal.ForEachClosedContact(func(_,_,_ int){ t.Fatal("released key remained connected") })
}

func TestPedalRoutingAndKeyBounds(t *testing.T) {
 var routing PedalRouting
 routing[24][7]=92
 if _,err:=NewPedalContactsState(routing);err==nil { t.Fatal("accepted wheel 92") }
 routing[24][7]=91
 p,err:=NewPedalContactsState(routing)
 if err!=nil { t.Fatal(err) }
 for _,key:=range []int{-1,25} {
  if err:=p.Key(key,true);err==nil { t.Fatalf("accepted pedal key %d",key) }
 }
 if err:=p.Key(24,true);err!=nil { t.Fatal(err) }
 count:=0
 p.ForEachClosedContact(func(key,contact,wheel int) {
  if key!=24||contact!=7||wheel!=91 { t.Fatal("incorrect last pedal contact") }
  count++
 })
 if count!=1 { t.Fatalf("expected one routed contact, got %d",count) }
}

func TestPedalContactsTraversalNoAllocations(t *testing.T) {
 var routing PedalRouting
 routing[0][0]=1
 p,_:=NewPedalContactsState(routing)
 _=p.Key(0,true)
 if n:=testing.AllocsPerRun(100,func(){
  p.ForEachClosedContact(func(_,_,_ int){})
 });n!=0 { t.Fatalf("allocations: %v",n) }
}
