# Latest Go verification

- Source commit: 580e909a1288c4ad6c5750a629d8aef1c419c39d
- Run: https://github.com/cpu-model/hammond/actions/runs/40221482580
- UTC: 2026-10-25T18:29:46Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.449s
ok  	github.com/cpu-model/hammond/internal/engine	0.461s
ok  	github.com/cpu-model/hammond/internal/generator	0.352s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.006s
```

## vet output
```text
```

## build output
```text
```
