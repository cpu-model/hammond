# Latest Go verification

- Source commit: 42befbc119461e817e8a921a1754244c4f623676
- Run: https://github.com/cpu-model/hammond/actions/runs/40219832949
- UTC: 2026-10-25T18:09:03Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.588s
ok  	github.com/cpu-model/hammond/internal/engine	0.556s
ok  	github.com/cpu-model/hammond/internal/generator	0.365s
ok  	github.com/cpu-model/hammond/internal/instrument	0.005s
ok  	github.com/cpu-model/hammond/internal/registration	0.005s
```

## vet output
```text
```

## build output
```text
```
