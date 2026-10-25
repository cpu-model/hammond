# Latest Go verification

- Source commit: 18c07a0dab4bc34fdf3d7c3a29352bfab320ca16
- Run: https://github.com/cpu-model/hammond/actions/runs/40214049398
- UTC: 2026-10-25T16:56:34Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.006s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.357s
ok  	github.com/cpu-model/hammond/internal/generator	0.296s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
