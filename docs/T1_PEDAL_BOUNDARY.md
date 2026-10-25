# T1 pedal contribution interface

The B3/C3 schematic labels T1 as the matching transformer for
**lower manual and pedals**. The current electrical engine computes
only manual bus voltages; it does not synthesize pedal tones.

`PedalTransformerInputs` is a separate source record containing a
connection flag, source voltage and a provenance-tagged positive
source resistance in ohms. `NewT1InputBoundary` combines that
record with lower-manual T1 tap *metadata* without electrically
mixing them.

A disconnected pedal is allowed. A connected pedal requires a
finite voltage and a validated source resistance. Experimental
values are explicitly tagged; no authentic B3 impedance is asserted.

This does not establish the actual pedal winding connection,
transformer tap arrangement, loading or transfer function. These
still require terminal-level transcription of the service schematic.
