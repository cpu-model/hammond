# Experimental post-mix RC response

`engine.NewFilteredTimed` connects the independently tested first-order
RC-equivalent lowpass to the output of the two-manual timed engine.
The existing `engine.NewTimed` remains unchanged as an unfiltered
reference, enabling matched comparisons.

The cutoff is explicit and must be below Nyquist. Generator phases,
contact timing and drawbar registration are otherwise identical.
`ResetFilter` resets only the filter memory, not the generator.

**Authenticity boundary:** This is an experimental post-mix topology,
not a circuit-derived Hammond B3 filter placement. Neither its cutoff
nor its insertion point is asserted as a B3 specification. In particular,
filtering the summed audio does not independently model nine key
contacts, bus loading, leakage, or key click. No default cutoff is
chosen for the instrument.

Next step: identify the actual electrical nodes and component values
from the B3/C3 schematic, then replace this test topology with a
circuit-grounded implementation.
