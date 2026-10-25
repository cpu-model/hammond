# Latest Go verification

- Source commit: 325a988865f0c3ce2ecbe3dc3a2a5230dd988770
- Run: https://github.com/cpu-model/hammond/actions/runs/40228500394
- UTC: 2026-10-25T19:58:15Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	3.803s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.005s
ok  	github.com/cpu-model/hammond/internal/electrical	0.677s
ok  	github.com/cpu-model/hammond/internal/engine	0.900s
ok  	github.com/cpu-model/hammond/internal/generator	0.527s
ok  	github.com/cpu-model/hammond/internal/instrument	0.004s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
