package registration

import "testing"

func TestActiveContactVisitDoesNotAdvanceClock(t *testing.T) {
 var route Routing
 route[0][0]=1
 route[0][1]=2
 route[1][0]=3
 var timing ContactTiming
 timing.Make[0]=1
 m,err:=NewTimedManual(route,timing)
 if err!=nil { t.Fatal(err) }
 if err:=m.Key(0,true);err!=nil { t.Fatal(err) }
 if err:=m.Key(1,true);err!=nil { t.Fatal(err) }
 count:=0
 m.ForEachActiveContact(func(bus,wheel int){ count++ })
 if count!=0 { t.Fatalf("contact active before rendering: %d",count) }
 frame:=make([]float64,91)
 m.Mix(frame)
 var seen [9][92]int
 m.ForEachActiveContact(func(bus,wheel int){ seen[bus][wheel]++ })
 if seen[1][2]!=1||seen[0][1]!=0||seen[0][3]!=0 {
  t.Fatalf("incorrect initial contacts: %v",seen)
 }
 m.ForEachActiveContact(func(bus,wheel int){ seen[bus][wheel]++ })
 if seen[1][2]!=2 { t.Fatal("visit mutated state") }
 m.Mix(frame)
 m.ForEachActiveContact(func(bus,wheel int){ seen[bus][wheel]++ })
 if seen[0][1]!=1||seen[0][3]!=1 { t.Fatalf("delayed contacts did not close: %v",seen) }
}

func TestActiveContactVisitSkipsDisconnectedWheel(t *testing.T) {
 var route Routing
 route[0][0]=0
 m,err:=NewTimedManual(route,ContactTiming{})
 if err!=nil { t.Fatal(err) }
 _=m.Key(0,true)
 m.Mix(make([]float64,91))
 m.ForEachActiveContact(func(bus,wheel int){ t.Fatal("disconnected contact visited") })
}
