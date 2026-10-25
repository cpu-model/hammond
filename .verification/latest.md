# Latest Go verification

- Source commit: 618fe9ef01feda846e1fb89d5579952c53260730
- Run: https://github.com/cpu-model/hammond/actions/runs/40226131908
- UTC: 2026-10-25T19:27:55Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.324s
ok  	github.com/cpu-model/hammond/internal/engine	0.412s
ok  	github.com/cpu-model/hammond/internal/generator	0.299s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.005s
```

## vet output
```text
```

## build output
```text
```
