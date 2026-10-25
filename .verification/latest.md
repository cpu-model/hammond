# Latest Go verification

- Source commit: 9dfc89094cc2d70d0a965ccd22a5dcae837b2448
- Run: https://github.com/cpu-model/hammond/actions/runs/40220961753
- UTC: 2026-10-25T18:23:08Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.471s
ok  	github.com/cpu-model/hammond/internal/engine	0.483s
ok  	github.com/cpu-model/hammond/internal/generator	0.403s
ok  	github.com/cpu-model/hammond/internal/instrument	0.004s
ok  	github.com/cpu-model/hammond/internal/registration	0.008s
```

## vet output
```text
```

## build output
```text
```
