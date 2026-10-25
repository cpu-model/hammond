# Latest Go verification

- Source commit: f282834505793908bee13144653924f53beb5d3d
- Run: https://github.com/cpu-model/hammond/actions/runs/40212723473
- UTC: 2026-10-25T16:39:45Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.005s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.241s
ok  	github.com/cpu-model/hammond/internal/generator	0.192s
ok  	github.com/cpu-model/hammond/internal/instrument	0.001s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
