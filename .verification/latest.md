# Latest Go verification

- Source commit: b4c7f64c030c605fac8e1736ef8912aca5bd37b9
- Run: https://github.com/cpu-model/hammond/actions/runs/40224025728
- UTC: 2026-10-25T19:01:48Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.575s
ok  	github.com/cpu-model/hammond/internal/engine	0.642s
ok  	github.com/cpu-model/hammond/internal/generator	0.467s
ok  	github.com/cpu-model/hammond/internal/instrument	0.003s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
