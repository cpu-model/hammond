# Timed two-manual instrument engine

`engine.NewTimed` integrates two independently scheduled nine-contact
manuals with one shared, continuously running mechanical generator.

Both manuals are advanced exactly once per audio sample, and they read
the same generator frame. Contact make/break delays are configured as
integer samples, separately for each harmonic contact and each manual.

This enables deterministic experiments with staggered key contacts.
**No actual B3 timing measurements or contact bounce parameters have
been asserted.** Zero delays are the default. This is not yet a
validated key-click simulation.

The existing `engine.New` remains unchanged for comparison.
Neither engine includes pedalboard, electrical loading, preamp, vibrato,
percussion, Leslie, or a realtime audio adapter.
