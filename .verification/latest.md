# Latest Go verification

- Source commit: d6d34c3ed3f967f94a8e1bd04ae1d071b819b304
- Run: https://github.com/cpu-model/hammond/actions/runs/40214395650
- UTC: 2026-10-25T17:01:07Z

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
ok  	github.com/cpu-model/hammond/internal/engine	0.406s
ok  	github.com/cpu-model/hammond/internal/generator	0.319s
ok  	github.com/cpu-model/hammond/internal/instrument	0.007s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
