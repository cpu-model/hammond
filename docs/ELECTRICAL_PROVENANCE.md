# Electrical parameter provenance

Every electrical model parameter must state its unit and evidence class:

- `schematic`: the primary source reference is required.
- `measurement`: a measurement record reference is required.
- `experimental`: the parameter is explicitly uncalibrated.

The code rejects nonfinite parameters, missing units, undocumented
schematic/measurement claims and unsupported provenance labels.

`ManualResponseConfig` currently records two experimental cutoff
frequencies and validates them against the sample rate. This is an
evidence-control mechanism, not evidence that an RC filter is the
correct Hammond B3 electrical circuit.

The current dual-filter engine still accepts numeric cutoff arguments;
migration to source-tagged configuration is a separate change.
