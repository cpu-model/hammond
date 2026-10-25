# Manual and registration model

The current implementation provides an explicit 61-key, nine-contact
routing matrix for one manual. Each contact refers to a shared physical
generator signal (number 1..91), or zero for disconnected.

This **does not yet contain the B3's actual wiring map or foldback**.
Those values must be transcribed and checked against the original
Hammond B3/C3 wiring guide:
https://www.hammondx66.com/wp-content/uploads/2022/09/Handy_Wiring_Guide_B3_C3.pdf

The mixer supports key-down/key-up and nine drawbar positions 0..8,
but its linear level = position/8 is **not** a verified Hammond
drawbar attenuation curve. It is a temporary engineering interface,
not an acoustic authenticity claim.

The manual does not own or advance oscillators. All manuals and pedals
must share one continuously running mechanical generator.

Next: transcribe the factory contact map and foldback, implement the
resistive drawbar mixing network, add second manual and pedals, and
integrate a playable input/output adapter.
