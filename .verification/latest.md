# Latest Go verification

- Source commit: f295c7a14c0db214c945e32d6f3eeb855d335f41
- Run: https://github.com/cpu-model/hammond/actions/runs/40210824547
- UTC: 2026-10-25T16:15:57Z

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
ok  	github.com/cpu-model/hammond/internal/engine	0.325s
ok  	github.com/cpu-model/hammond/internal/generator	0.240s
ok  	github.com/cpu-model/hammond/internal/instrument	0.001s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
