package registration

import "testing"

func TestPedalBusbarMappingAndContactPairs(t *testing.T) {
 var routing PedalRouting
 for i:=0;i<8;i++ { routing[0][i]=uint8(i+1) }
 p,err:=NewPedalContactsState(routing)
 if err!=nil { t.Fatal(err) }
 _=p.Key(0,true)
 mapping:=PedalBusRouting{1,1,2,2,3,3,4,4} // synthetic fixture, not B3 wiring
 var counts [4]int
 if err:=p.ForEachClosedBusContact(mapping,func(key,contact,busbar,wheel int){
  if key!=0||wheel!=contact+1 { t.Fatal("incorrect contact") }
  counts[busbar]++
 });err!=nil { t.Fatal(err) }
 for i,n:=range counts { if n!=2 { t.Fatalf("bus %d got %d contacts",i,n) } }
 _=p.Key(0,false)
 calls:=0
 if err:=p.ForEachClosedBusContact(mapping,func(_,_,_,_ int){calls++});err!=nil||calls!=0 {
  t.Fatalf("released pedal still connected: %d %v",calls,err)
 }
}

func TestPedalBusbarMappingValidation(t *testing.T) {
 p,_:=NewPedalContactsState(PedalRouting{})
 invalid:=PedalBusRouting{5}
 if err:=p.ForEachClosedBusContact(invalid,func(_,_,_,_ int){});err==nil { t.Fatal("accepted busbar 5") }
}

func TestPedalBusbarTraversalNoAllocations(t *testing.T) {
 var routing PedalRouting
 routing[0][0]=1
 p,_:=NewPedalContactsState(routing)
 _=p.Key(0,true)
 mapping:=PedalBusRouting{1}
 if n:=testing.AllocsPerRun(100,func(){
  _=p.ForEachClosedBusContact(mapping,func(_,_,_,_ int){})
 });n!=0 { t.Fatalf("allocations: %v",n) }
}
