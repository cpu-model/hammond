# Latest Go verification

- Source commit: 5462257fe461b274e1978e80ee165ad7957dae8b
- Run: https://github.com/cpu-model/hammond/actions/runs/40230067453
- UTC: 2026-10-25T20:17:54Z

| Check | Exit code |
| --- | ---: |
| go test ./... | 0 |
| go vet ./... | 0 |
| go build ./... | 0 |

## test output
```text
?   	github.com/cpu-model/hammond/cmd/contactdemo	[no test files]
?   	github.com/cpu-model/hammond/cmd/hammond	[no test files]
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/cmd/manualdemo	4.971s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.006s
ok  	github.com/cpu-model/hammond/internal/electrical	0.683s
ok  	github.com/cpu-model/hammond/internal/engine	0.831s
ok  	github.com/cpu-model/hammond/internal/generator	0.506s
ok  	github.com/cpu-model/hammond/internal/instrument	0.003s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
