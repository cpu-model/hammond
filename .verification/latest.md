# Latest Go verification

- Source commit: 07cabee03f8ff11aa5b7cf609bda18b6a5c0e361
- Run: https://github.com/cpu-model/hammond/actions/runs/40221584151
- UTC: 2026-10-25T18:30:58Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.555s
ok  	github.com/cpu-model/hammond/internal/engine	0.499s
ok  	github.com/cpu-model/hammond/internal/generator	0.398s
ok  	github.com/cpu-model/hammond/internal/instrument	0.003s
ok  	github.com/cpu-model/hammond/internal/registration	0.005s
```

## vet output
```text
```

## build output
```text
```
