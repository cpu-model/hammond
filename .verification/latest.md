# Latest Go verification

- Source commit: b9c6e1e3d9ace86ddc3e9dcff61b450ae8161fd4
- Run: https://github.com/cpu-model/hammond/actions/runs/40208444067
- UTC: 2026-10-25T15:45:44Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.012s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.339s
ok  	github.com/cpu-model/hammond/internal/generator	0.264s
ok  	github.com/cpu-model/hammond/internal/instrument	0.006s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
