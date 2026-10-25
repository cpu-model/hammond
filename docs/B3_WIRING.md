# B3/C3 factory-style manual routing

The 61 x 9 contact matrix in `internal/registration/b3_routing.go`
was transcribed from the **modern narrow 91-tone generator** table
published by Jeff Dairiki:

https://www.dairiki.org/hammond/wiring/

It is intended for the B3/C3 and related narrow-generator consoles.
The selected table explicitly includes upper-register foldback.
The first key is low C; the 61st key is high C.

Selected source cross-checks:

| Key | Drawbar | Frequency number |
|---:|---:|---:|
| 1 | 1 (16') | 13 |
| 1 | 2 (5⅓') | 20 |
| 1 | 3 (8') | 13 |
| 13 | 9 (1') | 61 |
| 43 | 9 (1') | 91 |
| 44 | 9 (1') | 80 |
| 61 | 9 (1') | 85 |

**Status:** Complete secondary-source transcription, with boundary tests.
An independent check against the factory wiring diagram remains necessary.
The wiring map is not proof of accurate tonewheel timbre or drawbar gains.

Original wiring reference:
https://www.hammondx66.com/wp-content/uploads/2022/09/Handy_Wiring_Guide_B3_C3.pdf

Use `registration.B3ManualRouting()` for both manuals in the
instrument engine. No modulo-based foldback approximation is used.
