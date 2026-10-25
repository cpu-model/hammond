package registration

import "testing"

func TestB3PedalHarmonics(t *testing.T) {
 table := B3PedalHarmonics()
 if err := table.Validate(); err != nil { t.Fatal(err) }
 for i,row := range table {
  if row[0] != uint8(i+1) { t.Fatalf("pedal %d fundamental = %d",i+1,row[0]) }
  if i >= 7 && i <= 11 {
   if row[1] != 0 { t.Fatalf("pedal %d third harmonic should be disconnected",i+1) }
  } else if row[1] == 0 {
   t.Fatalf("pedal %d missing third harmonic",i+1)
  }
 }
 if table[0] != ([8]uint8{1,13,13,25,32,37,41,44}) { t.Fatal("pedal 1 differs from Table III") }
 if table[12] != ([8]uint8{13,32,25,37,44,49,53,56}) { t.Fatal("pedal 13 differs from Table III") }
 if table[24] != ([8]uint8{25,44,37,49,56,61,65,68}) { t.Fatal("pedal 25 differs from Table III") }
}
