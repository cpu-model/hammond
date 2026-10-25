# Latest Go verification

- Source commit: 310b9bb52522ec1c7de2e740834633225051559e
- Run: https://github.com/cpu-model/hammond/actions/runs/40220806477
- UTC: 2026-10-25T18:21:11Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.434s
ok  	github.com/cpu-model/hammond/internal/engine	0.392s
ok  	github.com/cpu-model/hammond/internal/generator	0.318s
ok  	github.com/cpu-model/hammond/internal/instrument	0.001s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
