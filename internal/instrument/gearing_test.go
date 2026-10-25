package instrument

import (
 "math"
 "testing"
)

func TestMechanicalFrequencies(t *testing.T) {
 cases := []struct{ number int; want float64 }{
  {1, 40*85.0/104}, {10, 55}, {13, 80*85.0/104},
  {46, 440}, {82, 3520}, {84, 20*128*108.0/70},
 }
 for _, tc := range cases {
  got, err := VerifiedMechanicalFrequency(tc.number)
  if err != nil || math.Abs(got-tc.want)>1e-9 { t.Fatalf("%d: got %.12f want %.12f: %v",tc.number,got,tc.want,err) }
 }
 for _, n := range []int{0,85,91} {
  if _,err:=VerifiedMechanicalFrequency(n); err==nil { t.Fatalf("accepted unverified frequency %d",n) }
 }
}

func TestOctaveGearRelation(t *testing.T) {
 for n:=1;n<=72;n++ {
  a,_:=VerifiedMechanicalFrequency(n)
  b,_:=VerifiedMechanicalFrequency(n+12)
  if math.Abs(b-2*a)>1e-9 { t.Fatalf("octave mismatch at %d",n) }
 }
}
