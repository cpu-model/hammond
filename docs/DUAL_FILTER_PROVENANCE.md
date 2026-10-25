# Source-tagged dual manual filters

Use `engine.NewDualFilteredTimedWithConfig` with
`electrical.ManualResponseConfig` to declare the provenance of
both cutoff parameters.

The validator requires units, finite values and source references
for values marked `schematic` or `measurement`. Unreferenced
experimental values are allowed only when explicitly marked as such.

The older numeric constructor remains for existing experiments but
labels both values experimental internally.

This change enforces evidence bookkeeping. It does not validate that
a source describes a Hammond B3 component, nor does it turn a
first-order RC filter into a matching-transformer model.
