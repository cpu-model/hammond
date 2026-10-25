# B3 Figure 13 retrieval attempt — 2026-10-25

Primary manual: https://medias.audiofanzine.com/files/hammond-service-manual-a-a-100-ba-bc-bcv-bv-b2-b3-c-cv-c2-c2-text-470145.pdf

The web PDF index has 220 pages and explicitly identifies printed page 2-23 as **Pedal & Frequency Chart (Figure 13)**. It identifies printed page 2-22 as the Figure 11 circuit for B-3/C-3 consoles (serial 56000 and above). Source text on PDF page index 44 states that Figure 13 lists the frequency number at every pedal contact.

**Attempted retrieval:** direct PDF screenshot of page index 48 returned a cache-miss error. Downloading the full PDF in the working container also failed. Text extraction for the diagram page was not sufficient to recover its matrix. No table values have been inferred or certified.

**Critical distinction:** Figure 13's frequency-number columns are the wiring-chart columns referenced by Figure 11. These are not necessarily in the harmonic-column order of the provisional matrix in `internal/registration/b3_pedal_harmonics.go`. The existing matrix's first rows include duplicate numbers in the purported third and second harmonic columns. This is a reason to independently verify the table rather than extrapolate values.

**Acceptance criteria before factory B3 preset:**
1. Acquire a legible scan of Figure 13 and identify the eight columns Figure 11 actually references.
2. Independently transcribe 25 rows x 8 columns, including every absent connection.
3. Cross-check against physical harmonic ordering from Figure 7.
4. Trace the four switch busbars through Figure 11 to resistor panel and two pedal drawbars.
5. Preserve source page, figure, model and serial applicability in the checked-in evidence.

Until then, `B3PedalHarmonics()` remains provisional and cannot be claimed as an authentic transcription.
