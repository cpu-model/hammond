# T1 and T2 transformer connection records

`ManualTransformerPair` enforces the separately identified manual
transformer paths: upper manual to T2, lower manual to T1. The
`ManualTransformerConnection` retains each of nine raw bus voltages
along with the drawbar's selected tap (1..8), or marks it disconnected
for position 0.

This is **wiring metadata**, not a transformer electrical model.
It deliberately contains no turns ratios, tap impedances, winding
resistances, leakage inductances, magnetic saturation, load impedance
or voltage transfer. Those require traceable circuit evidence.

The pair does not combine manual signals. Tests cover transformer
identity, selected taps, isolation and allocation-free traversal.
