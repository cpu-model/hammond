# Latest Go verification

- Source commit: 1a65a1f160d8675df26084253530508ffdf68d95
- Run: https://github.com/cpu-model/hammond/actions/runs/40207834583
- UTC: 2026-10-25T15:37:59Z

| Check | Exit code |
| --- | ---: |
| go test ./... | 1 |
| go vet ./... | 1 |
| go build ./... | 1 |

## test output
```text
# github.com/cpu-model/hammond/internal/engine
internal/engine/electrical_bus.go:129:88: syntax error: mixed named and unnamed parameters
FAIL	github.com/cpu-model/hammond/cmd/contactdemo [build failed]
?   	github.com/cpu-model/hammond/cmd/hammond	[no test files]
FAIL	github.com/cpu-model/hammond/cmd/manualdemo [build failed]
FAIL	github.com/cpu-model/hammond/internal/engine [build failed]
ok  	github.com/cpu-model/hammond/internal/electrical	0.002s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/generator	0.253s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
FAIL
```

## vet output
```text
# github.com/cpu-model/hammond/internal/engine
internal/engine/electrical_bus.go:129:88: syntax error: mixed named and unnamed parameters
# github.com/cpu-model/hammond/internal/engine
# [github.com/cpu-model/hammond/internal/engine]
vet: internal/engine/electrical_bus.go:129:88: mixed named and unnamed parameters
```

## build output
```text
# github.com/cpu-model/hammond/internal/engine
internal/engine/electrical_bus.go:129:88: syntax error: mixed named and unnamed parameters
```
