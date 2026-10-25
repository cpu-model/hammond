# B3 pedal analog values: conflicting source labels

Reviewed 2026-10-25.

Two accessible copies of the 1995 B3/C3 schematic present ambiguous extracted text around the pedal filter and resistor panel:
- https://device.report/m/1764f5c997292f510b1451dbaa1e3c2ac48ee76b3628eff16a05a580795ece2b
- https://www.scribd.com/doc/36619802/Hammond-B3-Service-Manual

Both show R81 10, R84 5, R82 47, R85 5, R83 470 and R86 20 in the pedal area. However the extracted text near R80 includes both 10 and 50/300, with overlapping nearby schematic labels. The latter may be unrelated and cannot be assigned to R80 by text extraction alone.

**Critical designator collision:** A different B-3 service PDF's indexed text assigns R80 through R86 to the vibrato phase-shift line (18MEG resistors), not to the pedal circuit:
https://www.lost-manuals.com/pdf/hammond--b-3--service--ID201.pdf

Consequently, designators such as R80 are NOT globally unique across all Hammond service documents and drawing revisions. Never import an R80 value from another drawing without verifying the circuit identity and schematic revision.

Implementation rule: no default analog values or factory pedal netlist are authorized by these excerpts. The existing evidence-bearing pedal RL compiler must continue to require explicit topology and component values. A readable schematic with terminal junctions, units, and revision is required before authenticating the circuit.
