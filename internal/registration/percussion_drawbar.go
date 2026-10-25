package registration

// SetPercussionDrawbarCut applies the requested ninth-drawbar state.
// The drawbar's stored registration is unchanged.
func (m *TimedManual) SetPercussionDrawbarCut(enabled bool) {
 m.percussionDrawbarCut=enabled
 if enabled {m.ninthTarget=0} else {m.ninthTarget=1}
}

// SetNinthDrawbarTransition configures the gain transition in samples.
// One sample preserves the immediate-switch behavior.
func (m *TimedManual) SetNinthDrawbarTransition(samples int) {
 if samples<1 {samples=1}
 m.ninthStep=1/float64(samples)
}
