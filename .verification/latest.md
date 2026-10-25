# Latest Go verification

- Source commit: 10c84529c42eb484570865905e0036595f11a040
- Run: https://github.com/cpu-model/hammond/actions/runs/40208123379
- UTC: 2026-10-25T15:41:42Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.006s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.318s
ok  	github.com/cpu-model/hammond/internal/generator	0.258s
ok  	github.com/cpu-model/hammond/internal/instrument	0.004s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
