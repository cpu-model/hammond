# B3/C3 schematic transcription — pedal and manual boundaries

**Primary source:** Hammond B3/C3 service manual, *B3/C3 SCHEMATIC*,
dated December 11, 1995, PDF page 72 (page index 71):
https://www.lost-manuals.com/pdf/hammond--b-3--service--ID201.pdf

This is a **label-level transcription**. The extracted text does not
preserve conductor crossings and junction dots reliably. No
terminal-to-terminal edge is asserted from text proximity alone.

| Printed label | Location / significance | Confidence |
| --- | --- | --- |
| T1 — MATCHING TRANSFORMER (LOWER AND PEDAL) | Lower manual and pedal region | Explicit printed label |
| T2 — MATCHING TRANSFORMER UPPER | Upper manual region | Explicit printed label |
| PEDAL DRAWBARS | Adjacent to matching transformer region | Explicit printed label |
| 8 CONTACTS UNDER TYPICAL PEDAL KEY | Pedal switch wiring region | Explicit printed label |
| PEDAL SWITCH BUSBARS | Pedal switch wiring region | Explicit printed label |
| TERMINALS TO GENERATOR (SEE PEDAL WIRING CHART) | Pedal generator wiring | Explicit printed label |
| 9 CONTACTS UNDER TYPICAL PLAYING KEY | Manual key contact wiring | Explicit printed label |
| UPPER MANUAL BUSBARS | Upper manual wiring | Explicit printed label |
| LOWER MANUAL BUSBARS | Lower manual wiring | Explicit printed label |
| UPPER MANUAL PRESET PANEL | Upper preset wiring | Explicit printed label |
| LOWER MANUAL PRESET PANEL | Lower preset wiring | Explicit printed label |
| GENERATOR OUTPUT TERMINALS (GEN 1–43, GEN 44–48, GEN 49–91) | Generator outputs divided into three printed ranges | Explicit printed labels |

## Consequences for implementation

1. A pedal key has **eight** depicted contacts, unlike the
   nine-contact typical manual key. Do not reuse the manual's
   nine-contact representation for pedal keys.
2. The existing `PedalTransformerInputs` is an *external source
   interface*, not a pedal key-contact or pedal drawbar simulation.
3. `TerminalGraph` must remain empty of asserted B3 edges until
   conductor junctions and terminal identities are read visually.
4. The generator output ranges are printed labels, not evidence
   that the ranges correspond to three independent generators.

## Required next inspection

Inspect the pedal key group and its generator wiring chart at
sufficient resolution to trace eight contacts per pedal key.
Then inspect T1 winding terminals and the two pedal drawbars,
marking actual junction dots and crossover wires. Record
page/diagram-region references on every graph edge.
