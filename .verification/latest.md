# Latest Go verification

- Source commit: 6c0cb7ed73524ff3b978444a33ec8ffc0b66e439
- Run: https://github.com/cpu-model/hammond/actions/runs/40220038045
- UTC: 2026-10-25T18:11:35Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.452s
ok  	github.com/cpu-model/hammond/internal/engine	0.477s
ok  	github.com/cpu-model/hammond/internal/generator	0.369s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.007s
```

## vet output
```text
```

## build output
```text
```
