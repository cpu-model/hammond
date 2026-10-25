# Latest Go verification

- Source commit: 8d155df2293d99bc464f27752089cf061d1a0c5e
- Run: https://github.com/cpu-model/hammond/actions/runs/40227716286
- UTC: 2026-10-25T19:48:09Z

| Check | Exit code |
| --- | ---: |
| go test ./... | 0 |
| go vet ./... | 0 |
| go build ./... | 0 |

## test output
```text
?   	github.com/cpu-model/hammond/cmd/contactdemo	[no test files]
?   	github.com/cpu-model/hammond/cmd/hammond	[no test files]
ok  	github.com/cpu-model/hammond/cmd/manualdemo	0.946s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.005s
ok  	github.com/cpu-model/hammond/internal/electrical	0.632s
ok  	github.com/cpu-model/hammond/internal/engine	0.817s
ok  	github.com/cpu-model/hammond/internal/generator	0.497s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
