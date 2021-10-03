# tinymusicjam

Make music using a Arduino-based customized MIDI controller.

By running an AU host program such as Hosting AU along with a software synthesizer such as Surge, the router will forward MIDI commands from a Arduino Uno connected via serial interface to the instrument.

## Controller

Controller is intended to run directly on Arduino to send MIDI commands via serial interface.

To build/flash on Arduino:

        cd controller
        tinygo flash -target arduino .

## Router

Router is intended to run on notebook computer to connect via serial to controller, and then forward the MIDI commands to a specific MIDI device.

You must be running a program that can host AU or VST plugins, and also a soft synth.

For example, Hosting AU (http://ju-x.com/hostingau.html) works well with Surge (https://surge-synthesizer.github.io/).

First, flash your controller. Then run Hosting AU, and activate Surge on "Track A".

Now you can run the `router` as follows:

        cd router
        go run . -port /dev/cu.XXX -device=0
