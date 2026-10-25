# B-3 pedal wiring: Figure 11 is the applicable circuit

Primary source: Hammond Organ Service Manual (Early Models), printed pages 2-20 to 2-22, PDF pages 44-46 (zero-based), https://medias.audiofanzine.com/files/hammond-service-manual-a-a-100-ba-bc-bcv-bv-b2-b3-c-cv-c2-c2-text-470145.pdf

## Source-backed findings

- The manual identifies four pedal busbars for 25-pedal consoles (printed p. 2-20).
- Figure 7 gives the eight-contact arrangement for later 25-pedal consoles (printed p. 2-21).
- **Figure 11 explicitly applies to B-3 and C-3 consoles, serial number 56000 and above** (printed p. 2-22). This is not a conjecture based on other models.
- Figure 11's caption points to **Figure 13, columns 1, 2, 4, 6, 8, 9, 10, 11** for pedal wiring. The OCR of the column list is imperfect; this list MUST be confirmed against a legible image before treating it as exact.
- Figure 13 is the 25-pedal frequency wiring chart (printed p. 2-23). **Figure 24 concerns 32-pedal RT/D-100 consoles, not B-3.** Do not use it for B-3 pedal frequency routing.

## Implementation constraints

- No busbar assignment has been transcribed from Figure 11. Do not fabricate one from harmonic groupings.
- The 25x8 matrix currently returned by B3PedalHarmonics is unverified and must be compared with Figure 13 before it is presented as authentic.
- The Figure 7 physical contact order is independent of the Figure 11 busbar circuit and Figure 13 frequency numbers.
- Resistive panel and choke/resistor filter remain separate physical elements; do not invent resistance or inductance values from OCR.
- Model and serial applicability must be recorded with the future factory wiring configuration.

## Required verification

Obtain a readable scan of printed pages 2-22 and 2-23. Record the eight Figure 13 column identifiers for Figure 11, all 25 frequency rows, and Figure 11 busbar connections. Only then populate a production B-3 pedal wiring fixture.
