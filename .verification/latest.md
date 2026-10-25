# Latest Go verification

- Source commit: 49e6b4c4b873feb35d5ef1610bb3baa4fad6cf77
- Run: https://github.com/cpu-model/hammond/actions/runs/40229956029
- UTC: 2026-10-25T20:16:29Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	4.630s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.014s
ok  	github.com/cpu-model/hammond/internal/electrical	0.481s
ok  	github.com/cpu-model/hammond/internal/engine	0.647s
ok  	github.com/cpu-model/hammond/internal/generator	0.390s
ok  	github.com/cpu-model/hammond/internal/instrument	0.004s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
