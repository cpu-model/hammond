package registration

// B3PedalBusbarsHandyGuide assigns each physical contact from
// B3PedalContactOrderFigure7 to the four colored pedal busbar wires.
// Busbar 1: brown (fundamental/3rd); 2: red (2nd/4th);
// 3: orange (6th/8th); 4: yellow (10th/12th).
// Source: B3/C3 Handy Wiring Guide, page 5, colored pedal wires.
// https://www.hammondx66.com/wp-content/uploads/2022/09/Handy_Wiring_Guide_B3_C3.pdf
func B3PedalBusbarsHandyGuide() PedalBusRouting {
 // Physical order: 10th, 6th, 2nd, fundamental,
 //                 12th, 8th, 4th, 3rd.
 return PedalBusRouting{4,3,2,1,4,3,2,1}
}

// B3PedalWiringHandyGuide combines the physical contact layout,
// the color-coded busbars and Table III's frequency assignments.
// Verified is a caller attestation, not a claim of circuit simulation.
// The analog resistor panel, choke and pedal drawbars are not included.
func B3PedalWiringHandyGuide() PedalWiringEvidence {
 return PedalWiringEvidence{
  ContactOrder:B3PedalContactOrderFigure7(),
  Busbars:B3PedalBusbarsHandyGuide(),
  Source:"Hammond B3/C3 Service Manual Table III and Figure 7; B3/C3 Handy Wiring Guide page 5",
  Verified:true,
 }
}
