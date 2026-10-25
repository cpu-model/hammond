# Latest Go verification

- Source commit: 887f96b8bde62e00e50d1204468aa841e9e92101
- Run: https://github.com/cpu-model/hammond/actions/runs/40222984051
- UTC: 2026-10-25T18:48:37Z

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
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/electrical	0.364s
ok  	github.com/cpu-model/hammond/internal/engine	0.344s
ok  	github.com/cpu-model/hammond/internal/generator	0.235s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
