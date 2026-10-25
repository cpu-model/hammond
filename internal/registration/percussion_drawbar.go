package registration

// SetPercussionDrawbarCut controls the B3 ninth-drawbar cancellation.
// The stored drawbar position is retained when cancellation is enabled.
func (m *TimedManual) SetPercussionDrawbarCut(enabled bool) {
 m.percussionDrawbarCut=enabled
}
