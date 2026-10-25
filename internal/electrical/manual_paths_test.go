package electrical

import "testing"

func TestManualPathsPreserveSeparation(t *testing.T) {
 p:=ManualPaths{Upper:0.25,Lower:-0.75}
 if p.Upper!=0.25 || p.Lower!=-0.75 { t.Fatal("manual path changed") }
 if p.Sum()!=-0.5 { t.Fatal("diagnostic sum mismatch") }
}
