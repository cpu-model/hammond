# Hammond — Go tonewheel emulator research

An authenticity-focused Hammond B3 emulator project. macOS is the first
target; the DSP core is platform-independent Go.

## Implemented

- Nominal 91-wheel frequencies derived from documented gear ratios.
- Continuous, shared 12-speed-group mechanical phases (ideal sine output).
- Offline WAV rendering for any of the 91 frequency numbers.
- Source-traceable catalog validation and generator filter topology metadata.
- Explicit 61-key / nine-contact manual routing interface and drawbar controls (factory wiring not yet populated).
- Unit tests and GitHub Actions workflow.

## Run

```sh
go test ./...
go vet ./...
go build ./...
go run ./cmd/hammond -wheel 46 -seconds 3 -out wheel46.wav
```

Wheel 46 is nominally 440 Hz. The rendered tone is an ideal sine, **not yet
an authentic Hammond timbre**.

## Sound-and-behavior milestone

The primary acceptance target is a playable instrument whose audible output and
response match a real Hammond B3, not a component-by-component electrical clone.
The existing electrical pedal circuit is experimental and must not block the
playable signal path.

Implementation order:

1. Render a complete upper/lower manual and pedal mix from shared tonewheels,
   verified contact routing, and adjustable registrations.
2. Add audible key-click and contact timing, including repeatable note-on/off
   tests and audio comparison fixtures.
3. Implement percussion behavior and scanner vibrato/chorus.
4. Calibrate generator spectra, leakage, expression and nonlinear amplification
   against recorded reference instruments.
5. Add optional Leslie processing and validate stereo movement and dynamics.

For every stage, require an audible render plus deterministic automated tests.
Treat undocumented circuit values as calibration parameters, never factory facts.

## Next

Verify mechanical phase offsets for the 48 shaft assemblies; transcribe
manual/pedal contact routing and foldback; model pickups, filtering,
percussion, scanner vibrato, preamp and Leslie 122.

Documentation: [registration](docs/REGISTRATION.md),
[gearing](docs/GEARING.md),
[mechanical model](docs/MECHANICAL.md),
[generator](docs/GENERATOR.md),
[reference status](docs/REFERENCE.md).

**Status:** Research prototype. No acoustic equivalence claim. CI workflow
exists, but its run result must be checked separately.
