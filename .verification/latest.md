# Latest Go verification

- Source commit: 36776847f5bf2401d74f3282a85745fea4f38a04
- Run: https://github.com/cpu-model/hammond/actions/runs/40214187058
- UTC: 2026-10-25T16:58:28Z

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
ok  	github.com/cpu-model/hammond/internal/engine	0.468s
ok  	github.com/cpu-model/hammond/internal/generator	0.353s
ok  	github.com/cpu-model/hammond/internal/instrument	0.003s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
