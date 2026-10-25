# Latest Go verification

- Source commit: 9d2c32801b53828edf39c03c6dcf770f8654fb2a
- Run: https://github.com/cpu-model/hammond/actions/runs/40230433530
- UTC: 2026-10-25T20:22:29Z

| Check | Exit code |
| --- | ---: |
| go test ./... | 1 |
| go vet ./... | 0 |
| go build ./... | 0 |

## test output
```text
?   	github.com/cpu-model/hammond/cmd/contactdemo	[no test files]
?   	github.com/cpu-model/hammond/cmd/hammond	[no test files]
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/cmd/manualdemo	4.342s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.006s
ok  	github.com/cpu-model/hammond/internal/electrical	0.422s
ok  	github.com/cpu-model/hammond/internal/engine	0.487s
ok  	github.com/cpu-model/hammond/internal/generator	0.357s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
--- FAIL: TestTonewheelSwitchingResidual (0.00s)
    percussion_tonewheel_residual_test.go:42: 3ms transition did not reduce residual peak: immediate=0.05758770738600319 smooth=0.05761622727202991
FAIL
FAIL	github.com/cpu-model/hammond/internal/registration	0.006s
FAIL
```

## vet output
```text
```

## build output
```text
```
