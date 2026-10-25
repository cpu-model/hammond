package registration

import "testing"

func TestPercussionSingleTrigger(t *testing.T) {
 m,err:=NewTimedManual(Routing{},ContactTiming{})
 if err!=nil{t.Fatal(err)}
 if m.TakePercussionTrigger(){t.Fatal("initial trigger")}
 if err:=m.Key(0,true);err!=nil{t.Fatal(err)}
 if !m.TakePercussionTrigger(){t.Fatal("first press did not trigger")}
 if m.TakePercussionTrigger(){t.Fatal("trigger not consumed")}
 if err:=m.Key(1,true);err!=nil{t.Fatal(err)}
 if m.TakePercussionTrigger(){t.Fatal("legato press retriggered")}
 if err:=m.Key(0,false);err!=nil{t.Fatal(err)}
 if err:=m.Key(0,true);err!=nil{t.Fatal(err)}
 if m.TakePercussionTrigger(){t.Fatal("partial release rearmed percussion")}
 if err:=m.Key(0,false);err!=nil{t.Fatal(err)}
 if err:=m.Key(1,false);err!=nil{t.Fatal(err)}
 if err:=m.Key(1,true);err!=nil{t.Fatal(err)}
 if !m.TakePercussionTrigger(){t.Fatal("all-key release did not rearm")}
}

func TestPercussionTriggerIgnoresRepeatedKeyRequest(t *testing.T) {
 m,err:=NewTimedManual(Routing{},ContactTiming{})
 if err!=nil{t.Fatal(err)}
 _=m.Key(0,true)
 if !m.TakePercussionTrigger(){t.Fatal("missing initial trigger")}
 _=m.Key(0,true)
 if m.TakePercussionTrigger(){t.Fatal("repeated key down retriggered")}
}
