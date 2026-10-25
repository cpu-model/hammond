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
