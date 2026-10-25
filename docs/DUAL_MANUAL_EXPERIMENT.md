# Separate manual electrical-response experiment

The experimental `engine.NewDualFilteredTimed` preserves independent
upper and lower manual paths from the shared mechanical generator.
Each path has its own RC-equivalent filter state and explicit cutoff.

This lets us test that two electrical paths can be processed separately
without accidentally combining signals or coupling their filter memory.

Neither cutoff is a Hammond B3 specification. These filters are **not**
models of T1 and T2: their winding, tap and load characteristics remain
unmeasured and unimplemented. The diagnostic mono sum remains only a
listening convenience.

Tests compare both output paths sample-for-sample against independent
reference recurrences and check allocation-free rendering.
