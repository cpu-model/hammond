package registration

// B3PedalContactOrderFigure7 is the physical top-then-bottom contact
// order shown in Figure 7 (service manual, printed page 2-21) for
// later 25-pedal Hammond consoles, including the B-3.
// Top: 10th, 6th, 2nd, fundamental.
// Bottom: 12th, 8th, 4th, 3rd.
// Indices address the harmonic columns in B3PedalHarmonics:
// fundamental, 3rd, 2nd, 4th, 6th, 8th, 10th, 12th.
// This is NOT a busbar assignment or a verification of the frequency
// numbers in the provisional harmonic table.
func B3PedalContactOrderFigure7() PedalContactOrder {
 return PedalContactOrder{6,4,2,0,7,5,3,1}
}
