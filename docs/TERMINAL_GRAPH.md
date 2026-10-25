# Schematic terminal graph

`electrical.TerminalGraph` is a source-referenced graph of
**conductive terminal-to-terminal connections**. Each edge must
name two distinct terminals and identify its documentary source.
Reversed duplicate edges are rejected.

`Connected(a,b)` answers whether a conductive path has been
transcribed. It does not assume that magnetically coupled windings
are electrically shorted, and it computes no voltage transfer.

The tests use **synthetic fixtures**, not claimed B3 terminal names.
No real T1/T2 winding edges have been asserted in production data.

To transcribe the actual schematic, record each terminal label,
the exact diagram page and drawing region for each edge, and
separately represent winding coupling. Do not add speculative
links or fill the graph merely to pass tests.
