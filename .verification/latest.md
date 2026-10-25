# Latest Go verification

- Source commit: e5ddadfe9b5e8b439c1026c701ac98f7eae2f364
- Run: https://github.com/cpu-model/hammond/actions/runs/40212667209
- UTC: 2026-10-25T16:39:06Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.006s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.441s
ok  	github.com/cpu-model/hammond/internal/generator	0.354s
ok  	github.com/cpu-model/hammond/internal/instrument	0.006s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
