package main

import (
	"machine"
	"time"

	"github.com/hybridgroup/tinymusicjam/controller/makeybutton"
	"github.com/hybridgroup/tinymusicjam/controller/midi"
)

const (
	middleC = 60
)

var (
	buttonC machine.Pin = machine.D2

	key         *makeybutton.Button
	sender      *midi.Sender
	midichannel uint8 = 0 // MIDI channels are 0-15 e.g. 1-16
)

func main() {
	// open MIDI connection to serial
	sender = midi.NewSender()

	key = makeybutton.NewButton(buttonC)

	for {
		switch key.Get() {
		case makeybutton.Pressed:
			sender.NoteOn(midichannel, uint8(middleC), 100)
		case makeybutton.Released:
			sender.NoteOff(midichannel, uint8(middleC), 100)
		}
		time.Sleep(30 * time.Millisecond)
	}
}
