# Experimental generator-to-bus signal path

`engine.ElectricalBusEngine` connects a single continuously running
91-output mechanical generator to two independent timed manuals.
For each sample it advances the generator once, advances each manual's
contact events once, and injects each active routed tonewheel voltage
into the corresponding one of nine independent resistive buses.

`NextBuses()` returns two nine-element arrays of nodal voltages.
The engine does not sum manual outputs or implement drawbar taps,
matching transformers, vibrato, or AO-28 amplification. It is separate
from the existing audio engine.

All 18 bus configurations are explicit, source-tagged parameters.
No Hammond resistor or transformer value has been inferred.

**Limitations:** ideal sine generator outputs, ideal switched contacts,
identical source resistance within each bus, simplified ground loading,
and no electrical back-reaction onto a shared tonewheel pickup.
The engine uses `TimedManual.AdvanceContacts` to update contact events without evaluating the legacy digital mix. The original `Mix` retains its sample timing for existing callers.
