# B3 electrical traceability

Primary reference: Hammond B3/C3 service manual, organ schematic.

The 91-wheel generator is shared by the manuals. Each manual has nine
contact buses. The upper manual connects through its own matching
transformer T2. The lower manual and pedals use matching transformer T1.

Current software coverage:
- Shared tonewheel generation: implemented, not amplitude calibrated.
- Nine contact switching: implemented, timing unmeasured.
- Drawbar registration: implemented as approximate gain.
- Separate electrical bus loading: not implemented.
- T1 and T2 transformer response: not implemented.
- Vibrato selector and AO-28 input: not implemented.

Engineering constraint: retain distinct upper and lower electrical
paths until their separate transformer stages. Do not treat the current
digital sum or experimental post-mix RC filter as B3 circuit topology.

No resistor-wire, transformer, or preamp component values are asserted
without a primary schematic or measured source.
