# Latest Go verification

- Source commit: 4ed7b5cd096e4fa6bfa1dfd44dc23bfc706d5ca2
- Run: https://github.com/cpu-model/hammond/actions/runs/40229249120
- UTC: 2026-10-25T20:07:34Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	5.653s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.007s
ok  	github.com/cpu-model/hammond/internal/electrical	0.663s
ok  	github.com/cpu-model/hammond/internal/engine	0.912s
ok  	github.com/cpu-model/hammond/internal/generator	0.496s
ok  	github.com/cpu-model/hammond/internal/instrument	0.003s
ok  	github.com/cpu-model/hammond/internal/registration	0.006s
```

## vet output
```text
```

## build output
```text
```
