# Latest Go verification

- Source commit: cf5a95ccdd643eb63a5ebd9413d535dc233b85a0
- Run: https://github.com/cpu-model/hammond/actions/runs/40221252197
- UTC: 2026-10-25T18:26:46Z

| Check | Exit code |
| --- | ---: |
| go test ./... | 0 |
| go vet ./... | 0 |
| go build ./... | 0 |

## test output
```text
?   	github.com/cpu-model/hammond/cmd/contactdemo	[no test files]
?   	github.com/cpu-model/hammond/cmd/manualdemo	[no test files]
?   	github.com/cpu-model/hammond/cmd/hammond	[no test files]
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/electrical	0.410s
ok  	github.com/cpu-model/hammond/internal/engine	0.379s
ok  	github.com/cpu-model/hammond/internal/generator	0.319s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
