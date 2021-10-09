package makeybutton

import "machine"

// ButtonState represents the state of a MakeyButton.
type ButtonState int

const (
	NeverPressed ButtonState = 0
	Press                    = 1
	Release                  = 2
)

// ButtonEvent represents when the state of a Button changes.
type ButtonEvent int

const (
	NotChanged ButtonEvent = 0
	Pressed                = 1
	Released               = 2
)

// Button is a "button"-like device that acts like a MakeyMakey.
type Button struct {
	pin      machine.Pin
	state    ButtonState
	readings *Buffer
}

// NewButton creates a new Button.
func NewButton(pin machine.Pin) *Button {
	pin.Configure(machine.PinConfig{Mode: machine.PinInput})
	pin.Set(false)

	return &Button{
		pin:      pin,
		state:    NeverPressed,
		readings: NewBuffer(),
	}
}

// Get returns a ButtonEvent based on the most recent state of the button,
// and if it has changed by being pressed or released.
func (b *Button) Get() ButtonEvent {
	// since pin is pulled up, a low value means the key is pressed
	pressed := !b.pin.Get()
	avg := b.readings.Avg()
	b.readings.Put(pressed)

	if pressed && avg > 0 {
		if b.state == Press {
			return NotChanged
		}

		b.state = Press
		return Pressed
	} else if !pressed {
		if b.state == Press {
			b.state = Release
			return Released
		}
	}

	return NotChanged
}
