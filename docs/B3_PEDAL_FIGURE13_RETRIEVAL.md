# Figure 13 verification blocker and figure-number disambiguation

Primary source: Hammond Organ Service Manual - Early Models, 220-page PDF, https://medias.audiofanzine.com/files/hammond-service-manual-a-a-100-ba-bc-bcv-bv-b2-b3-c-cv-c2-c2-text-470145.pdf

Verified from the extracted source:
- PDF page 44 (zero based): the 25-pedal assembly has four busbars and eight contacts; the text explicitly directs the reader to Figure 13 for the pedal frequency numbers.
- PDF page 45: Figure 7 labels the eight contact harmonics.
- PDF page 46: Figure 11 is the pedal circuit applicable to B-3/C-3 serial 56000 and later.
- PDF page 47: a *different* pedal assembly for 32-pedal RT and D-100 models is described, and its Figure 24 is a 32-pedal frequency chart.
- The contents also call a **Figure 24** the B-3/C-3/RT-3 *console schematic* (section III, printed page 3-12). Figure numbers repeat between manual sections. Thus a bare reference to Figure 24 is ambiguous and must include section/page.

Unverified:
- The provisional 25x8 numbers in B3PedalHarmonics() have not been checked against Figure 13.
- Exact Figure 11 connections to four busbars remain untranscribed.
- The OCR of Figure 11's referenced Figure 13 columns is too garbled to trust.

Retrieval result:
- PDF text was available from web retrieval, but screenshot of PDF page 48 failed with a cache-miss error. The runtime could not download the PDF by direct HTTPS (DNS unavailable). No numerical transcription from the unseen diagram was attempted.

Next action:
Acquire a readable image of PDF pages 46 and the actual Figure 13 page (located by its printed 2-23 label, potentially after inserted 2-22A-D pages). Verify every number and trace the physical busbar connections before enabling an authenticated factory preset.
