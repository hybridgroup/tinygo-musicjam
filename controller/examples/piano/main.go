package main

import (
	"machine"
	"time"

	"github.com/hybridgroup/tinymusicjam/controller/makeybutton"
	"github.com/hybridgroup/tinymusicjam/controller/midi"
	"github.com/hybridgroup/tinymusicjam/controller/notes"
)

const (
	numberOfKeys = 1
	keyOfMusic   = notes.C
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

	keys        [12]*makeybutton.Button
	sender      *midi.Sender
	midichannel uint8 = 0 // MIDI channels are 0-15 e.g. 1-16
)

func main() {
	// open MIDI connection to serial
	sender = midi.NewSender()

	initKeys()

	for {
		readKeys()
		time.Sleep(30 * time.Millisecond)
	}
}

func readKeys() {
	for key := 0; key < numberOfKeys; key++ {
		switch keys[key].Get() {
		case makeybutton.Pressed:
			sender.NoteOn(midichannel, uint8(keyOfMusic+key), 100)
		case makeybutton.Released:
			sender.NoteOff(midichannel, uint8(keyOfMusic+key), 100)
		}
	}
}

func initKeys() {
	keys[0] = makeybutton.NewButton(buttonC)
	keys[1] = makeybutton.NewButton(buttonDb)
	keys[2] = makeybutton.NewButton(buttonD)
	keys[3] = makeybutton.NewButton(buttonEb)
	keys[4] = makeybutton.NewButton(buttonE)
	keys[5] = makeybutton.NewButton(buttonF)
	keys[6] = makeybutton.NewButton(buttonGb)
	keys[7] = makeybutton.NewButton(buttonG)
	keys[8] = makeybutton.NewButton(buttonAb)
	keys[9] = makeybutton.NewButton(buttonA)
	keys[10] = makeybutton.NewButton(buttonBb)
	keys[11] = makeybutton.NewButton(buttonB)
}
