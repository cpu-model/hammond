# Shared speed-group model

`internal/generator/mechanical.go` advances **12 continuous mechanical speed-group phases**
rather than 91 independent oscillator phases. Each wheel waveform is evaluated
using its number of teeth and the corresponding group phase.

For frequency numbers 1..84:
- group = (number-1) modulo 12
- teeth = 2 << ((number-1)/12)

For frequency numbers 85..91:
- teeth = 192
- groups = F, F#, G, G#, A, A#, B (indices 5..11)

The common 1200 RPM drive and the documented gear ratios establish nominal
pitch. The current phase origins are all zero by convention. **This does not
model the 48 individual shaft phase offsets, nor the relative placement of
pickup magnets and tonewheel teeth.** Those parameters are not established by
the frequency table. Phase-accurate mechanical authenticity remains open.

No waveform distortion, leakage, key click, filters, drawbars, vibrato or
Leslie are simulated. Output remains ideal sine waves.

The old `generator.Bank` remains as a generic oscillator test utility.

## Regression check

The generator's upper-wheel speed-group offset was corrected to index 5
for wheel 85 (F group). The new cross-component test compares all 91
rendered signals against the independently calculated nominal frequency
model across 2,400 samples. This is a consistency test, not a physical
measurement of a Hammond generator.
