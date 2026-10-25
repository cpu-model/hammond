# Resistive-node analysis primitive

`electrical.ResistiveNode` solves the voltage at one ideal resistive
node using conductance-weighted source voltages. For source voltages
`V_i`, series resistances `R_i`, and ground loads `L_j`:

`V_node = sum(V_i/R_i) / (sum(1/R_i) + sum(1/L_j))`

This implements Kirchhoff's current law for a memoryless linear node.
The code validates source voltages and resistance values and detects
floating nodes. Tests cover voltage division, unequal source impedances,
invalid inputs and allocation-free evaluation.

**Authenticity boundary:** This primitive has not been connected to
the B3 engine. The Hammond resistance-wire values, transformer loading
and capacitive coupling are not established. This circuit equation
cannot yet be presented as a complete key-bus or key-click model.
