package main

const bufferSize = 10

// ReadingBuffer is a buffer to keep track of the most recent readings for a key.
type ReadingBuffer struct {
	readings [bufferSize]bool
	index    int
}

// NewReadingBuffer returns a new buffer.
func NewReadingBuffer() *ReadingBuffer {
	return &ReadingBuffer{}
}

// Used returns how many bytes in buffer have been used.
func (rb *ReadingBuffer) Used() int {
	return rb.index
}

// Put stores a byte in the buffer.
func (rb *ReadingBuffer) Put(val bool) bool {
	rb.index++
	if rb.index >= bufferSize {
		rb.index = 0
	}

	rb.readings[rb.index] = val

	return true
}

// Avg returns the "average" of all the readings in the buffer, by
// treating a true as 1 and a false as -1.
func (rb *ReadingBuffer) Avg() int {
	avg := 0
	for i := 0; i < bufferSize; i++ {
		if rb.readings[i] {
			avg += 1
		} else {
			avg -= 1
		}
	}

	return avg
}
