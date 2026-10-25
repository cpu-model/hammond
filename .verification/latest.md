# Latest Go verification

- Source commit: 18449b3b8fd7ef99e4a3a1576f15fb07a81523e4
- Run: https://github.com/cpu-model/hammond/actions/runs/40228568951
- UTC: 2026-10-25T19:59:12Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	4.532s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.013s
ok  	github.com/cpu-model/hammond/internal/electrical	0.627s
ok  	github.com/cpu-model/hammond/internal/engine	0.856s
ok  	github.com/cpu-model/hammond/internal/generator	0.485s
ok  	github.com/cpu-model/hammond/internal/instrument	0.009s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
