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

## Corrected primary-source cross-check — 2026-10-25

A fresh direct inspection of the **actual PDF page image** (page index 71) and its embedded text extraction establishes that **L20 is printed in the pedal-switch area**. This supersedes the preceding tentative statement that L20 could not be confirmed. The printed sequence in the PDF's schematic text is `L20 R83 470`; it does **not** establish L20's inductance, circuit role, or endpoints. The label is therefore **confirmed**, but the electrical component model remains **unverified**.

Source: https://www.lost-manuals.com/pdf/hammond--b-3--service--ID201.pdf#page=72

The page image shows a **resistor ladder** above the pedal contacts. The resistors labeled R81/R84, R82/R85, and R83/R86 occupy three visible horizontal rows, while R80 appears on a left-hand vertical branch. The drawing is not equivalent to four independent pedal RL filters. The exact conductive junctions, source terminals, and output loading still require detailed tracing at higher resolution before entering a production netlist.

**Additional directly visible circuit fact:** the pedal drawbars are drawn adjacent to T1 at the top left, while the resistor ladder is at the bottom left. Their apparent proximity on the same sheet does not by itself establish a complete end-to-end circuit path.

The confirmed schematic label `L20` is not permission to assign it an arbitrary inductance in henries. No value for it has been established.

## Page-image trace: resistor-network geometry (2026-10-25)

**Source actually inspected in this increment:** page 72 image (PDF index 71) at
https://www.lost-manuals.com/pdf/hammond--b-3--service--ID201.pdf#page=72
The following observations refer to the **bottom-left PEDAL SWITCH inset**,
not to the nearby mains motor or vibrato delay-line circuits.

| Visible row (top to bottom) | Left resistor | Right resistor | Printed numbers |
| --- | --- | --- | --- |
| Top | R81 | R84 | 10, 5 |
| Middle | R82 | R85 | 47, 5 |
| Bottom | R83 | R86 | 470, 20 |

R80 (printed 10) is drawn vertically on the left, between the pedal
switch contacts and the three-row network. The schematic shows a common
right-hand vertical rail at the ends of R84, R85 and R86. The row
ordering and shared right rail are **visual observations**; they do not
establish the signal source, ground reference, total impedance, or the
units of the printed numbers. The network is **not** represented by
three independent series R-L branches in the current emulator.

The label L20 appears immediately adjacent to the vertical R80 branch
in the drawing. Its **component type, inductance and terminals have not
been established** from the visible resolution; it is not safe to infer
that this is a discrete inductor simply from the letter L.

### Trace checkpoints still required before compilation

1. Resolve the left-side junctions of R81/R82/R83, including whether
   they are a common conductor or three distinct pedal-bus taps.
2. Resolve the right rail's destination (including any drawbar contact
   and load), not just its visible continuity.
3. Resolve R80's two terminals, the meaning of L20, and their junctions
   with the eight pedal contacts.
4. Verify units from the schematic legend, parts list or independent
   service evidence before entering any ohmic value.

No `PedalCircuitTopology` edge or `PedalComponentValue` is promoted
to schematic provenance by these observations alone.

## Independent wiring-diagram cross-check (2026-10-25)

A **second drawing**, distinct from the electronic schematic, is included
in the same service manual: **WIRING DIAGRAM - HAMMOND B3 OR C3**,
revision B, dated December 12, 1995. Source:
https://synthmania.com/wp-content/uploads/2020/05/Hammond-Organ-B3-C3-Service-Manual.pdf

The wiring diagram explicitly identifies **two separate physical assemblies**
on the back of the manual assembly:

| Assembly label on wiring diagram | Designators visibly assigned |
| --- | --- |
| PEDAL FILTER | R80, L20 |
| RESISTOR PANEL | R81, R82, R83, R84, R85, R86 |

This is an important distinction: the schematic's visually adjacent symbols
do **not** mean all seven resistors and L20 are physically mounted together.
It also strengthens identification of **L20 as part of the pedal filter**,
though it does **not** supply an inductance in henries, winding resistance,
or verified electrical terminal connections.

**Modeling consequence:** represent pedal filter and resistor panel as
separate source-traceable assembly boundaries when tracing wiring. Do not
equate assembly placement with circuit connectivity; a wire-color label or
a component's location in the diagram is not by itself a verified netlist
edge. The earlier notes on a common right-hand rail remain *schematic
observations*, not verified wiring-diagram endpoints.

**Next source requirement:** reconcile individual color-coded wires between
the resistor panel, pedal filter, pedal switch, drawbar base and T1 using
the wiring diagram's printed color legend, then cross-check each resulting
electrical net against the schematic before encoding it.
