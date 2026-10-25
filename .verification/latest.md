# Latest Go verification

- Source commit: cc96baf53ffce89d8d1cbceee8065005a31e5c9e
- Run: https://github.com/cpu-model/hammond/actions/runs/40220459327
- UTC: 2026-10-25T18:16:49Z

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
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/electrical	0.526s
ok  	github.com/cpu-model/hammond/internal/engine	0.510s
ok  	github.com/cpu-model/hammond/internal/generator	0.379s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.005s
```

## vet output
```text
```

## build output
```text
```
