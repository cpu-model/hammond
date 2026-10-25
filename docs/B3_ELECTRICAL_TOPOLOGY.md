# B3/C3 electrical signal path: source-grounded inventory

Status: research inventory only. No unverified component values have been
introduced into the DSP signal chain.

## Primary source

Hammond B3/C3 service manual, manual chassis and preamplifier sections
5-31 through 5-42, plus B3/C3 schematic (PDF pages 71-72, zero-based
pages 70-71):

https://www.lost-manuals.com/pdf/hammond--b-3--service--ID201.pdf

## Confirmed topology

1. Each tone-generator output reaches one or more playing-key contacts
   through **resistance wires**. The wires decouple the shared generator
   from simultaneously connected key circuits (section 5-32).
2. A depressed playing key connects up to nine generator signals to
   the corresponding **nine manual busbars** (section 5-32).
3. A selected preset/adjust key routes those busbars to either a preset
   registration or nine drawbars (sections 5-32 to 5-34).
4. Each drawbar selects a tap on its manual's **mixing transformer**.
   There are separate transformers for upper and lower manual paths,
   T2 and T1 respectively (sections 5-33 to 5-35, schematic).
5. The transformer secondary feeds the preamplifier through a manual
   vibrato on/off selector. Vibrato and non-vibrato inputs follow
   distinct paths (sections 5-39 to 5-42).

These are separate electrical stages. A single post-mix low-pass filter
does **not** substitute for the resistance-wire / busbar / transformer
network and is retained solely as an experiment.

## Component-specific work still required

- Resistance wire impedance and its frequency dependence; effects of
  multiple held keys on the same generator output.
- Nine busbar loads and selected drawbar-tap electrical coupling.
- Transformer turns/tap ratios, magnetization and loading.
- Vibrato selector loading and the AO-28 input impedance.
- Contact bounce, leakage, parasitic capacitance and transient behavior.

The manual and schematic establish topology but do not, by themselves,
supply enough verified values for a calibrated nodal model. No circuit
values should be inferred from the existing 3 dB drawbar approximation.

## Proposed implementation boundary

Tonewheel source -> decoupling impedance -> switched key contact ->
manual busbar -> preset/adjust switch -> drawbar tap ->
manual mixing transformer -> vibrato selector -> AO-28 input.

Preserve one shared generator and two separate manual networks until
the actual schematic calls for combining signals. Use nodal/circuit
analysis only after component parameters and connectivity are traced
to documented evidence.
