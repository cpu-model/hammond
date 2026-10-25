# Key-contact timing research model

The B3's nine harmonic key contacts are independent mechanical switches.
The new `registration.TimedManual` supports configurable per-contact
make and break delays (integer audio samples), a continuously running
shared generator input, and deterministic event ordering.

This implementation is **not yet an authentic key-click model**.
It uses ideal switch transitions and no measured timing, bounce,
contact resistance, bus capacitance, leakage, or preamplifier response.
A nonzero delay must be explicitly supplied by the caller. The default
timing is zero for all nine contacts; no undocumented values are hidden
in the model.

The existing `registration.Manual` remains available for direct
zero-latency routing and for the engine's current integration. The
timed manual is an independently tested research component and is
not yet connected to the two-manual engine.

Next: integrate timed contacts into the engine, establish source-based
timing ranges and electrical transient response, then validate attack
recordings against reference material.
