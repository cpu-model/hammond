# B3/C3 pedal resistor and reactor transcription — evidence ledger

Primary schematic: https://synthmania.com/wp-content/uploads/2020/05/Hammond-Organ-B3-C3-Service-Manual.pdf
Diagram title: B3/C3 SCHEMATIC (1995 reproduction).

## Legible printed labels in pedal area

| Designator | Adjacent printed number | Electrical unit independently verified? |
| --- | --- | --- |
| R80 | 10 | No |
| R81 | 10 | No |
| R82 | 47 | No |
| R83 | 470 | No |
| R84 | 5 | No |
| R85 | 5 | No |
| R86 | 20 | No |
| L20 | none confirmed | No |

These are *transcriptions of visible number strings*, not validated
ohmic/henry component specifications. The diagram also shows
eight contacts under a typical pedal key and a group of pedal
busbars. Neither visual proximity nor OCR reading order is a
sufficient substitute for tracing junction dots and conductive
lines at high resolution.

The currently implemented independent RL channels remain synthetic.
Do not enter these bare numbers into `BusConfig`, `PedalRLConfig`
or `PedalCircuitEvidence` as schematic-derived physical values
until units, connections and component roles are checked.

## Next transcription requirements

1. Identify each of four pedal busbar terminals on the schematic.
2. Trace each R80–R86 endpoint and L20 winding endpoint, preserving
   junctions and distinguishing crossing wires.
3. Identify the two pedal drawbar connections and T1 input boundary.
4. Confirm units, winding properties and revision applicability.
5. Record source page/region for every edge and parameter.

## Visual source inspection — 2026-10-25

**Inspected directly:** [B3/C3 service manual, schematic page 72 (PDF index 71)](https://www.lost-manuals.com/pdf/hammond--b-3--service--ID201.pdf), lower-left **PEDAL SWITCH** region, visually checked against the full-page rendering.

The schematic shows a distinct pedal-switch region with eight contacts under a typical pedal key, a labeled group of pedal busbars, generator terminals to the left, and a resistor network above the contacts. The network visibly contains the designators **R80–R86** and printed adjacent numbers 10, 10, 47, 470, 5, 5, 20. This confirms label presence and approximate region, **not** physical resistance units or junction connectivity.

**Important correction to earlier transcription:** the printed designator **L20** has *not* been independently confirmed in this pedal-switch region by this visual inspection. It must not be instantiated as an inductor in the emulator without a separate verified schematic or parts-list reference. The previous ledger row is an unverified candidate, not established circuit evidence.

The full schematic also places the two **PEDAL DRAWBARS** adjacent to the **T1 MATCHING TRANSFORMER (LOWER AND PEDAL)** at the upper-left, spatially separate from the pedal-switch drawing. The exact path between these regions has **not** yet been traced terminal by terminal.

### Evidence status and transcription discipline

- **Verified from visible drawing:** eight typical pedal contacts; pedal busbar label; R80–R86 designators and their neighboring number strings; T1 and pedal drawbar labels.
- **Not verified:** resistor units and tolerances, whether printed numbers are ohms or another notation, exact resistor endpoints and crossover junctions, reactor L20's identity/existence in this region, drawbar-to-T1 electrical edges.
- **Implementation consequence:** no new B3-sourced `PedalComponentValue` or `TerminalEdge` should be added yet. The existing synthetic RL tests remain synthetic.

The diagram itself is a 1995 reproduction of the B3/C3 schematic; a reproduction date is not evidence that every circuit revision uses identical components.
