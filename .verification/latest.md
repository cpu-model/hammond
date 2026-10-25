# Latest Go verification

- Source commit: cc0ce39976f541fb7ccfafe5d0d88e3b890619c6
- Run: https://github.com/cpu-model/hammond/actions/runs/40226900191
- UTC: 2026-10-25T19:37:43Z

| Check | Exit code |
| --- | ---: |
| go test ./... | 1 |
| go vet ./... | 0 |
| go build ./... | 0 |

## test output
```text
?   	github.com/cpu-model/hammond/cmd/contactdemo	[no test files]
?   	github.com/cpu-model/hammond/cmd/hammond	[no test files]
?   	github.com/cpu-model/hammond/cmd/manualdemo	[no test files]
--- FAIL: TestMagnitudeKnownSine (0.00s)
    main_test.go:17: unexpected second harmonic 0.000000005
FAIL
FAIL	github.com/cpu-model/hammond/cmd/wavcompare	0.013s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/electrical	0.618s
ok  	github.com/cpu-model/hammond/internal/engine	0.672s
ok  	github.com/cpu-model/hammond/internal/generator	0.512s
ok  	github.com/cpu-model/hammond/internal/instrument	0.003s
ok  	github.com/cpu-model/hammond/internal/registration	0.014s
FAIL
```

## vet output
```text
```

## build output
```text
```
