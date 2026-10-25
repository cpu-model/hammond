# Latest Go verification

- Source commit: 85a3da54c25e5fb645d556ef118a657815aeee01
- Run: https://github.com/cpu-model/hammond/actions/runs/40212420705
- UTC: 2026-10-25T16:36:02Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.007s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.377s
ok  	github.com/cpu-model/hammond/internal/generator	0.306s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
