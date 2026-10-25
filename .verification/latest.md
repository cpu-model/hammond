# Latest Go verification

- Source commit: 72bc1d89ac17187fc6307132aed6d98353ecf1fb
- Run: https://github.com/cpu-model/hammond/actions/runs/40206324789
- UTC: 2026-10-25T15:19:02Z

| Check | Exit code |
| --- | ---: |
| go test ./... | 1 |
| go vet ./... | 1 |
| go build ./... | 0 |

## test output
```text
# github.com/cpu-model/hammond/internal/electrical [github.com/cpu-model/hammond/internal/electrical.test]
internal/electrical/resistive_bus_test.go:8:6: experimentalOhms redeclared in this block
	internal/electrical/contact_network_test.go:8:6: other declaration of experimentalOhms
?   	github.com/cpu-model/hammond/cmd/contactdemo	[no test files]
?   	github.com/cpu-model/hammond/cmd/hammond	[no test files]
?   	github.com/cpu-model/hammond/cmd/manualdemo	[no test files]
FAIL	github.com/cpu-model/hammond/internal/electrical [build failed]
?   	github.com/cpu-model/hammond/internal/wav	[no test files]
ok  	github.com/cpu-model/hammond/internal/engine	0.388s
ok  	github.com/cpu-model/hammond/internal/generator	0.283s
ok  	github.com/cpu-model/hammond/internal/instrument	0.002s
ok  	github.com/cpu-model/hammond/internal/registration	0.004s
FAIL
```

## vet output
```text
# github.com/cpu-model/hammond/internal/electrical
# [github.com/cpu-model/hammond/internal/electrical]
vet: internal/electrical/resistive_bus_test.go:8:6: experimentalOhms redeclared in this block
```

## build output
```text
```
