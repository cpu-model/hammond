# Latest Go verification

- Source commit: 0ec871b52c7c779eaae0eeb23c7ebedc7b982bfe
- Run: https://github.com/cpu-model/hammond/actions/runs/40213445149
- UTC: 2026-10-25T16:48:46Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.013s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.268s
ok  	github.com/cpu-model/hammond/internal/generator	0.220s
ok  	github.com/cpu-model/hammond/internal/instrument	0.006s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
