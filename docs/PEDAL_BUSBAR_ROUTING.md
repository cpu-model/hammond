# Pedal contact-to-busbar interface

The B3 pedal assembly uses **eight contacts per pedal key** but
only **four busbars**. The technical description states that
the four busbars feed the pedal drawbars through a filter reactor
and resistor network.

Source:
https://b3world.com/hammond-technical-information-03.html
(paragraph 5-36).

`registration.PedalBusRouting` is an explicit mapping of each
of eight contacts to one of four busbars, encoded 1..4, or 0
for no connection. `PedalContactsState.ForEachClosedBusContact`
emits active key, contact, zero-based busbar and tonewheel
reference without mixing signals.

The mapping is **not yet transcribed from the B3 wiring chart**.
The paired-contact mapping in tests is a synthetic fixture only.
Do not treat the current interface as a completed pedal electrical
circuit. The filter reactor, resistor network, drawbar attenuation
and T1 coupling are still unimplemented.
