# TinyGo Music Jam

Make music using your own Arduino-based customized MIDI controller using audio software running on your notebook computer.

```
┌────────────────────────┐          ┌──────────────────────┐
│                        │          │                      │
│                        │          │                      │
│     MIDI Controller    │          │     MIDI Router      │
│                        ├──────────►                      │
│                        │          │                      │
│                        │          │                      │
│                        │          │                      │
│                        │          │                      │
└────────────────────────┘          └──────────┬───────────┘
                                               │
                                               │
                                               │
                                               │
                                               │
                                    ┌──────────▼───────────┐
                                    │                      ├─────────────────────┐
                                    │                      │                     │
                                    │    AU or VST Host    │   Software Synth    │
                                    │                      │                     │
                                    │                      ├─────────────────────┘
                                    │                      │
                                    │                      │
                                    │                      │
                                    └──────────────────────┘
```

By running a AU or VST host program such as Hosting AU along with a software synthesizer such as Surge, the MIDI router will forward MIDI commands from a Arduino Uno connected via serial interface to the virtual instrument.

## Installation

You will need to install Go and TinyGo to compile the code.

You will also need to install audio programs on your local computer in order to produce beautiful music.

### macOS

Hosting AU (http://ju-x.com/hostingau.html)

Surge (https://surge-synthesizer.github.io/)

### Windows

VSTHost is an open source VST host program for Windows. You can obtain it from here: https://www.hermannseib.com/english/vsthost.htm

Surge (https://surge-synthesizer.github.io/)

### Linux

Please add instructions if you try this.

## Controller

The MIDI Controller is intended to run directly on Arduino to send MIDI commands via serial interface.

There are several different kinds of controllers in the `examples` folder.

To build/flash the `onenote` example on Arduino:

        cd controller
        tinygo flash -target arduino ./examples/onenote/

## Router

Router is intended to run on notebook computer to connect via serial to controller, and then forward the MIDI commands to a specific MIDI device.

You must be running a program that can host AU or VST plugins, and also a soft synth.

First, flash your controller as described above. Then run Hosting AU, and activate Surge on "Track A".

Now you can run the `router` as follows:

        cd router
        go run . -port /dev/cu.XXX -device=0

Have fun!