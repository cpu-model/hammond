# Gear ratios — US 60 Hz B3 research reference

The reference table at https://www.goodeveca.net/RotorOrgan/ToneWheelSpec.html
lists 12 chromatic drive/driven tooth pairs for a 1200 RPM drive shaft:

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

For generator frequency numbers 1–84, tooth counts are 2, 4, 8,
16, 32, 64, 128 by ascending octave. The model computes
frequency = (1200 / 60) * (drive / driven) * teeth.

This yields A440 at terminal 46 and differs slightly from equal temperament.
Numbers 85–91 use 192-tooth wheels, and their pitch mapping is **not**
implemented here: those wheels are not simply the next octave.

The gearing data comes from a secondary technical table, not a direct
transcription of a factory gear drawing. Independent primary-source
verification remains required before declaring the instrument specification
complete.

This stage models only frequency, not mechanical coupling, pickup output,
filters, manual wiring, or audible B3 character.
