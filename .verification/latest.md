# Latest Go verification

- Source commit: 585d9b296330486ff25ac61154e1da0666c3dafb
- Run: https://github.com/cpu-model/hammond/actions/runs/40222194884
- UTC: 2026-10-25T18:38:40Z

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
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
--- FAIL: TestConnectPedalPanelInputsExplicitMapping (0.00s)
    b3_pedal_panel_bus_mapping_test.go:28: same-bus fanout rejected: pedal bus 1 assigned to multiple panel rows
FAIL
FAIL	github.com/cpu-model/hammond/internal/electrical	0.558s
ok  	github.com/cpu-model/hammond/internal/engine	0.507s
ok  	github.com/cpu-model/hammond/internal/generator	0.387s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.008s
FAIL
```

## vet output
```text
```

## build output
```text
```
