# Resistive bus experiment

The `ResistiveBus` component models a single electrically loaded
node fed by zero or more closed-contact sources. Each source uses
an explicitly configured series resistance; the bus has an explicit
resistance to ground.

Both resistances are tagged with provenance. Schematic or measured
values require source references; experimental values are explicitly
identified as such. No default B3 values are provided.

`BeginSample` clears the node, `AddSource` adds a closed contact,
and `Voltage` solves the node. Tests demonstrate how a second
contact changes the load-dependent bus voltage.

This is deliberately separate from the Hammond audio engine. Real
B3 busbars, resistance wires, drawbar tap loads, and transformer
coupling require a more complete schematic-derived network. The
one-node model is not an authentic B3 bus circuit.
