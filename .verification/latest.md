# Latest Go verification

- Source commit: 2b3846045ce0c87f6ff1c15a71866cd9ca23dbb6
- Run: https://github.com/cpu-model/hammond/actions/runs/40219932933
- UTC: 2026-10-25T18:10:16Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.439s
ok  	github.com/cpu-model/hammond/internal/engine	0.390s
ok  	github.com/cpu-model/hammond/internal/generator	0.327s
ok  	github.com/cpu-model/hammond/internal/instrument	0.004s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
