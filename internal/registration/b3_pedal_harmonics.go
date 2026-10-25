package registration

// B3PedalHarmonics gives generator frequency numbers in the harmonic-column
// order printed in Table III of the Hammond B3/C3 Service Manual:
// fundamental, third, second, fourth, sixth, eighth, tenth, twelfth.
// A zero is a printed dash (no connection), not an inferred tonewheel.
// This is NOT the physical switch-contact order. Do not pass the returned
// table directly to NewPedalContactsState without verifying that mapping.
//
// Source: Hammond B3/C3 Service Manual, TABLE III,
// "Frequency Numbers Assigned to Keys and Pedals" (continued).
// https://synthmania.com/wp-content/uploads/2020/05/Hammond-Organ-B3-C3-Service-Manual.pdf
func B3PedalHarmonics() PedalRouting {
 return PedalRouting{
  {1,13,13,25,32,37,41,44},
  {2,14,14,26,33,38,42,45},
  {3,15,15,27,34,39,43,46},
  {4,16,16,28,35,40,44,47},
  {5,17,17,29,36,41,45,48},
  {6,18,18,30,37,42,46,49},
  {7,19,19,31,38,43,47,50},
  {8,0,20,32,39,44,48,51},
  {9,0,21,33,40,45,49,52},
  {10,0,22,34,41,46,50,53},
  {11,0,23,35,42,47,51,54},
  {12,0,24,36,43,48,52,55},
  {13,32,25,37,44,49,53,56},
  {14,33,26,38,45,50,54,57},
  {15,34,27,39,46,51,55,58},
  {16,35,28,40,47,52,56,59},
  {17,36,29,41,48,53,57,60},
  {18,37,30,42,49,54,58,61},
  {19,38,31,43,50,55,59,62},
  {20,39,32,44,51,56,60,63},
  {21,40,33,45,52,57,61,64},
  {22,41,34,46,53,58,62,65},
  {23,42,35,47,54,59,63,66},
  {24,43,36,48,55,60,64,67},
  {25,44,37,49,56,61,65,68},
 }
}
