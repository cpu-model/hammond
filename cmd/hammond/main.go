package main

import (
 "flag"
 "log"
 "os"

 "github.com/cpu-model/hammond/internal/generator"
 "github.com/cpu-model/hammond/internal/instrument"
 "github.com/cpu-model/hammond/internal/wav"
)

func main() {
 out := flag.String("out", "prototype.wav", "output WAV path")
 seconds := flag.Int("seconds", 3, "duration in seconds")
 wheel := flag.Int("wheel", 46, "generator frequency number (1..91)")
 flag.Parse()
 if *seconds <= 0 || *seconds > 120 { log.Fatal("seconds must be 1..120") }
 if *wheel < 1 || *wheel > instrument.TonewheelCount { log.Fatal("wheel must be 1..91") }
 const rate = 48000
 frequencies := instrument.MechanicalFrequencies()
 bank, err := generator.New(rate, frequencies[:])
 if err != nil { log.Fatal(err) }
 samples := make([]float64, rate * *seconds)
 frame := make([]float64, bank.Count())
 for i := range samples {
  bank.Next(frame)
  samples[i] = 0.25 * frame[*wheel-1]
 }
 f, err := os.Create(*out)
 if err != nil { log.Fatal(err) }
 if err := wav.WriteMono16(f, rate, samples); err != nil { _ = f.Close(); log.Fatal(err) }
 if err := f.Close(); err != nil { log.Fatal(err) }
 log.Printf("wrote %s (wheel %d, nominal %.3f Hz, ideal sine only)", *out, *wheel, frequencies[*wheel-1])
}
