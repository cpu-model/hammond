package main

import (
 "flag"
 "log"
 "os"

 "github.com/cpu-model/hammond/internal/generator"
 "github.com/cpu-model/hammond/internal/wav"
)

func main() {
 out := flag.String("out", "prototype.wav", "output WAV path")
 seconds := flag.Int("seconds", 3, "duration in seconds")
 flag.Parse()
 if *seconds <= 0 || *seconds > 120 { log.Fatal("seconds must be 1..120") }
 const rate = 48000
 // Deliberately synthetic test tone, NOT a verified Hammond tonewheel.
 bank, err := generator.New(rate, []float64{440})
 if err != nil { log.Fatal(err) }
 samples := make([]float64, rate * *seconds)
 frame := make([]float64, bank.Count())
 for i := range samples {
  bank.Next(frame)
  samples[i] = 0.25 * frame[0]
 }
 f, err := os.Create(*out)
 if err != nil { log.Fatal(err) }
 if err := wav.WriteMono16(f, rate, samples); err != nil { _ = f.Close(); log.Fatal(err) }
 if err := f.Close(); err != nil { log.Fatal(err) }
 log.Printf("wrote %s (synthetic 440 Hz prototype)", *out)
}
