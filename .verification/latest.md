# Latest Go verification

- Source commit: 4dbdfbb277e2e895960bce504e548e9b4de9aa9b
- Run: https://github.com/cpu-model/hammond/actions/runs/40210604484
- UTC: 2026-10-25T16:13:12Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.003s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.260s
ok  	github.com/cpu-model/hammond/internal/generator	0.210s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.002s
```

## vet output
```text
```

## build output
```text
```
