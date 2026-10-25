# Latest Go verification

- Source commit: 6eb515e8b7f4d07708ea58c81779c92b835539de
- Run: https://github.com/cpu-model/hammond/actions/runs/40228912653
- UTC: 2026-10-25T20:03:22Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	4.622s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.009s
ok  	github.com/cpu-model/hammond/internal/electrical	0.564s
ok  	github.com/cpu-model/hammond/internal/engine	0.618s
ok  	github.com/cpu-model/hammond/internal/generator	0.401s
ok  	github.com/cpu-model/hammond/internal/instrument	0.001s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
