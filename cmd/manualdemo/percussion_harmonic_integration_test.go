package main

import (
 "bytes"
 "os"
 "os/exec"
 "path/filepath"
 "testing"
)

func TestHarmonicPercussionSelectionChangesAudio(t *testing.T) {
 dir:=t.TempDir()
 render:=func(harmonic string)[]byte{
  t.Helper()
  output:=filepath.Join(dir,"harmonic"+harmonic+".wav")
  cmd:=exec.Command("go","run",".","-out",output,"-seconds","1","-release","0","-drawbars","000000000","-percussion-level","0.5","-percussion-harmonic",harmonic)
  if out,err:=cmd.CombinedOutput();err!=nil{t.Fatalf("render %s: %v: %s",harmonic,err,out)}
  data,err:=os.ReadFile(output);if err!=nil{t.Fatal(err)}
  return data
 }
 second:=render("2")
 third:=render("3")
 if bytes.Equal(second,third){t.Fatal("second and third harmonic percussion rendered identically")}
 if len(second)!=len(third){t.Fatal("different WAV durations")}
}
