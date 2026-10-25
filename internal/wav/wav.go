package wav

import (
 "encoding/binary"
 "errors"
 "io"
 "math"
)

// WriteMono16 writes a little-endian PCM WAV file.
func WriteMono16(w io.Writer, sampleRate uint32, samples []float64) error {
 if sampleRate == 0 || sampleRate > ^uint32(0)/2 || uint64(len(samples)) > uint64(^uint32(0)-36)/2 { return errors.New("invalid WAV dimensions") }
 dataSize := uint32(len(samples)) * 2
 if _, err := w.Write([]byte("RIFF")); err != nil { return err }
 if err := binary.Write(w, binary.LittleEndian, uint32(36)+dataSize); err != nil { return err }
 if _, err := w.Write([]byte("WAVEfmt ")); err != nil { return err }
 for _, v := range []any{uint32(16), uint16(1), uint16(1), sampleRate, sampleRate*2, uint16(2), uint16(16)} {
  if err := binary.Write(w, binary.LittleEndian, v); err != nil { return err }
 }
 if _, err := w.Write([]byte("data")); err != nil { return err }
 if err := binary.Write(w, binary.LittleEndian, dataSize); err != nil { return err }
 for _, x := range samples {
  if math.IsNaN(x) || math.IsInf(x, 0) { return errors.New("nonfinite audio sample") }
  x = math.Max(-1, math.Min(1, x))
  if err := binary.Write(w, binary.LittleEndian, int16(math.Round(x*32767))); err != nil { return err }
 }
 return nil
}
