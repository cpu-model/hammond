# Latest Go verification

- Source commit: 0705ca73cf901d2f90b79e23c489955b43a0eeed
- Run: https://github.com/cpu-model/hammond/actions/runs/40206815813
- UTC: 2026-10-25T15:25:15Z

| Check | Exit code |
| --- | ---: |
| go test ./... | 0 |
| go vet ./... | 0 |
| go build ./... | 0 |

## test output
```text
?   	github.com/cpu-model/hammond/cmd/contactdemo	[no test files]
?   	github.com/cpu-model/hammond/cmd/hammond	[no test files]
?   	github.com/cpu-model/hammond/cmd/manualdemo	[no test files]
ok  	github.com/cpu-model/hammond/internal/electrical	0.011s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.410s
ok  	github.com/cpu-model/hammond/internal/generator	0.317s
ok  	github.com/cpu-model/hammond/internal/instrument	0.004s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
