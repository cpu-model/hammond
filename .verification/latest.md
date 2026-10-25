# Latest Go verification

- Source commit: 7e7a1a4241f9d6a890edee4755250bb1f6c3f010
- Run: https://github.com/cpu-model/hammond/actions/runs/40227453625
- UTC: 2026-10-25T19:44:47Z

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
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.004s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/electrical	0.480s
ok  	github.com/cpu-model/hammond/internal/engine	0.545s
ok  	github.com/cpu-model/hammond/internal/generator	0.367s
ok  	github.com/cpu-model/hammond/internal/instrument	0.005s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
