# End-to-end transformer input records

`ElectricalBusEngine.NextTransformerInputs(pair)` connects the
mechanical tonewheel bank, independently timed manual contacts,
nine resistive buses per manual and drawbar tap selection metadata.

The upper manual is routed to T2, the lower to T1. Validation occurs
before advancing generator or contact clocks, so an invalid
transformer pair does not consume an audio sample.

Outputs retain individual bus voltages and selected tap positions.
There is no transformer transfer function, cross-coupling, winding
impedance or authentic drawbar attenuation yet. The resistive bus
parameters remain explicit, provenance-tagged experimental inputs.
