# Latest Go verification

- Source commit: cf1bfc9a19745f9f6f6b1e566e6c6c5ec31a586a
- Run: https://github.com/cpu-model/hammond/actions/runs/40227061810
- UTC: 2026-10-25T19:39:44Z

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
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.004s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/electrical	0.346s
ok  	github.com/cpu-model/hammond/internal/engine	0.441s
ok  	github.com/cpu-model/hammond/internal/generator	0.315s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
