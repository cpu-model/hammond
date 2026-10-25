# Latest Go verification

- Source commit: 8bc709b9dc8310e3ecc5a654208a53f309a9ee00
- Run: https://github.com/cpu-model/hammond/actions/runs/40212926418
- UTC: 2026-10-25T16:42:19Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.003s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.291s
ok  	github.com/cpu-model/hammond/internal/generator	0.224s
ok  	github.com/cpu-model/hammond/internal/instrument	0.004s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
