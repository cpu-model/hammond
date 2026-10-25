# Latest Go verification

- Source commit: 11a19a8b25660ab6a339b3a440c73b6174fc1a4e
- Run: https://github.com/cpu-model/hammond/actions/runs/40213122542
- UTC: 2026-10-25T16:44:50Z

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
ok  	github.com/cpu-model/hammond/internal/engine	0.430s
ok  	github.com/cpu-model/hammond/internal/generator	0.326s
ok  	github.com/cpu-model/hammond/internal/instrument	0.006s
ok  	github.com/cpu-model/hammond/internal/registration	0.005s
```

## vet output
```text
```

## build output
```text
```
