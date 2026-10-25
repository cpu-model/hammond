# Latest Go verification

- Source commit: ba44d56cb98454f6656ed3e7c98f836122b6cfdd
- Run: https://github.com/cpu-model/hammond/actions/runs/40205577799
- UTC: 2026-10-25T15:09:32Z

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
ok  	github.com/cpu-model/hammond/internal/engine	0.195s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/generator	0.279s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
