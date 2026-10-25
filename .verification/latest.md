# Latest Go verification

- Source commit: 8a1e2ab186ade553435b12bdc6b774f09c1e75a8
- Run: https://github.com/cpu-model/hammond/actions/runs/40228273997
- UTC: 2026-10-25T19:55:14Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	2.182s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.012s
ok  	github.com/cpu-model/hammond/internal/electrical	0.441s
ok  	github.com/cpu-model/hammond/internal/engine	0.472s
ok  	github.com/cpu-model/hammond/internal/generator	0.337s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.002s
```

## vet output
```text
```

## build output
```text
```
