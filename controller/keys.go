package main

import "machine"

// KeyState represents the state of a key.
type KeyState int

const (
	KeyNeverPressed KeyState = 0
	KeyPress                 = 1
	KeyRelease               = 2
)

// KeyEvent represents when the state of a Key changes.
type KeyEvent int

const (
	KeyNotChanged KeyEvent = 0
	KeyPressed             = 1
	KeyReleased            = 2
)

// Key is a button that is being used as a key for the MIDI controller.
type Key struct {
	button   machine.Pin
	state    KeyState
	readings *ReadingBuffer
}

// NewKey creates a new Key.
func NewKey(button machine.Pin) *Key {
	button.Configure(machine.PinConfig{Mode: machine.PinInput})
	button.Set(false)

	return &Key{
		button:   button,
		state:    KeyNeverPressed,
		readings: NewReadingBuffer(),
	}
}

// Get returns a KeyEvent based on the most recent state of the key,
// and if it has changed by being pressed or released.
func (k *Key) Get() KeyEvent {
	// since pin is pulled up, a low value means the key is pressed
	pressed := !k.button.Get()
	avg := k.readings.Avg()
	k.readings.Put(pressed)

	if pressed && avg > 0 {
		if k.state == KeyPress {
			return KeyNotChanged
		}

		k.state = KeyPress
		return KeyPressed
	} else if !pressed {
		if k.state == KeyPress {
			k.state = KeyRelease
			return KeyReleased
		}
	}

	return KeyNotChanged
}
