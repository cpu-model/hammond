package registration

import "testing"

func TestAdvanceContactsMatchesMixTiming(t *testing.T) {
 var route Routing
 route[0][0]=1
 route[0][1]=2
 var timing ContactTiming
 timing.Make[0]=2
 timing.Break[1]=1
 a,err:=NewTimedManual(route,timing)
 if err!=nil { t.Fatal(err) }
 b,err:=NewTimedManual(route,timing)
 if err!=nil { t.Fatal(err) }
 frame:=make([]float64,91)
 frame[0],frame[1]=1,1
 _=a.SetDrawbar(0,8)
 _=a.SetDrawbar(1,8)
 _=b.SetDrawbar(0,8)
 _=b.SetDrawbar(1,8)
 for sample:=0;sample<8;sample++ {
  if sample==0 { _=a.Key(0,true);_=b.Key(0,true) }
  if sample==4 { _=a.Key(0,false);_=b.Key(0,false) }
  a.AdvanceContacts()
  var activeA,activeB [9][92]int
  a.ForEachActiveContact(func(bus,wheel int){activeA[bus][wheel]++})
  b.Mix(frame)
  b.ForEachActiveContact(func(bus,wheel int){activeB[bus][wheel]++})
  if activeA!=activeB { t.Fatalf("contact mismatch at sample %d: %v vs %v",sample,activeA,activeB) }
 }
}

func TestAdvanceContactsNoAllocations(t *testing.T) {
 m,err:=NewTimedManual(Routing{},ContactTiming{})
 if err!=nil { t.Fatal(err) }
 if n:=testing.AllocsPerRun(100,func(){m.AdvanceContacts()});n!=0 {
  t.Fatalf("allocations: %v",n)
 }
}
