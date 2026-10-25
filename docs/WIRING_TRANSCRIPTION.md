# Factory wiring transcription gate

Before assigning actual B3 key contacts or claiming foldback authenticity,
transcribe the original wiring guide into a complete CSV table.

Primary document:
https://www.hammondx66.com/wp-content/uploads/2022/09/Handy_Wiring_Guide_B3_C3.pdf

## Required format

```csv
key,drawbar,wheel
```

- Key: 1..61 (the source's physical key order must be documented).
- Drawbar: 1..9 (the contact order must be documented).
- Wheel: generator frequency number 1..91; zero only if a
  source explicitly documents a disconnected contact.
- Exactly 549 unique key/drawbar combinations per manual.
- Any manual-specific differences must be represented in separate files.
- Record source page/diagram, transcription date and independent review
  in the accompanying provenance notes.

`registration.ParseRoutingCSV` rejects missing, duplicate and
out-of-range contacts. The unit tests use an **artificial test fixture**,
not Hammond wiring. No factory routing CSV is currently committed.

**Do not generate foldback with modulo arithmetic.** The wiring must
come from the original circuit documentation. A CSV passing structural
validation does not by itself establish historical correctness.
