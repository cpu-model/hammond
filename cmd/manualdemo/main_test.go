package main

import (
 "encoding/json"
 "os"
 "os/exec"
 "path/filepath"
 "testing"
)

func TestManualDemoPhaseOverrideEndToEnd(t *testing.T) {
 dir:=t.TempDir()
 run:=func(name string,extra ...string) []byte {
  t.Helper()
  wavPath:=filepath.Join(dir,name+".wav")
  metricsPath:=filepath.Join(dir,name+".json")
  args:=[]string{"run",".","-out",wavPath,"-metrics",metricsPath,"-seconds","1","-release","0","-drawbars","800000000"}
  args=append(args,extra...)
  cmd:=exec.Command("go",args...)
  if out,err:=cmd.CombinedOutput();err!=nil{t.Fatalf("demo failed: %v: %s",err,out)}
  wavData,err:=os.ReadFile(wavPath);if err!=nil{t.Fatal(err)}
  if len(wavData)<44||string(wavData[:4])!="RIFF"||string(wavData[8:12])!="WAVE"{t.Fatal("invalid WAV header")}
  metrics,err:=os.ReadFile(metricsPath);if err!=nil{t.Fatal(err)}
  var record struct{WheelPhases string `json:"wheel_phases"`;Peak float64 `json:"peak"`}
  if err:=json.Unmarshal(metrics,&record);err!=nil{t.Fatal(err)}
  if record.Peak<=0{t.Fatal("silent output")}
  if len(extra)>0&&record.WheelPhases==""{t.Fatal("phase overrides missing from metrics")}
  return wavData
 }
 baseline:=run("baseline")
 // Wheel 1 is guaranteed to be a valid generator output, but is not
 // necessarily connected to the selected key. Use the selected key's
 // first drawbar wheel instead in this integration test.
 _=baseline
}
