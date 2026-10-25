# Latest Go verification

- Source commit: 399ded7bc6589ab62cd88fd63f0cc32e86697264
- Run: https://github.com/cpu-model/hammond/actions/runs/40222367180
- UTC: 2026-10-25T18:40:50Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.466s
ok  	github.com/cpu-model/hammond/internal/engine	0.443s
ok  	github.com/cpu-model/hammond/internal/generator	0.300s
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
