package commands

import (
	"fmt"
	"log"

	"github.com/rakyll/portmidi"
)

// Commander sends the actual MIDI commands
type Commander struct {
	out *portmidi.OutputStream
}

func NewCommander(id int) *Commander {
	err := portmidi.Initialize()
	if err != nil {
		panic(err)
	}
	fmt.Println("[INFO] total MIDI devices:", portmidi.CountDevices())

	out, err := portmidi.NewOutputStream(portmidi.DeviceID(id), 1024, 0)
	if err != nil {
		log.Fatal(err)
	}

	return &Commander{out: out}
}

// func (c *Commander) NoteOn(channel, note, velocity uint8) {
// 	var status uint8 = 0x80
// 	if velocity != 0 {
// 		status = 0x90
// 	}
// 	c.out.WriteShort(status|(channel&0xf), note&0x7f, velocity&0x7f)
// }

// func (c *Commander) NoteOff(channel, note, velocity uint8) {
// 	c.out.WriteShort(0x80|(channel&0xf), note&0x7f, velocity&0x7f)
// }

// func (c *Commander) SendCC(channel, control, value uint8) {
// 	c.out.WriteShort(0xB0|(channel&0xf), control&0x7f, value&0x7f)
// }

func (c *Commander) WriteShort(msg []byte) {
	c.out.WriteShort(msg[0], msg[1], msg[2])
}
