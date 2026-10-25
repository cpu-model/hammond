# Nine independent manual bus conductors

`electrical.NineBusNetwork` contains nine independent resistive
nodes, one for each manual harmonic bus conductor. Each bus requires
an explicit source resistance and ground load with provenance.

At each sample the caller clears all nine nodes, adds voltages only
for closed contacts, and obtains nine separate bus voltages. Multiple
contacts on one bus share its resistive loading. Different buses are
electrically isolated in this simplified network.

The component is not yet connected to the keyboard engine. Actual
Hammond wiring can share tonewheel sources across contacts, and the
real buses couple through preset/drawbar transformer circuitry. Those
effects are not represented here. The load to ground is a test
simplification, not a claimed B3 schematic value.

Next integration requires exposing individual active contacts from
`registration.TimedManual` without advancing contact clocks twice.
The resulting nine bus voltages must remain separate until a
documented transformer tap model is available.
