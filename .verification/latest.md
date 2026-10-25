# Latest Go verification

- Source commit: 62770de9d0c3702617e4159cc504b581d90dcd81
- Run: https://github.com/cpu-model/hammond/actions/runs/40231084607
- UTC: 2026-10-25T20:30:54Z

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
ok  	github.com/cpu-model/hammond/cmd/manualdemo	5.411s
ok  	github.com/cpu-model/hammond/cmd/wavcompare	0.007s
ok  	github.com/cpu-model/hammond/internal/electrical	0.490s
ok  	github.com/cpu-model/hammond/internal/engine	0.623s
ok  	github.com/cpu-model/hammond/internal/generator	0.383s
ok  	github.com/cpu-model/hammond/internal/instrument	0.007s
ok  	github.com/cpu-model/hammond/internal/registration	0.167s
```

## vet output
```text
```

## build output
```text
```
