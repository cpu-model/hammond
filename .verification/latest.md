# Latest Go verification

- Source commit: 739a5f416c0f396c16f1ad4a1293f1688acf33ba
- Run: https://github.com/cpu-model/hammond/actions/runs/40228203587
- UTC: 2026-10-25T19:54:21Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	2.241s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.013s
ok  	github.com/cpu-model/hammond/internal/electrical	0.615s
ok  	github.com/cpu-model/hammond/internal/engine	0.775s
ok  	github.com/cpu-model/hammond/internal/generator	0.503s
ok  	github.com/cpu-model/hammond/internal/instrument	0.003s
ok  	github.com/cpu-model/hammond/internal/registration	0.006s
```

## vet output
```text
```

## build output
```text
```
