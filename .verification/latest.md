# Latest Go verification

- Source commit: cd1cf18c25610d67eee27bea4ffc0d3c5f0f542a
- Run: https://github.com/cpu-model/hammond/actions/runs/40226053001
- UTC: 2026-10-25T19:26:56Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.347s
ok  	github.com/cpu-model/hammond/internal/engine	0.401s
ok  	github.com/cpu-model/hammond/internal/generator	0.334s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.008s
```

## vet output
```text
```

## build output
```text
```
