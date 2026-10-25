# Latest Go verification

- Source commit: dfc1a3331f7c48fb34a9dcc0f62c779df790b51f
- Run: https://github.com/cpu-model/hammond/actions/runs/40230621951
- UTC: 2026-10-25T20:25:00Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	6.413s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.004s
ok  	github.com/cpu-model/hammond/internal/electrical	0.649s
ok  	github.com/cpu-model/hammond/internal/engine	0.829s
ok  	github.com/cpu-model/hammond/internal/generator	0.497s
ok  	github.com/cpu-model/hammond/internal/instrument	0.009s
ok  	github.com/cpu-model/hammond/internal/registration	0.006s
```

## vet output
```text
```

## build output
```text
```
