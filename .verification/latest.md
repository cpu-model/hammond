# Latest Go verification

- Source commit: bc0c922eefc1021eeb4d1056f0797a007630504a
- Run: https://github.com/cpu-model/hammond/actions/runs/40220724404
- UTC: 2026-10-25T18:20:13Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.483s
ok  	github.com/cpu-model/hammond/internal/engine	0.455s
ok  	github.com/cpu-model/hammond/internal/generator	0.302s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
