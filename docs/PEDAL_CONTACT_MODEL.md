# B3 pedal switch model

`registration.PedalContactsState` models 25 pedal keys, each with
eight switch contacts, independently from the 61-key, nine-contact
manuals.

Sources:
- Hammond B3/C3 service manual, section 1-8 (25 pedal levers):
  https://synthmania.com/wp-content/uploads/2020/05/Hammond-Organ-B3-C3-Service-Manual.pdf
- Hammond technical description, section 5-36 (eight contacts and
  four pedal switch busbars):
  https://b3world.com/hammond-technical-information-03.html

`PedalRouting` contains eight tonewheel references per pedal key.
Zero means unspecified/disconnected. **No historical B3 wiring
chart is populated yet**; users of this experimental model must
supply routing explicitly. Switch contacts currently close and
open together, without timing or bounce simulation.

Important distinction: the documented pedal switch has **four
busbars**, despite having eight contacts per pedal key. Do not
mistake contact count for the number of independent buses.
The present model enumerates individual contacts and does not
implement pedal busbar wiring, filter reactor or resistor network.
