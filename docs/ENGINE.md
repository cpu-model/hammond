# Instrument engine composition

`internal/engine` joins **one continuously running 12-speed-group
mechanical generator** with **two 61-key manuals**. Every output sample:

1. Advances all mechanical speed-group phases once.
2. Produces all 91 generator signals in a reusable frame.
3. Mixes the upper and lower manual contacts from that same frame.

Releasing all keys silences the mix but does not stop or reset the generator.
Repressing a key resumes the current mechanical phase. Rendering allocates
no memory in the sample loop.

**Not yet a playable authentic B3:** both manuals require an explicitly
provided wiring table. There is no invented default routing or foldback.
The mixer is linear and does not model contact sequencing, resistive networks,
pedals, percussion, scanner vibrato, expression, preamp or Leslie.

The engine is designed for subsequent keyboard/MIDI and audio adapters;
those adapters have not been implemented.
