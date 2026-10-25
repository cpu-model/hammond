# Latest Go verification

- Source commit: ffced6484c1f1caa6e235f1c95e636ddf98fdf66
- Run: https://github.com/cpu-model/hammond/actions/runs/40225882163
- UTC: 2026-10-25T19:24:47Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.467s
ok  	github.com/cpu-model/hammond/internal/engine	0.518s
ok  	github.com/cpu-model/hammond/internal/generator	0.382s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
