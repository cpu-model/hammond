# Latest Go verification

- Source commit: 946e5949a685edc82e3a5130adcdfb874a278aee
- Run: https://github.com/cpu-model/hammond/actions/runs/40205741406
- UTC: 2026-10-25T15:11:33Z

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
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.180s
ok  	github.com/cpu-model/hammond/internal/generator	0.254s
ok  	github.com/cpu-model/hammond/internal/instrument	0.001s
ok  	github.com/cpu-model/hammond/internal/registration	0.002s
```

## vet output
```text
```

## build output
```text
```
