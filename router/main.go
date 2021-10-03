package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"

	"github.com/hybridgroup/tinymusicjam/router/commands"
	"go.bug.st/serial"
)

var (
	deviceid int
	port     string
	sp       serial.Port
	cmdr     *commands.Commander
	msg      [3]byte
)

func init() {
	flag.IntVar(&deviceid, "deviceid", 0, "MIDI device to forward messages")
	flag.StringVar(&port, "port", "/dev/ttyACM0", "serial port")
}

func main() {
	// get device ID and serial port from flags
	flag.Parse()

	// open serial port
	sp, _ = serial.Open(port, &serial.Mode{BaudRate: 115200})
	reader := bufio.NewReader(sp)

	// open commander connection to MIDI
	cmdr = commands.NewCommander(int(deviceid))

	// start listening for MIDI messages coming from serial port
	for {
		_, err := io.ReadAtLeast(reader, msg[:], 3)
		if err != nil {
			fmt.Println(err)
		}

		cmdr.WriteShort(msg[:])
	}
}
