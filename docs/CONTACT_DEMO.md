# Contact-timing comparison

The offline demo lets us compare zero-delay key contacts with
deliberately staggered closures, without inventing B3 measurements.

```sh
go run ./cmd/contactdemo -step 0 -out simultaneous.wav
go run ./cmd/contactdemo -step 12 -out staggered.wav
```

At 48 kHz, `-step 12` means each successive contact closes 0.25 ms
after the previous one. This is an **experimental parameter**, not
a claimed measurement of Hammond B3 contacts.

Both recordings use the same transcribed B3 manual routing, the first
three drawbars at 8, and release the key at two-thirds of the recording.
Each recording is peak-normalized independently, so use these files to
compare attack shapes, not calibrated loudness.

No contact bounce, circuit loading, leakage, or preamp response is yet
modeled. Therefore neither recording is a validated B3 key-click sound.
