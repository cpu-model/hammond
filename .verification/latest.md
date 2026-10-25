# Latest Go verification

- Source commit: 95f4041bd3fb4aafd044f0c293c11bf928ccac9b
- Run: https://github.com/cpu-model/hammond/actions/runs/40222885419
- UTC: 2026-10-25T18:47:19Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.335s
ok  	github.com/cpu-model/hammond/internal/engine	0.313s
ok  	github.com/cpu-model/hammond/internal/generator	0.250s
ok  	github.com/cpu-model/hammond/internal/instrument	0.003s
ok  	github.com/cpu-model/hammond/internal/registration	0.002s
```

## vet output
```text
```

## build output
```text
```
