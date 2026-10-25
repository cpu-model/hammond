# Experimental resistive contact network

The new `electrical.ResistiveContact` models one ideal switch with
one source resistance and one load resistance.

For a closed contact the bus voltage is
`source * load / (source + load)`. An open contact returns zero.

Both resistances require explicit units and provenance. The example
values in the unit tests are **not Hammond B3 measurements**.

This is a small, independently testable circuit primitive, not a
complete B3 tonewheel resistance-wire or nine-busbar network.
It excludes parallel connected keys, generator source impedance,
transformer taps, parasitic capacitance and contact transients.

The component is not inserted into the audio engine until a
physically grounded connection topology and parameter set exist.
