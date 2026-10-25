# Latest Go verification

- Source commit: e4548910cde68367ce23ba1f56f91093c7818109
- Run: https://github.com/cpu-model/hammond/actions/runs/40231009881
- UTC: 2026-10-25T20:29:54Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	5.035s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.004s
ok  	github.com/cpu-model/hammond/internal/electrical	0.571s
ok  	github.com/cpu-model/hammond/internal/engine	0.637s
ok  	github.com/cpu-model/hammond/internal/generator	0.415s
ok  	github.com/cpu-model/hammond/internal/instrument	0.005s
ok  	github.com/cpu-model/hammond/internal/registration	0.168s
```

## vet output
```text
```

## build output
```text
```
