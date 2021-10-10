package midi

import (
	"fmt"
	"log"

	"github.com/rakyll/portmidi"
)

// Stream sends the actual MIDI commands
type Stream struct {
	out *portmidi.Stream
}

func NewStream(id int) *Stream {
	err := portmidi.Initialize()
	if err != nil {
		panic(err)
	}
	fmt.Println("[INFO] total MIDI devices:", portmidi.CountDevices())
	fmt.Println(portmidi.Info(portmidi.DeviceID(id)))
	out, err := portmidi.NewOutputStream(portmidi.DeviceID(id), 1024, 0)
	if err != nil {
		log.Fatal(err)
	}

	return &Stream{out: out}
}

func (s *Stream) WriteShort(msg []byte) {
	s.out.WriteShort(int64(msg[0]), int64(msg[1]), int64(msg[2]))
}
