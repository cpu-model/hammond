# Latest Go verification

- Source commit: 9a67959918508352f6c7f65610473245db8120d1
- Run: https://github.com/cpu-model/hammond/actions/runs/40207876275
- UTC: 2026-10-25T15:38:33Z

| Check | Exit code |
| --- | ---: |
| go test ./... | 1 |
| go vet ./... | 1 |
| go build ./... | 1 |

## test output
```text
# github.com/cpu-model/hammond/internal/engine
internal/engine/electrical_bus.go:133:5: no new variables on left side of :=
FAIL	github.com/cpu-model/hammond/cmd/contactdemo [build failed]
?   	github.com/cpu-model/hammond/cmd/hammond	[no test files]
FAIL	github.com/cpu-model/hammond/cmd/manualdemo [build failed]
FAIL	github.com/cpu-model/hammond/internal/engine [build failed]
ok  	github.com/cpu-model/hammond/internal/electrical	0.004s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/generator	0.291s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
FAIL
```

## vet output
```text
# github.com/cpu-model/hammond/internal/engine
internal/engine/electrical_bus.go:133:5: no new variables on left side of :=
# github.com/cpu-model/hammond/internal/engine
# [github.com/cpu-model/hammond/internal/engine]
vet: internal/engine/electrical_bus.go:133:5: no new variables on left side of :=
```

## build output
```text
# github.com/cpu-model/hammond/internal/engine
internal/engine/electrical_bus.go:133:5: no new variables on left side of :=
```
