# Latest Go verification

- Source commit: 3170adba8630794bb2caf3b8ed1af2fa86b6f834
- Run: https://github.com/cpu-model/hammond/actions/runs/40230140277
- UTC: 2026-10-25T20:18:51Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	5.574s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.013s
ok  	github.com/cpu-model/hammond/internal/electrical	0.627s
ok  	github.com/cpu-model/hammond/internal/engine	0.823s
ok  	github.com/cpu-model/hammond/internal/generator	0.502s
ok  	github.com/cpu-model/hammond/internal/instrument	0.003s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
