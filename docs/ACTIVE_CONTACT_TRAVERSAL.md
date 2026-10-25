# Active-contact traversal for electrical integration

`TimedManual.ForEachActiveContact` exposes each active, connected
contact as a pair of zero-based harmonic bus index and one-based
generator wheel index. It does not apply drawbar gain and does not
advance the sample clock.

The existing `Mix(frame)` method still advances contact events.
Callers must advance once per sample before traversing contacts.
Calling `Mix` merely to advance events also computes the legacy
digital mix; it is not an electrical model.

This is a read-only interface for a future bridge from mechanical
generator frames to the nine resistive bus nodes. It avoids imposing
unverified B3 transformer or drawbar loading values.
