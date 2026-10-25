# Latest Go verification

- Source commit: 648bbde8c29542ad89b3c44c6c328709af92b400
- Run: https://github.com/cpu-model/hammond/actions/runs/40220568420
- UTC: 2026-10-25T18:18:13Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.398s
ok  	github.com/cpu-model/hammond/internal/engine	0.378s
ok  	github.com/cpu-model/hammond/internal/generator	0.319s
ok  	github.com/cpu-model/hammond/internal/instrument	0.004s
ok  	github.com/cpu-model/hammond/internal/registration	0.006s
```

## vet output
```text
```

## build output
```text
```
