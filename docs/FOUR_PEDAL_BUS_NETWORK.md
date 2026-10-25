# Four experimental pedal bus nodes

`electrical.FourPedalBusNetwork` implements four independently
loaded resistive nodes, corresponding to the **count** of pedal
switch busbars described in Hammond's technical manual (5-36):
https://b3world.com/hammond-technical-information-03.html

The four buses accept closed contact voltages and calculate their
conductance-weighted node voltages. All source and ground-load
resistances are explicit, validated, provenance-tagged inputs.

This is **not** a B3 pedal tone circuit: the manual states that
the four buses connect through a filter reactor and resistor
network to the pedal drawbars. That cross-bus/downstream circuit
has not yet been transcribed or solved.

Schematic component labels include R80–R86 and L20 in the
pedal-filter/resistor-panel region, but label proximity alone
is not sufficient to infer individual electrical edges or
component values. No such edges are claimed in this module.
