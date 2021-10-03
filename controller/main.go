package main

import (
	"machine"
	"time"

	"github.com/hybridgroup/tinymusicjam/controller/commands"
)

type Key struct {
	button  machine.Pin
	pressed bool
}

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

	keys [12]Key
	cmdr *commands.Commander
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
	for i := 0; i < 12; i++ {
		if keys[i].button.Get() {
			if keys[i].pressed {
				cmdr.NoteOn(0, uint8(i+58), 100)
				keys[i].pressed = false
			}
		} else {
			if !keys[i].pressed {
				cmdr.NoteOff(0, uint8(i+58), 100)
				keys[i].pressed = true
			}
		}
	}
}

func initButtons() {
	buttonC.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonDb.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonD.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonEb.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonE.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonF.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonGb.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonG.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonAb.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonA.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonBb.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	buttonB.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

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
