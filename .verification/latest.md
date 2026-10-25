# Latest Go verification

- Source commit: 34679489a4f1e53dc23ea976900f1edcb058a059
- Run: https://github.com/cpu-model/hammond/actions/runs/40230537406
- UTC: 2026-10-25T20:23:49Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	4.339s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.006s
ok  	github.com/cpu-model/hammond/internal/electrical	0.375s
ok  	github.com/cpu-model/hammond/internal/engine	0.517s
ok  	github.com/cpu-model/hammond/internal/generator	0.309s
ok  	github.com/cpu-model/hammond/internal/instrument	0.006s
--- FAIL: TestTonewheelSwitchingResidual (0.00s)
    percussion_tonewheel_residual_test.go:48: 3ms transition did not reduce residual peak: immediate=0 smooth=0.006944444444444376
FAIL
FAIL	github.com/cpu-model/hammond/internal/registration	0.004s
FAIL
```

## vet output
```text
```

## build output
```text
```
