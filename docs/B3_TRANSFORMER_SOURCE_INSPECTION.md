# B3/C3 transformer source inspection — initial findings

Primary diagram inspected: **B3/C3 SCHEMATIC**, revision 1.0,
dated December 11, 1995, available at:
https://www.lost-manuals.com/pdf/hammond--b-3--service--ID201.pdf

Corroborating B3/C3 wiring diagram:
https://www.lost-manuals.com/pdf/hammond--c-3--service--ID201.pdf

## Directly identified topology

- T1 is explicitly labelled **MATCHING TRANSFORMER (LOWER AND PEDAL)**.
- T2 is explicitly labelled **MATCHING TRANSFORMER UPPER**.
- The diagram distinguishes **UPPER MANUAL BUSBARS** from
  **LOWER MANUAL BUSBARS**, with drawbar and preset circuitry.
- The wiring diagram separately labels T1 and T2.

## Still unverified

- Winding terminal-to-terminal connectivity and exact tap sequence.
- Actual turns ratios for positions 1–8.
- Winding resistance, leakage/magnetizing inductance and load.
- Frequency response and source impedance at the matching stage.
- Whether all nine conductors can be treated as independent loaded
  nodes (the current experimental model assumes so).

**Do not** populate `TransformerEvidence` with generic URLs as a
substitute for verified, specific topology and measured parameters.
The diagram establishes transformer identity, not the numerical
circuit transfer function.

## Next source-analysis task

Inspect the diagram at full resolution, transcribe T1 and T2
terminal labels and winding/tap connections separately, record
diagram coordinates or component references, and cross-check the
wiring drawing. Then define a graph of electrical nodes and edges
before choosing a numerical transformer approximation.
