# Latest Go verification

- Source commit: 3eea6495c75362f12b755295f0c40f0326a3008f
- Run: https://github.com/cpu-model/hammond/actions/runs/40228667012
- UTC: 2026-10-25T20:00:23Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	3.631s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.007s
ok  	github.com/cpu-model/hammond/internal/electrical	0.533s
ok  	github.com/cpu-model/hammond/internal/engine	0.658s
ok  	github.com/cpu-model/hammond/internal/generator	0.385s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
