# tinymusicjam

Make music using a Arduino-based customized MIDI controller.

By running an AU host program such as AU Lab along with a software synthesizer such as Surge, the router will forward MIDI commands from a Arduino Uno connected via serial interface to the instrument.

## Controller

Controller is intended to run directly on Arduino to send MIDI commands via serial interface.

## Router

Router is intended to run on notebook computer to connect via serial to controller, and then forward the MIDI commands to a specific MIDI device.

