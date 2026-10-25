# Nominal gearing for a US 60 Hz Hammond B3

Secondary reference: https://www.goodeveca.net/RotorOrgan/ToneWheelSpec.html
Cross-check: https://www.dairiki.org/HammondWiki/GearRatio?version=6

The drive shaft rotates at 1200 RPM. Each of the 12 speed groups has
a drive/driven tooth pair, in chromatic order:

| Note | Drive | Driven |
|---|---:|---:|
| C | 85 | 104 |
| C# | 71 | 82 |
| D | 67 | 73 |
| D# | 105 | 108 |
| E | 103 | 100 |
| F | 84 | 77 |
| F# | 74 | 64 |
| G | 98 | 80 |
| G# | 96 | 74 |
| A | 88 | 64 |
| A# | 67 | 46 |
| B | 108 | 70 |

Nominal frequency = (1200/60) * drive/driven * wheel teeth.

- Generator frequency numbers 1..84: 12 pitches per octave, with
  2, 4, 8, 16, 32, 64 and 128 teeth.
- Generator frequency numbers 85..91: 192 teeth, with **gear groups
  F through B** mapped to **output pitches C7 through F#7**.
- Number 46 is exactly 440 Hz; number 86 is 4440 Hz.

The Go model now calculates all 91 nominal frequencies. This does
**not** validate individual shaft phasing, pickup amplitude/waveform,
tonewheel defects, generator filtering or keyboard wiring. Independent
primary-source confirmation of the gear table is still desirable.

Render an isolated test wheel (ideal sine, not authentic B3 timbre):

```sh
go run ./cmd/hammond -wheel 86 -seconds 3 -out wheel86.wav
```
