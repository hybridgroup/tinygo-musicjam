package main

import (
	"machine"
	"time"

	"github.com/hybridgroup/tinymusicjam/controller/makeybutton"
	"github.com/hybridgroup/tinymusicjam/controller/midi"
	"github.com/hybridgroup/tinymusicjam/controller/notes"
)

const (
	numberOfKeys = 4
	keyOfMusic   = notes.C
)

var (
	buttonC machine.Pin = machine.D2
	buttonD machine.Pin = machine.D4
	buttonE machine.Pin = machine.D6
	buttonG machine.Pin = machine.D9

	keys        [4]*makeybutton.Button
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
	keys[1] = makeybutton.NewButton(buttonD)
	keys[2] = makeybutton.NewButton(buttonE)
	keys[3] = makeybutton.NewButton(buttonG)
}
