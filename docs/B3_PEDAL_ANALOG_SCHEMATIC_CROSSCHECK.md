# Pedal analog schematic: cross-check against independent copies

Primary B3/C3 circuit scan:
https://manuals.plus/m/1764f5c997292f510b1451dbaa1e3c2ac48ee76b3628eff16a05a580795ece2b

Independent 1995 wiring drawing:
https://schematicsforfree.com/files/Audio/Products/Musician/H-I/Hammond%20Organ/B3.pdf

The first document explicitly lists R81=10, R84=5, R82=47, R85=5, R83=470, R86=20, and R80=10, adjacent to the pedal switch schematic. These printed numbers must NOT be treated as ohms until the original schematic's resistance convention is established. L20 is labelled but its inductance is not provided by the accessible text.

The independent wiring drawing confirms physical separation between a PEDAL FILTER (R80/L20) and a RESISTOR PANEL (R81-R86), mounted on the back of the manual assembly. It labels brown, red, orange and yellow signal wires, and black ground, but its extracted text is not sufficient to reconstruct every internal node.

The resistor panel has three drawn rows, with pairs R81/R84, R82/R85, R83/R86. This is a *layout observation*, not proof of series, parallel, or common-node connections. Do not create a factory netlist by treating schematic drawing proximity as electrical connectivity.

Before using CompilePedalRLNetwork for a B3 preset:
- Obtain a readable resistor-network drawing and identify all junction dots and wire crossings.
- Establish component units and L20 inductance from a parts list or measured data.
- Map the four colored busbar wires to the two pedal drawbar inputs and T1.
- Preserve the resistor panel and pedal filter as distinct physical subassemblies.

No component values or terminal connectivity were added to the runtime from this cross-check.
