# Latest Go verification

- Source commit: 4e496b506c49c185fb44b4f429253778b9fda25c
- Run: https://github.com/cpu-model/hammond/actions/runs/40206748249
- UTC: 2026-10-25T15:24:21Z

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
ok  	github.com/cpu-model/hammond/internal/engine	0.279s
ok  	github.com/cpu-model/hammond/internal/generator	0.223s
ok  	github.com/cpu-model/hammond/internal/instrument	0.001s
ok  	github.com/cpu-model/hammond/internal/registration	0.002s
```

## vet output
```text
```

## build output
```text
```
