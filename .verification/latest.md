# Latest Go verification

- Source commit: ee703b71f2c4bf5945391cf51caebb4e5609f041
- Run: https://github.com/cpu-model/hammond/actions/runs/40217268138
- UTC: 2026-10-25T17:36:39Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.532s
ok  	github.com/cpu-model/hammond/internal/engine	0.456s
ok  	github.com/cpu-model/hammond/internal/generator	0.357s
ok  	github.com/cpu-model/hammond/internal/instrument	0.005s
ok  	github.com/cpu-model/hammond/internal/registration	0.005s
```

## vet output
```text
```

## build output
```text
```
