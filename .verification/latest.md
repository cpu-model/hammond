# Latest Go verification

- Source commit: 5802d1a6d2c356cb7ced5a3d87d23b46b2e14427
- Run: https://github.com/cpu-model/hammond/actions/runs/40227980987
- UTC: 2026-10-25T19:51:25Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	1.751s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.006s
ok  	github.com/cpu-model/hammond/internal/electrical	0.492s
ok  	github.com/cpu-model/hammond/internal/engine	0.594s
ok  	github.com/cpu-model/hammond/internal/generator	0.379s
ok  	github.com/cpu-model/hammond/internal/instrument	0.004s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
