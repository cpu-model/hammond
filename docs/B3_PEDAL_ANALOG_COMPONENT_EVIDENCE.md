# B3/C3 pedal resistor panel and filter — component evidence

## Primary-source cross-check

B3/C3 schematic: https://manuals.plus/m/1764f5c997292f510b1451dbaa1e3c2ac48ee76b3628eff16a05a580795ece2b

B3/C3 Handy Wiring Guide: https://www.hammondx66.com/wp-content/uploads/2022/09/Handy_Wiring_Guide_B3_C3.pdf

The schematic labels **R80 and L20** in the pedal filter, and **R81–R86** in the resistor panel. These are separate physical assemblies on the back of the manual assembly. The guide identifies four signal wires (brown, red, orange, yellow) plus a black ground wire. Brown carries fundamental/third; red second/fourth; orange sixth/eighth; yellow tenth/twelfth.

### Printed schematic labels — NOT calibrated component values

| Designator | Adjacent printed number | Unit verified? | Connection verified? |
|---|---:|---|---|
| R80 | 10 | No | No |
| L20 | — | No inductance available | No |
| R81 | 10 | No | No |
| R82 | 47 | No | No |
| R83 | 470 | No | No |
| R84 | 5 | No | No |
| R85 | 5 | No | No |
| R86 | 20 | No | No |

**Do not convert these numbers to ohms, kilohms, or henries without checking the schematic legend and the original parts list.** The OCR output from some copies has a confusing '50/300' near R80; that annotation belongs to adjacent components in some text extractions and cannot be assigned reliably from OCR alone.

## Implementation boundary

- Four pedal bus voltages are available in the engine.
- `CompilePedalRLNetwork` can solve an explicitly wired resistive/inductive network once terminals and sourced values are known.
- It cannot derive the real network from component labels or color groups alone.
- A valid B3 pedal circuit still requires: exact terminals for R80–R86 and L20, units and values, ground path, resistor-panel outputs, and connection to pedal drawbars / T1.

No factory netlist or physical values have been asserted in code by this audit.
