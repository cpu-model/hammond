# Latest Go verification

- Source commit: b7ae35082ae9a0c43ad9d3d3458dcf5a359657ce
- Run: https://github.com/cpu-model/hammond/actions/runs/40229102335
- UTC: 2026-10-25T20:05:44Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	5.640s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.021s
ok  	github.com/cpu-model/hammond/internal/electrical	0.652s
ok  	github.com/cpu-model/hammond/internal/engine	0.885s
ok  	github.com/cpu-model/hammond/internal/generator	0.489s
ok  	github.com/cpu-model/hammond/internal/instrument	0.003s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
