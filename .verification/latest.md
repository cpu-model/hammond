# Latest Go verification

- Source commit: 4ae231aafde8c6f8ebd09e969d9286ef459ae082
- Run: https://github.com/cpu-model/hammond/actions/runs/40230842616
- UTC: 2026-10-25T20:27:50Z

| Check | Exit code |
| --- | ---: |
| go test ./... | 0 |
| go vet ./... | 0 |
| go build ./... | 0 |

## test output
```text
?   	github.com/cpu-model/hammond/cmd/contactdemo	[no test files]
?   	github.com/cpu-model/hammond/cmd/hammond	[no test files]
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/cmd/manualdemo	6.561s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.015s
ok  	github.com/cpu-model/hammond/internal/electrical	0.677s
ok  	github.com/cpu-model/hammond/internal/engine	0.897s
ok  	github.com/cpu-model/hammond/internal/generator	0.523s
ok  	github.com/cpu-model/hammond/internal/instrument	0.003s
ok  	github.com/cpu-model/hammond/internal/registration	0.037s
```

## vet output
```text
```

## build output
```text
```
