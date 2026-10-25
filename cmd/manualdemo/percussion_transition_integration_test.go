package main

import (
 "bytes"
 "os"
 "os/exec"
 "path/filepath"
 "testing"
)

func TestPercussionToggleTransitionChangesWAV(t *testing.T) {
 dir:=t.TempDir()
 render:=func(ms string)[]byte{
  t.Helper()
  path:=filepath.Join(dir,"transition"+ms+".wav")
  cmd:=exec.Command("go","run",".","-out",path,"-seconds","1","-release","0","-drawbars","000000008","-percussion-preset","fast-normal","-percussion-toggle","0.5","-ninth-transition-ms",ms)
  if output,err:=cmd.CombinedOutput();err!=nil{t.Fatalf("render failed: %v: %s",err,output)}
  data,err:=os.ReadFile(path);if err!=nil{t.Fatal(err)}
  return data
 }
 immediate:=render("0")
 smooth:=render("10")
 if len(immediate)!=len(smooth){t.Fatal("transition altered WAV duration")}
 if bytes.Equal(immediate,smooth){t.Fatal("transition setting did not affect WAV")}
}
