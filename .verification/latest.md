# Latest Go verification

- Source commit: af2527e6b88deae16f8215002616641524bb91c5
- Run: https://github.com/cpu-model/hammond/actions/runs/40227811273
- UTC: 2026-10-25T19:49:19Z

| Check | Exit code |
| --- | ---: |
| go test ./... | 0 |
| go vet ./... | 0 |
| go build ./... | 0 |

## test output
```text
?   	github.com/cpu-model/hammond/cmd/contactdemo	[no test files]
?   	github.com/cpu-model/hammond/cmd/hammond	[no test files]
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/cmd/manualdemo	1.630s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.005s
ok  	github.com/cpu-model/hammond/internal/electrical	0.523s
ok  	github.com/cpu-model/hammond/internal/engine	0.588s
ok  	github.com/cpu-model/hammond/internal/generator	0.366s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
