package commands

import (
	"fmt"
	"log"

	"github.com/rakyll/portmidi"
)

// Commander sends the actual MIDI commands
type Commander struct {
	out *portmidi.Stream
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

func (c *Commander) WriteShort(msg []byte) {
	c.out.WriteShort(int64(msg[0]), int64(msg[1]), int64(msg[2]))
}
