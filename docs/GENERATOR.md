# Hammond B3 generator — verified facts and unresolved data

## Documented facts

Source: Hammond B3/C3 Service Manual, section 5, paragraphs 5-4 through 5-13:
https://synthmania.com/wp-content/uploads/2020/05/Hammond-Organ-B3-C3-Service-Manual.pdf

- 91 active tonewheels in 48 two-wheel mechanical assemblies (five blank positions).
- 12 mechanical speed groups; four assemblies per speed group.
- Active tooth counts: 12 each of 2, 4, 8, 16, 32, 64, 128; 7 of 192.
- Two-wheel assemblies: 2+32, 4+64, 8+128, and 16+192 (five 16-tooth wheels pair with blanks).
- Frequencies are assigned terminal numbers 1 through 91.
- The 60 Hz synchronous run motor has a stated 1200 rpm synchronous speed.
- Filter networks: 1–43 resistive shunt; 44–48 reactor only; 49–54 reactor with 0.255 µF capacitor; 55–91 reactor with 0.105 µF capacitor.
- Mechanical springs reduce motor-speed pulsation reaching the wheels.

The motor RPM is **not** a tonewheel shaft RPM. Derive each wheel frequency only after the gear train and shaft speeds are verified.

## Not yet established

- Exact 12 gear ratios and their association with frequency numbers.
- Full 91-frequency reference table with traceable rows.
- Per-wheel pickup waveform, amplitude, phase and filter response.
- Exact upper/lower manual and pedal wiring.

Do not infer an instrument-ready frequency table from equal temperament.
Do not use this prototype as a reference for acoustic authenticity.

## Current implementation

`internal/instrument/filter.go` classifies the documented filter topology but **does not simulate** it.
`internal/generator/bank.go` provides phase-continuous ideal sine oscillators for testing.
