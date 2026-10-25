# Latest Go verification

- Source commit: 488665e04584942230024de5ed288bd386192600
- Run: https://github.com/cpu-model/hammond/actions/runs/40230917421
- UTC: 2026-10-25T20:28:47Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	6.528s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.011s
ok  	github.com/cpu-model/hammond/internal/electrical	0.626s
ok  	github.com/cpu-model/hammond/internal/engine	0.821s
ok  	github.com/cpu-model/hammond/internal/generator	0.501s
ok  	github.com/cpu-model/hammond/internal/instrument	0.005s
ok  	github.com/cpu-model/hammond/internal/registration	0.219s
```

## vet output
```text
```

## build output
```text
```
