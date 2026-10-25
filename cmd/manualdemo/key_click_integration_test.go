package main

import (
 "bytes"
 "os"
 "os/exec"
 "path/filepath"
 "testing"
)

func TestKeyClickChangesRenderedWAV(t *testing.T) {
 dir:=t.TempDir()
 render:=func(name,click string)[]byte{
  t.Helper()
  output:=filepath.Join(dir,name+".wav")
  cmd:=exec.Command("go","run",".","-out",output,"-seconds","1","-release","0.5","-drawbars","800000000","-key-click",click)
  if result,err:=cmd.CombinedOutput();err!=nil {t.Fatalf("render failed: %v: %s",err,result)}
  data,err:=os.ReadFile(output);if err!=nil{t.Fatal(err)}
  if len(data)<44||string(data[:4])!="RIFF" {t.Fatal("invalid WAV")}
  return data
 }
 off:=render("off","0")
 on:=render("on","0.5")
 if bytes.Equal(off,on){t.Fatal("key click did not change rendered samples")}
 if len(off)!=len(on){t.Fatal("click changed recording length")}
}
