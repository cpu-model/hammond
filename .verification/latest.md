# Latest Go verification

- Source commit: 34f06492c10b701a5e49127d736fb2da2f2eebd7
- Run: https://github.com/cpu-model/hammond/actions/runs/40221800656
- UTC: 2026-10-25T18:33:33Z

| Check | Exit code |
| --- | ---: |
| go test ./... | 1 |
| go vet ./... | 0 |
| go build ./... | 0 |

## test output
```text
?   	github.com/cpu-model/hammond/cmd/contactdemo	[no test files]
?   	github.com/cpu-model/hammond/cmd/hammond	[no test files]
?   	github.com/cpu-model/hammond/cmd/manualdemo	[no test files]
--- FAIL: TestConnectPedalPanelInputsExplicitMapping (0.00s)
    b3_pedal_panel_bus_mapping_test.go:28: same-bus fanout rejected: pedal bus 1 assigned to multiple panel rows
FAIL
FAIL	github.com/cpu-model/hammond/internal/electrical	0.179s
ok  	github.com/cpu-model/hammond/internal/engine	0.292s
ok  	github.com/cpu-model/hammond/internal/generator	0.241s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/registration	0.003s
FAIL
```

## vet output
```text
```

## build output
```text
```
