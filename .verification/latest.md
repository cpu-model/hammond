# Latest Go verification

- Source commit: 26c12de0c82ff9cdb44ba743f1b30649edcf676d
- Run: https://github.com/cpu-model/hammond/actions/runs/40226825823
- UTC: 2026-10-25T19:36:46Z

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
    main_test.go:16: fundamental amplitude 0.249999999, want 0.25
FAIL
FAIL	github.com/cpu-model/hammond/cmd/wavcompare	0.005s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/electrical	0.381s
ok  	github.com/cpu-model/hammond/internal/engine	0.534s
ok  	github.com/cpu-model/hammond/internal/generator	0.396s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.005s
FAIL
```

## vet output
```text
```

## build output
```text
```
