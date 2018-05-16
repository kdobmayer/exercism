package circular

import (
	"errors"
)

type Buffer struct {
	data      []byte
	lastWrite int
}

func (b *Buffer) next() int {
	if b.lastWrite == len(b.data)-1 {
		return 0
	}
	return b.lastWrite + 1
}

func NewBuffer(size int) *Buffer {
	return &Buffer{
		data:      make([]byte, size),
		lastWrite: size - 1,
	}
}

func (b *Buffer) ReadByte() (byte, error) {
	for i := b.next(); i < len(b.data); i++ {
		if b.data[i] != 0 {
			read := b.data[i]
			b.data[i] = 0
			return read, nil
		}
	}
	for i := 0; i <= b.lastWrite; i++ {
		if b.data[i] != 0 {
			read := b.data[i]
			b.data[i] = 0
			return read, nil
		}
	}
	return 0, errors.New("Empty buffer")
}

func (b *Buffer) WriteByte(c byte) error {
	next := b.next()
	if b.data[next] != 0 {
		return errors.New("The buffer is full")
	}
	b.data[next] = c
	b.lastWrite = next
	return nil
}

func (b *Buffer) Overwrite(c byte) {
	next := b.next()
	b.data[next] = c
	b.lastWrite = next
}

func (b *Buffer) Reset() {
	for i := range b.data {
		b.data[i] = 0
	}
}