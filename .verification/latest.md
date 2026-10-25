# Latest Go verification

- Source commit: af1f688cb13c23bd144842ba745a371f4ff51bab
- Run: https://github.com/cpu-model/hammond/actions/runs/40222403766
- UTC: 2026-10-25T18:41:16Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.399s
ok  	github.com/cpu-model/hammond/internal/engine	0.394s
ok  	github.com/cpu-model/hammond/internal/generator	0.313s
ok  	github.com/cpu-model/hammond/internal/instrument	0.005s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
