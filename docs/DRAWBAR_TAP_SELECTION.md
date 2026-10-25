# Drawbar selection without unverified electrical transfer

`electrical.DrawbarTapSelection` records the nine drawbar positions
(0..8) and validates their range.

`SelectedBusVoltages` is a **diagnostic mask**: position 0 suppresses
the corresponding bus from its returned array, while any position
1..8 retains the raw bus voltage. It neither modifies the electrical
bus loading nor claims to model transformer tap ratios.

In a real B3, moving a drawbar changes connections to transformer
taps. Electrical transfer and impedance can depend on the circuit
and on simultaneous drawbar settings. A numeric 3 dB-per-step
gain or a masked bus voltage is not an adequate substitute for
that network.

This stage deliberately avoids summing nine conductors or collapsing
upper and lower manuals. Future implementation requires documented
tap wiring and impedances, followed by an explicit circuit solver.
