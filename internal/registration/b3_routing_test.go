package registration

import "testing"

func TestB3ManualRouting(t *testing.T) {
 route:=B3ManualRouting()
 if _,err:=NewManual(route);err!=nil { t.Fatal(err) }
 for k,contacts:=range route {
  for d,w:=range contacts {
   if w<13 || w>91 { t.Fatalf("key %d contact %d: wheel %d",k+1,d+1,w) }
  }
 }
 cases:=[]struct{ key,drawbar int; wheel uint8 }{
  {1,1,13}, {1,2,20}, {1,3,13},
  {13,1,13}, {13,2,32}, {13,9,61},
  {43,9,91}, {44,9,80}, {61,1,61}, {61,9,85},
 }
 for _,tc:=range cases {
  got:=route[tc.key-1][tc.drawbar-1]
  if got!=tc.wheel { t.Fatalf("key %d drawbar %d: %d != %d",tc.key,tc.drawbar,got,tc.wheel) }
 }
}

func TestB3Foldback(t *testing.T) {
 route:=B3ManualRouting()
 if route[42][8]!=91 || route[43][8]!=80 {
  t.Fatal("top harmonic does not fold back at key 44")
 }
 if route[55][5]!=80 || route[60][5]!=85 {
  t.Fatal("upper-register harmonic foldback incorrect")
 }
}
