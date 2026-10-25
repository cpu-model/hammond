# Latest Go verification

- Source commit: 6180998f48f8fe2f3e54cf8ba75b66f3860a6416
- Run: https://github.com/cpu-model/hammond/actions/runs/40218469433
- UTC: 2026-10-25T17:51:44Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.538s
ok  	github.com/cpu-model/hammond/internal/engine	0.513s
ok  	github.com/cpu-model/hammond/internal/generator	0.394s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.005s
```

## vet output
```text
```

## build output
```text
```
