# Latest Go verification

- Source commit: 40c1431fd714685cfe8d1d60eb0024455c5f9c0c
- Run: https://github.com/cpu-model/hammond/actions/runs/40218102621
- UTC: 2026-10-25T17:47:05Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.438s
ok  	github.com/cpu-model/hammond/internal/engine	0.412s
ok  	github.com/cpu-model/hammond/internal/generator	0.302s
ok  	github.com/cpu-model/hammond/internal/instrument	0.003s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
