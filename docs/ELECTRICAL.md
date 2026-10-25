# Electrical response: first experimental building block

The `internal/electrical` package introduces a first-order low-pass
filter with a documented mathematical step response.

The model uses the exact sampled form of a continuous first-order RC
response, with coefficient `exp(-2*pi*cutoff/sampleRate)`.
It can be used in controlled experiments to study how an electrical
time constant changes contact-switching transients.

**This is not a calibrated Hammond B3 circuit.** No component values,
pickup impedance, contact resistance, bus capacitance, matching
transformer, or preamplifier topology have yet been established from
a primary-source schematic. In particular, a single RC filter must
not be presented as an authentic key-click emulation.

The package is independent of the audio engine so that filter behavior
can be tested separately before choosing a physically justified
insertion point. The cutoff is always supplied by the caller.
