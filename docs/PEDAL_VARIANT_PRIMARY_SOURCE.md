# Primary-source finding: pedal wiring depends on console variant

Source: Hammond Organ Service Manual, Early Models, section MANUALS & PEDALS, printed page 2-20 (PDF page index 44), and illustrations on subsequent pages.
https://medias.audiofanzine.com/files/hammond-service-manual-a-a-100-ba-bc-bcv-bv-b2-b3-c-cv-c2-c2-text-470145.pdf

The service text explicitly distinguishes early seven-contact pedal assemblies from later eight-contact assemblies. Later 25-pedal consoles use eight contacts, with a different harmonic contact arrangement and a circuit that mixes 16-foot and 8-foot pedal tones. Four colored wires connect pedal switch busbars to pedal drawbars, sometimes via a resistor panel. The lower pedal drawbar is filtered by a choke and resistor.

The manual states that specific pedal wiring must be selected from Figures 8 through 11 using the console serial number, and that Figure 13 is the pedal-to-frequency chart. Therefore, a single universal contact order/bus mapping for every Hammond organ would be incorrect. Figure 7 identifies the eight-contact arrangement for later 25-pedal consoles. Figures 10 and 11 are the relevant later mixing circuit family; the precise B3 variant must still be identified.

Implementation consequences:
1. Do not promote the provisional B3PedalHarmonics matrix to verified factory wiring without comparing all rows against Figure 13.
2. Transcribe Figure 7 contact order, Figure 13 frequency numbers, and the B3-applicable Figure 10/11 connections independently.
3. Record the source figure, applicable console model/serial range, and verification method for each table. The existing Verified flag only records caller attestation.
4. Keep the resistor/choke filter and the resistor panel separate in the electrical netlist.

This is an evidence inventory, not a completed transcription.
