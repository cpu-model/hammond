# Experimental pedal reactor stage

The B3 pedal circuit includes a filter reactor and resistor network,
but the actual connectivity and values are not yet verified here.

`electrical.PedalRLStage` is a separate, **synthetic topology**:
ideal input source -> series resistor -> series inductor -> load
resistor to ground. Its state is the load voltage.

For sample interval dt, total series-plus-load resistance R,
load resistor Rload and inductance L, backward Euler uses:

- alpha = 1 / (1 + dt*R/L)
- output[n] = alpha*output[n-1] + (1-alpha)*(Rload/R)*input[n]

This permits stable time-domain experiments with explicit
provenance-tagged H/ohm parameters. It is not connected to the
four pedal buses, because the actual bus-to-reactor topology
has not been transcribed from the B3 schematic. No component
values or physical transfer function are claimed as authentic.

Reference for presence of the reactor/resistor network:
https://b3world.com/hammond-technical-information-03.html
