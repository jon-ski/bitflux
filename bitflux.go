package bitflux

import (
	"bytes"
)

var (
	Le = new(le) // Little-Endian functions
	Be = new(be) // Big-Endian functions
)

// Buffer wraps `bytes.Buffer` and provides serialization methods on top of it.
type Buffer struct {
	buf bytes.Buffer
	err error // Latest/Last Error
}

// NewBuffer creates an initialized internal buffer from `buf`.
// Useful as a reader on an existing buffer.
func NewBuffer(buf []byte) *Buffer {
	return &Buffer{
		buf: *bytes.NewBuffer(buf),
	}
}

// Err returns an error if an error has occurred on the buffer
func (b *Buffer) Err() error {
	return b.err
}

// Write writes the contents of p to the underlying buffer.
// Will NOP if Buffer has stored an error and return that error.
func (b *Buffer) Write(p []byte) (n int, err error) {
	// Return early if an error already in buffer
	if b.err != nil {
		return 0, b.err
	}

	n, err = b.buf.Write(p)
	if err != nil {
		b.err = err
	}
	return n, err
}

// WriteByte writes a byte to the underlying buffer.
func (b *Buffer) WriteByte(v byte) error {
	_, err := b.Write([]byte{v})
	return err
}

// WriteUint8 writes a uint8 to the underlying buffer.
func (b *Buffer) WriteUint8(v uint8) error {
	_, err := b.Write(Le.WriteUint8(v))
	return err
}

// WriteLUint16 writes a uint16 (little-endian) to the underlying buffer.
func (b *Buffer) WriteLUint16(v uint16) error {
	_, err := b.Write(Le.WriteUint16(v))
	return err
}

// WriteBUint16 writes a uint16 (big-endian) to the underlying buffer.
func (b *Buffer) WriteBUint16(v uint16) error {
	_, err := b.Write(Be.WriteUint16(v))
	return err
}

// WriteLUint32 writes a uint32 (little-endian) to the underlying buffer.
func (b *Buffer) WriteLUint32(v uint32) error {
	_, err := b.Write(Le.WriteUint32(v))
	return err
}

// WriteBUint32 writes a uint32 (big-endian) to the underlying buffer.
func (b *Buffer) WriteBUint32(v uint32) error {
	_, err := b.Write(Be.WriteUint32(v))
	return err
}

// WriteLUint64 writes a uint64 (little-endian) to the underlying buffer.
func (b *Buffer) WriteLUint64(v uint64) error {
	_, err := b.Write(Le.WriteUint64(v))
	return err
}

// WriteBUint64 writes a uint64 (big-endian) to the underlying buffer.
func (b *Buffer) WriteBUint64(v uint64) error {
	_, err := b.Write(Be.WriteUint64(v))
	return err
}

// WriteLFloat32 writes a float32 (little-endian) to the underlying buffer.
func (b *Buffer) WriteLFloat32(v float32) error {
	_, err := b.Write(Le.WriteFloat32(v))
	return err
}

// WriteBFloat32 writes a float32 (big-endian) to the underlying buffer.
func (b *Buffer) WriteBFloat32(v float32) error {
	_, err := b.Write(Be.WriteFloat32(v))
	return err
}

// WriteLFloat64 writes a float64 (little-endian) to the underlying buffer.
func (b *Buffer) WriteLFloat64(v float64) error {
	_, err := b.Write(Le.WriteFloat64(v))
	return err
}

// WriteBFloat64 writes a float64 (big-endian) to the underlying buffer.
func (b *Buffer) WriteBFloat64(v float64) error {
	_, err := b.Write(Be.WriteFloat64(v))
	return err
}

// WriteString writes a string to the underlying buffer.
func (b *Buffer) WriteString(v string) error {
	if b.err != nil {
		return b.err
	}
	_, err := b.buf.WriteString(v)
	return err
}

// Bytes returns the bytes buffered.
func (b *Buffer) Bytes() []byte {
	return b.buf.Bytes()
}

// Read reads into p bytes and returns the number of bytes read and an error, if any.
// Will NOP if the buffer has stored an error.
func (b *Buffer) Read(p []byte) (n int, err error) {
	// Return early
	if b.err != nil {
		return 0, b.err
	}

	n, err = b.buf.Read(p)
	if err != nil {
		b.err = err
		return n, err
	}
	return n, nil
}

// ReadByte reads a single byte from the underlying buffer.
func (b *Buffer) ReadByte() (byte, error) {
	if b.err != nil {
		return 0, b.err
	}

	return b.buf.ReadByte()
}

// ReadUint8 reads a single uint8 (byte) from the underlying buffer.
func (b *Buffer) ReadUint8() (uint8, error) {
	return Le.ReadUint8(b)
}

// ReadLUint16 reads a uint16 (little-endian) from the underlying buffer.
func (b *Buffer) ReadLUint16() (uint16, error) {
	return Le.ReadUint16(b)
}

// ReadBUint16 reads a uint16 (big-endian) from the underlying buffer.
func (b *Buffer) ReadBUint16() (uint16, error) {
	return Be.ReadUint16(b)
}

// ReadLUint32 reads a uint32 (little-endian) from the underlying buffer.
func (b *Buffer) ReadLUint32() (uint32, error) {
	return Le.ReadUint32(b)
}

// ReadBUint32 reads a uint32 (big-endian) from the underlying buffer.
func (b *Buffer) ReadBUint32() (uint32, error) {
	return Be.ReadUint32(b)
}

// ReadLUint64 reads a uint64 (little-endian) from the underlying buffer.
func (b *Buffer) ReadLUint64() (uint64, error) {
	return Le.ReadUint64(b)
}

// ReadBUint64 reads a uint64 (big-endian) from the underlying buffer.
func (b *Buffer) ReadBUint64() (uint64, error) {
	return Be.ReadUint64(b)
}

// ReadLFloat32 reads a float32 (little-endian) from the underlying buffer.
func (b *Buffer) ReadLFloat32() (float32, error) {
	return Le.ReadFloat32(b)
}

// ReadBFloat32 reads a float32 (big-endian) from the underlying buffer.
func (b *Buffer) ReadBFloat32() (float32, error) {
	return Be.ReadFloat32(b)
}

// ReadLFloat64 reads a float64 (little-endian) from the underlying buffer.
func (b *Buffer) ReadLFloat64() (float64, error) {
	return Le.ReadFloat64(b)
}

// ReadBFloat64 reads a float64 (big-endian) from the underlying buffer.
func (b *Buffer) ReadBFloat64() (float64, error) {
	return Be.ReadFloat64(b)
}
