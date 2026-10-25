# Latest Go verification

- Source commit: d21352f2e725bd63c501fc9cb3cd626a18d0d62c
- Run: https://github.com/cpu-model/hammond/actions/runs/40227249814
- UTC: 2026-10-25T19:42:10Z

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
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.009s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/electrical	0.582s
ok  	github.com/cpu-model/hammond/internal/engine	0.677s
ok  	github.com/cpu-model/hammond/internal/generator	0.497s
ok  	github.com/cpu-model/hammond/internal/instrument	0.003s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
