# Latest Go verification

- Source commit: ef3bda9396a1117d41cb98e7845034b6592d167d
- Run: https://github.com/cpu-model/hammond/actions/runs/40222757379
- UTC: 2026-10-25T18:45:46Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.496s
ok  	github.com/cpu-model/hammond/internal/engine	0.477s
ok  	github.com/cpu-model/hammond/internal/generator	0.376s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.006s
```

## vet output
```text
```

## build output
```text
```
