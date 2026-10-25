# Verification protocol

The repository has a GitHub Actions workflow running:

```sh
go test ./...
go vet ./...
go build ./...
```

## End-to-end regression

`internal/engine/audio_integration_test.go` exercises:

1. The transcribed B3 routing table.
2. Two timed manuals and one continuously running mechanical generator.
3. A drawbar at position 8 on both manuals.
4. The complete PCM WAV encoder, including RIFF header and payload length.

It also checks that two manuals using the same physical generator
signal produce twice the single-manual amplitude. This is a **digital
consistency check**, not a measured B3 electrical response.

`internal/generator/mechanical_reference_test.go` compares all 91
oscillator outputs to the independently derived nominal frequency table.

## Limits

Passing these tests does not establish acoustic authenticity.
No primary-source electrical calibration, key-click transient,
tonewheel pickup spectrum, Leslie, vibrato or real-time device adapter
is included.

CI results should be confirmed in GitHub Actions. A workflow file
existing in the repository is not evidence of a passing run.
