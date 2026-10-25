# Latest Go verification

- Source commit: 9778518b9032192178cf365fface37fc9a148d5d
- Run: https://github.com/cpu-model/hammond/actions/runs/40228109874
- UTC: 2026-10-25T19:53:08Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	2.026s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.009s
ok  	github.com/cpu-model/hammond/internal/electrical	0.654s
ok  	github.com/cpu-model/hammond/internal/engine	0.790s
ok  	github.com/cpu-model/hammond/internal/generator	0.497s
ok  	github.com/cpu-model/hammond/internal/instrument	0.005s
ok  	github.com/cpu-model/hammond/internal/registration	0.005s
```

## vet output
```text
```

## build output
```text
```
