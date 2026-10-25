# Latest Go verification

- Source commit: 33d2db80597f71bb227b28e1b80488900412aabe
- Run: https://github.com/cpu-model/hammond/actions/runs/40221697012
- UTC: 2026-10-25T18:32:20Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.573s
ok  	github.com/cpu-model/hammond/internal/engine	0.476s
ok  	github.com/cpu-model/hammond/internal/generator	0.415s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.005s
```

## vet output
```text
```

## build output
```text
```
