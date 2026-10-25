# Latest Go verification

- Source commit: ec958e07ef79e6bde1f2147a164888c54c5a8219
- Run: https://github.com/cpu-model/hammond/actions/runs/40223905966
- UTC: 2026-10-25T19:00:29Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.488s
ok  	github.com/cpu-model/hammond/internal/engine	0.489s
ok  	github.com/cpu-model/hammond/internal/generator	0.341s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.014s
```

## vet output
```text
```

## build output
```text
```
