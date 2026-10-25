# Latest Go verification

- Source commit: cc8f5e7f78ec25b791874c0c78a6ee293d6de144
- Run: https://github.com/cpu-model/hammond/actions/runs/40230718746
- UTC: 2026-10-25T20:26:10Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	5.498s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.006s
ok  	github.com/cpu-model/hammond/internal/electrical	0.480s
ok  	github.com/cpu-model/hammond/internal/engine	0.617s
ok  	github.com/cpu-model/hammond/internal/generator	0.377s
ok  	github.com/cpu-model/hammond/internal/instrument	0.006s
--- FAIL: TestNinthDrawbarSwitchingHighFrequencyEnergy (0.02s)
    percussion_spectral_test.go:61: high-frequency transition energy not reduced: immediate=0 smooth=138.2616640208104
FAIL
FAIL	github.com/cpu-model/hammond/internal/registration	0.021s
FAIL
```

## vet output
```text
```

## build output
```text
```
