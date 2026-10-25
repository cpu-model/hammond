# Latest Go verification

- Source commit: d75998062478cd8f95585529ad2bc1a375eec725
- Run: https://github.com/cpu-model/hammond/actions/runs/40230362786
- UTC: 2026-10-25T20:21:44Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	6.459s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.005s
ok  	github.com/cpu-model/hammond/internal/electrical	0.652s
ok  	github.com/cpu-model/hammond/internal/engine	0.855s
ok  	github.com/cpu-model/hammond/internal/generator	0.490s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.007s
```

## vet output
```text
```

## build output
```text
```
