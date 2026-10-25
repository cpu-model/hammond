# Latest Go verification

- Source commit: 6262ffa0c0297d0f558db42cbcf982c2cdc05e23
- Run: https://github.com/cpu-model/hammond/actions/runs/40226617371
- UTC: 2026-10-25T19:34:03Z

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
?   	github.com/cpu-model/hammond/cmd/wavcompare	[no test files]
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/electrical	0.360s
ok  	github.com/cpu-model/hammond/internal/engine	0.448s
ok  	github.com/cpu-model/hammond/internal/generator	0.313s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
