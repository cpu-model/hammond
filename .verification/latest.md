# Latest Go verification

- Source commit: fdd99e01598800f2db99d879b6e4214a9bba3ffb
- Run: https://github.com/cpu-model/hammond/actions/runs/40226732812
- UTC: 2026-10-25T19:35:34Z

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
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.008s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/electrical	0.517s
ok  	github.com/cpu-model/hammond/internal/engine	0.595s
ok  	github.com/cpu-model/hammond/internal/generator	0.444s
ok  	github.com/cpu-model/hammond/internal/instrument	0.003s
ok  	github.com/cpu-model/hammond/internal/registration	0.006s
```

## vet output
```text
```

## build output
```text
```
