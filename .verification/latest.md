# Latest Go verification

- Source commit: 621924da43058fd623fc63b1801433db4ffd4706
- Run: https://github.com/cpu-model/hammond/actions/runs/40212340996
- UTC: 2026-10-25T16:34:59Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.006s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.364s
ok  	github.com/cpu-model/hammond/internal/generator	0.301s
ok  	github.com/cpu-model/hammond/internal/instrument	0.004s
ok  	github.com/cpu-model/hammond/internal/registration	0.005s
```

## vet output
```text
```

## build output
```text
```
