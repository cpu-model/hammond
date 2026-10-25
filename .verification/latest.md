# Latest Go verification

- Source commit: abf9c983f6f417027f2484f805a58a245674d820
- Run: https://github.com/cpu-model/hammond/actions/runs/40213285094
- UTC: 2026-10-25T16:46:46Z

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
ok  	github.com/cpu-model/hammond/internal/engine	0.336s
ok  	github.com/cpu-model/hammond/internal/generator	0.272s
ok  	github.com/cpu-model/hammond/internal/instrument	0.008s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
