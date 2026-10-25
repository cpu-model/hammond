# Latest Go verification

- Source commit: 0783a330275f79221570e01401b4d0a53c2b9669
- Run: https://github.com/cpu-model/hammond/actions/runs/40229000263
- UTC: 2026-10-25T20:04:28Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	4.606s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.009s
ok  	github.com/cpu-model/hammond/internal/electrical	0.517s
ok  	github.com/cpu-model/hammond/internal/engine	0.652s
ok  	github.com/cpu-model/hammond/internal/generator	0.378s
ok  	github.com/cpu-model/hammond/internal/instrument	0.003s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
