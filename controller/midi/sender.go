package midi

import (
	"machine"
)

// Sender sends MIDI messages over the serial port.
type Sender struct {
	msg [3]byte
}

// NewSender returns a new Sender.
func NewSender() *Sender {
	return &Sender{}
}

// NoteOn sends a note on message.
func (s *Sender) NoteOn(channel, note, velocity uint8) {
	s.msg[0], s.msg[1], s.msg[2] = 0x90|(channel&0xf), note&0x7f, velocity&0x7f
	s.send()
}

// NoteOff sends a note off message.
func (s *Sender) NoteOff(channel, note, velocity uint8) {
	s.msg[0], s.msg[1], s.msg[2] = 0x80|(channel&0xf), note&0x7f, velocity&0x7f
	s.send()
}

// SendCC sends a continuous controller message.
func (s *Sender) SendCC(channel, control, value uint8) {
	s.msg[0], s.msg[1], s.msg[2] = 0xB0|(channel&0xf), control&0x7f, value&0x7f
	s.send()
}

func (s *Sender) send() {
	machine.Serial.Write(s.msg[:])
}
