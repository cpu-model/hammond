# B3/C3 matching-transformer topology — visual inspection

Source: Hammond B3/C3 service PDF, **PDF page 72 (zero-based
page 71)**, sheet labelled **B3/C3 SCHEMATIC**, dated
December 11, 1995:
https://www.lost-manuals.com/pdf/hammond--b-3--service--ID201.pdf

## Visible, source-grounded connections

- T1 is labelled **MATCHING TRANSFORMER (LOWER AND PEDAL)**.
  It is not exclusively a lower-manual transformer.
- T2 is labelled **MATCHING TRANSFORMER UPPER**.
- Separate upper and lower manual drawbar groups are drawn with
  multiple individual connections toward the transformer windings.
- Pedal drawbars and lower manual drawbars appear in the T1 region.
- Upper and lower manual busbars and their respective preset panels
  are separately drawn.
- Transformer outputs are routed onward through vibrato selection.

## Model consequences

1. The current `ManualTransformerPair` describes only the
   **manual contribution** to T1 and T2; it does not represent
   the complete T1 circuit when pedals are active.
2. Do not assert that all nine drawbar settings are simply
   independent voltage gains. Tap position selects physical
   connections on a winding.
3. A circuit graph must account for T1's pedal connections,
   manual presets, and downstream loading.
4. The current resistive bus model assumes independent ground
   loads. That is an experimental simplification, not a
   transcription of the illustrated transformer network.

## Not established by this inspection

The raster schematic does not supply trustworthy numeric turns
ratios, impedance values, or frequency-dependent transfer
characteristics. Precise terminal-level netlist transcription
remains pending. Do not mark `TransformerEvidence` complete
on the basis of this visual identification alone.
