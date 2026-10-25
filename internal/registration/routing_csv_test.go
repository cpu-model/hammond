package registration

import (
 "fmt"
 "strings"
 "testing"
)

func routingFixture() string {
 var b strings.Builder
 b.WriteString("key,drawbar,wheel\n")
 for k:=1;k<=Keys;k++ {
  for d:=1;d<=Harmonics;d++ {
   fmt.Fprintf(&b,"%d,%d,%d\n",k,d,((k+d-2)%91)+1)
  }
 }
 return b.String()
}

func TestParseRoutingCSV(t *testing.T) {
 route,err:=ParseRoutingCSV(strings.NewReader(routingFixture()))
 if err!=nil { t.Fatal(err) }
 if route[0][0]!=1 || route[60][8]!=69 { t.Fatalf("unexpected test routing: %d %d",route[0][0],route[60][8]) }
}

func TestParseRoutingCSVRejectsMissing(t *testing.T) {
 fixture:=routingFixture()
 lines:=strings.Split(strings.TrimSuffix(fixture,"\n"),"\n")
 if _,err:=ParseRoutingCSV(strings.NewReader(strings.Join(lines[:len(lines)-1],"\n")));err==nil { t.Fatal("accepted incomplete wiring") }
}

func TestParseRoutingCSVRejectsDuplicate(t *testing.T) {
 fixture:=routingFixture()+"1,1,46\n"
 if _,err:=ParseRoutingCSV(strings.NewReader(fixture));err==nil { t.Fatal("accepted duplicate contact") }
}

func TestParseRoutingCSVRejectsBadWheel(t *testing.T) {
 fixture:=strings.Replace(routingFixture(),"1,1,1\n","1,1,92\n",1)
 if _,err:=ParseRoutingCSV(strings.NewReader(fixture));err==nil { t.Fatal("accepted invalid wheel") }
}
