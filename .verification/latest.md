# Latest Go verification

- Source commit: 35eea51a9cda7e4b4d6535acf9f84442a789b8d6
- Run: https://github.com/cpu-model/hammond/actions/runs/40210989806
- UTC: 2026-10-25T16:18:03Z

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
ok  	github.com/cpu-model/hammond/internal/engine	0.289s
ok  	github.com/cpu-model/hammond/internal/generator	0.228s
ok  	github.com/cpu-model/hammond/internal/instrument	0.006s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
