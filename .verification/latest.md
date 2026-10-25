# Latest Go verification

- Source commit: c401e85fa7323609e292cefbc4f8ae191f7fb0d7
- Run: https://github.com/cpu-model/hammond/actions/runs/40206475966
- UTC: 2026-10-25T15:20:56Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.005s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.381s
ok  	github.com/cpu-model/hammond/internal/generator	0.307s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
