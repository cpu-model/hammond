# Latest Go verification

- Source commit: d87113a9b4160d67c7fb8c5956e530111de865a6
- Run: https://github.com/cpu-model/hammond/actions/runs/40206260735
- UTC: 2026-10-25T15:18:08Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.002s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.293s
ok  	github.com/cpu-model/hammond/internal/generator	0.238s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.005s
```

## vet output
```text
```

## build output
```text
```
