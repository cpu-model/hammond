package main

import (
 "bytes"
 "fmt"
 "os"
 "os/exec"
 "path/filepath"
 "testing"

 "github.com/cpu-model/hammond/internal/registration"
)

func TestActiveWheelPhaseChangesRenderedWAV(t *testing.T) {
 dir:=t.TempDir()
 render:=func(name string,phase string)[]byte{
  t.Helper()
  output:=filepath.Join(dir,name+".wav")
  args:=[]string{"run",".","-out",output,"-seconds","1","-release","0","-drawbars","800000000"}
  if phase!="" {args=append(args,"-wheel-phases",phase)}
  cmd:=exec.Command("go",args...)
  if result,err:=cmd.CombinedOutput();err!=nil {t.Fatalf("render failed: %v: %s",err,result)}
  data,err:=os.ReadFile(output);if err!=nil {t.Fatal(err)}
  return data
 }
 wheel:=registration.B3ManualRouting()[36][0]
 if wheel==0 {t.Fatal("unconnected first drawbar")}
 baseline:=render("baseline","")
 altered:=render("altered",fmt.Sprintf("%d:0.25",wheel))
 if bytes.Equal(baseline,altered) {t.Fatal("phase change did not affect WAV samples")}
}
