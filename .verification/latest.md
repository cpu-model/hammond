# Latest Go verification

- Source commit: 4c6ea597a102561b7be3b48e65c876f047b375bd
- Run: https://github.com/cpu-model/hammond/actions/runs/40230258254
- UTC: 2026-10-25T20:20:19Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	4.376s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.005s
ok  	github.com/cpu-model/hammond/internal/electrical	0.413s
ok  	github.com/cpu-model/hammond/internal/engine	0.481s
ok  	github.com/cpu-model/hammond/internal/generator	0.314s
ok  	github.com/cpu-model/hammond/internal/instrument	0.004s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
