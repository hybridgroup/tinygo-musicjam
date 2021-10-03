package main

import (
	"machine"
	"time"

	"github.com/hybridgroup/tinymusicjam/controller/commands"
)

type KeyState int

const (
	KeyNeverPressed KeyState = 0
	KeyPress                 = 1
	KeyRelease               = 2
)

type Key struct {
	button machine.Pin
	state  KeyState
}

const numberOfKeys = 1

var (
	buttonC  machine.Pin = machine.D2
	buttonDb machine.Pin = machine.D3
	buttonD  machine.Pin = machine.D4
	buttonEb machine.Pin = machine.D5
	buttonE  machine.Pin = machine.D6
	buttonF  machine.Pin = machine.D7
	buttonGb machine.Pin = machine.D8
	buttonG  machine.Pin = machine.D9
	buttonAb machine.Pin = machine.D10
	buttonA  machine.Pin = machine.D11
	buttonBb machine.Pin = machine.D12
	buttonB  machine.Pin = machine.D13

	keys        [12]Key
	cmdr        *commands.Commander
	midichannel uint8 = 0 // MIDI channels are 0-15 e.g. 1-16
)

func main() {
	// open commander connection to serial
	cmdr = commands.NewCommander()

	initButtons()

	for {
		readButtons()
		time.Sleep(30 * time.Millisecond)
	}
}

func readButtons() {
	for i := 0; i < numberOfKeys; i++ {
		if keys[i].button.Get() {
			if keys[i].state == KeyPress {
				continue
			}

			cmdr.NoteOn(midichannel, uint8(i+60), 100)
			keys[i].state = KeyPress
		} else {
			if keys[i].state != KeyPress {
				continue
			}

			cmdr.NoteOff(midichannel, uint8(i+60), 100)
			keys[i].state = KeyRelease
		}
	}
}

func initButtons() {
	buttonC.Configure(machine.PinConfig{Mode: machine.PinInput})
	buttonDb.Configure(machine.PinConfig{Mode: machine.PinInput})
	buttonD.Configure(machine.PinConfig{Mode: machine.PinInput})
	buttonEb.Configure(machine.PinConfig{Mode: machine.PinInput})
	buttonE.Configure(machine.PinConfig{Mode: machine.PinInput})
	buttonF.Configure(machine.PinConfig{Mode: machine.PinInput})
	buttonGb.Configure(machine.PinConfig{Mode: machine.PinInput})
	buttonG.Configure(machine.PinConfig{Mode: machine.PinInput})
	buttonAb.Configure(machine.PinConfig{Mode: machine.PinInput})
	buttonA.Configure(machine.PinConfig{Mode: machine.PinInput})
	buttonBb.Configure(machine.PinConfig{Mode: machine.PinInput})
	buttonB.Configure(machine.PinConfig{Mode: machine.PinInput})

	keys[0] = Key{button: buttonC}
	keys[1] = Key{button: buttonDb}
	keys[2] = Key{button: buttonD}
	keys[3] = Key{button: buttonEb}
	keys[4] = Key{button: buttonE}
	keys[5] = Key{button: buttonF}
	keys[6] = Key{button: buttonGb}
	keys[7] = Key{button: buttonG}
	keys[8] = Key{button: buttonAb}
	keys[9] = Key{button: buttonA}
	keys[10] = Key{button: buttonBb}
	keys[11] = Key{button: buttonB}
}
