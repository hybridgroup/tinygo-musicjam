package commands

import (
	"machine"
)

type Commander struct {
	msg [3]byte
}

func NewCommander() *Commander {
	return &Commander{}
}

func (c *Commander) NoteOn(channel, note, velocity uint8) {
	c.msg[0], c.msg[1], c.msg[2] = 0x90|(channel&0xf), note&0x7f, velocity&0x7f
	c.send()
}

func (c *Commander) NoteOff(channel, note, velocity uint8) {
	c.msg[0], c.msg[1], c.msg[2] = 0x80|(channel&0xf), note&0x7f, velocity&0x7f
	c.send()
}

func (c *Commander) SendCC(channel, control, value uint8) {
	c.msg[0], c.msg[1], c.msg[2] = 0xB0|(channel&0xf), control&0x7f, value&0x7f
	c.send()
}

func (c *Commander) send() {
	machine.Serial.Write(c.msg[:])
}
