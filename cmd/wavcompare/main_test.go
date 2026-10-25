package main

import (
 "encoding/binary"
 "math"
 "os"
 "path/filepath"
 "testing"
)

func TestMagnitudeKnownSine(t *testing.T) {
 const rate=48000
 const freq=1000
 x:=make([]float64,4800)
 for i:=range x{x[i]=0.25*math.Sin(2*math.Pi*freq*float64(i)/rate)}
 if got:=magnitude(x,freq,rate);math.Abs(got-0.25)>1e-6 {t.Fatalf("fundamental amplitude %.9f, want 0.25",got)}
 if got:=magnitude(x,2*freq,rate);got>1e-9 {t.Fatalf("unexpected second harmonic %.9f",got)}
 if got:=rms(x);math.Abs(got-0.25/math.Sqrt2)>1e-9 {t.Fatalf("RMS %.9f",got)}
}

func TestReadWAVMono16(t *testing.T){
 path:=filepath.Join(t.TempDir(),"test.wav")
 data:=make([]byte,48)
 copy(data[:4],"RIFF")
 binary.LittleEndian.PutUint32(data[4:],40)
 copy(data[8:12],"WAVE")
 copy(data[12:16],"fmt ")
 binary.LittleEndian.PutUint32(data[16:],16)
 binary.LittleEndian.PutUint16(data[20:],1)
 binary.LittleEndian.PutUint16(data[22:],1)
 binary.LittleEndian.PutUint32(data[24:],48000)
 binary.LittleEndian.PutUint32(data[28:],96000)
 binary.LittleEndian.PutUint16(data[32:],2)
 binary.LittleEndian.PutUint16(data[34:],16)
 copy(data[36:40],"data")
 binary.LittleEndian.PutUint32(data[40:],4)
 binary.LittleEndian.PutUint16(data[44:],16384)
 binary.LittleEndian.PutUint16(data[46:],uint16(0xC000))
 if err:=os.WriteFile(path,data,0600);err!=nil{t.Fatal(err)}
 a,err:=readWAV(path);if err!=nil{t.Fatal(err)}
 if a.rate!=48000||len(a.samples)!=2||a.samples[0]!=0.5||a.samples[1]!=-0.5 {t.Fatalf("decoded %+v",a)}
}

func TestReadWAVRejectsMalformed(t *testing.T){
 path:=filepath.Join(t.TempDir(),"bad.wav")
 if err:=os.WriteFile(path,[]byte("RIFF"),0600);err!=nil{t.Fatal(err)}
 if _,err:=readWAV(path);err==nil{t.Fatal("accepted truncated WAV")}
}

func TestMagnitudeNonBinCenteredTone(t *testing.T) {
 const rate=48000
 const frequency=997.3
 x:=make([]float64,4800)
 for i:=range x{x[i]=0.4*math.Sin(2*math.Pi*frequency*float64(i)/rate)}
 if got:=magnitude(x,frequency,rate);math.Abs(got-0.4)>0.001 {
  t.Fatalf("off-bin fundamental amplitude %.6f, want approximately 0.4",got)
 }
 if got:=magnitude(x,2*frequency,rate);got>0.001 {
  t.Fatalf("off-bin second harmonic leakage %.6f",got)
 }
}

func TestMagnitudeShortWindow(t *testing.T) {
 if got:=magnitude(nil,1000,48000);got!=0 {t.Fatalf("empty magnitude %v",got)}
 if got:=magnitude([]float64{0.5},1000,48000);got!=0 {t.Fatalf("single-sample magnitude %v",got)}
}

func TestMagnitudeRejectsDistantFrequencyLeakage(t *testing.T) {
 const rate=48000
 x:=make([]float64,4800)
 for i:=range x{x[i]=0.3*math.Sin(2*math.Pi*1003.7*float64(i)/rate)}
 if got:=magnitude(x,3000,rate);got>0.0001 {
  t.Fatalf("distant frequency leakage %.9f",got)
 }
}
