package main

import (
	"machine"
	"time"

	"github.com/hybridgroup/tinymusicjam/controller/commands"
)

const (
	numberOfKeys = 1
	middleC      = 60
)

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

	initKeys()

	for {
		readKeys()
		time.Sleep(30 * time.Millisecond)
	}
}

func readKeys() {
	for key := 0; key < numberOfKeys; key++ {
		switch keys[key].Get() {
		case KeyPressed:
			cmdr.NoteOn(midichannel, uint8(middleC+key), 100)
		case KeyReleased:
			cmdr.NoteOff(midichannel, uint8(middleC+key), 100)
		}
	}
}

func initKeys() {
	keys[0] = NewKey(buttonC)
	keys[1] = NewKey(buttonDb)
	keys[2] = NewKey(buttonD)
	keys[3] = NewKey(buttonEb)
	keys[4] = NewKey(buttonE)
	keys[5] = NewKey(buttonF)
	keys[6] = NewKey(buttonGb)
	keys[7] = NewKey(buttonG)
	keys[8] = NewKey(buttonAb)
	keys[9] = NewKey(buttonA)
	keys[10] = NewKey(buttonBb)
	keys[11] = NewKey(buttonB)
}
