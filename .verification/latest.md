# Latest Go verification

- Source commit: e23c707822007b0ea645ba2169d3e7d4e895eba7
- Run: https://github.com/cpu-model/hammond/actions/runs/40220218979
- UTC: 2026-10-25T18:13:51Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.541s
ok  	github.com/cpu-model/hammond/internal/engine	0.490s
ok  	github.com/cpu-model/hammond/internal/generator	0.385s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.005s
```

## vet output
```text
```

## build output
```text
```
