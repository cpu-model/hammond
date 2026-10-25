# T1/T2 physical-model evidence checklist

The engine currently has separate T1 and T2 connection records, not
transformer circuit models. `TransformerEvidencePair` records whether
four evidence categories have source identifiers for **each**
transformer: winding topology, tap topology, impedance and transfer.

`ReadyForCircuitModel` rejects missing references and swapped T1/T2
identities. Passing the check means references were supplied; it does
**not** mean the referenced data were checked, that numerical values
were transcribed, or that an electrical solver exists.

Current status for both T1 and T2:
- Winding topology: unverified.
- Tap connectivity: unverified at circuit level.
- Impedance and load: unverified.
- Frequency-dependent transfer: unverified.

Do not populate these fields with placeholder citations. Source
documents must be inspected and values cross-checked before enabling
a physically asserted model. The experimental resistive bus path
remains separate from the legacy audio output.
