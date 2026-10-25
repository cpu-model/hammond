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
