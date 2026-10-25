# Latest Go verification

- Source commit: c8da0c15dccdc086d282e219fa10c3f60c35e6ec
- Run: https://github.com/cpu-model/hammond/actions/runs/40210524146
- UTC: 2026-10-25T16:12:15Z

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
ok  	github.com/cpu-model/hammond/internal/electrical	0.013s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.315s
ok  	github.com/cpu-model/hammond/internal/generator	0.227s
ok  	github.com/cpu-model/hammond/internal/instrument	0.007s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
```

## vet output
```text
```

## build output
```text
```
