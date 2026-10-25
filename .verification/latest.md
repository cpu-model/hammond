# Latest Go verification

- Source commit: 89a17f61ca2f6c339acf9a189df3a14864ee0a25
- Run: https://github.com/cpu-model/hammond/actions/runs/40228402495
- UTC: 2026-10-25T19:56:51Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	2.773s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.011s
ok  	github.com/cpu-model/hammond/internal/electrical	0.512s
ok  	github.com/cpu-model/hammond/internal/engine	0.636s
ok  	github.com/cpu-model/hammond/internal/generator	0.380s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
```

## vet output
```text
```

## build output
```text
```
