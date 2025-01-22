package bitflux

import (
	"bytes"
)

var (
	Le = new(le)
	Be = new(be)
)

type Buffer struct {
	buf *bytes.Buffer
	err error // Latest/Last Error
}

func NewBuffer(buf []byte) *Buffer {
	return &Buffer{
		buf: bytes.NewBuffer(buf),
	}
}

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

func (b *Buffer) WriteByte(v byte) error {
	_, err := b.Write([]byte{v})
	return err
}

func (b *Buffer) WriteUint8(v uint8) error {
	_, err := b.Write(Le.WriteUint8(v))
	return err
}

func (b *Buffer) WriteLUint16(v uint16) error {
	_, err := b.Write(Le.WriteUint16(v))
	return err
}

func (b *Buffer) WriteBUint16(v uint16) error {
	_, err := b.Write(Be.WriteUint16(v))
	return err
}

func (b *Buffer) WriteLUint32(v uint32) error {
	_, err := b.Write(Le.WriteUint32(v))
	return err
}

func (b *Buffer) WriteBUint32(v uint32) error {
	_, err := b.Write(Be.WriteUint32(v))
	return err
}

func (b *Buffer) WriteLUint64(v uint64) error {
	_, err := b.Write(Le.WriteUint64(v))
	return err
}

func (b *Buffer) WriteBUint64(v uint64) error {
	_, err := b.Write(Be.WriteUint64(v))
	return err
}

func (b *Buffer) WriteLFloat32(v float32) error {
	_, err := b.Write(Le.WriteFloat32(v))
	return err
}

func (b *Buffer) WriteBFloat32(v float32) error {
	_, err := b.Write(Be.WriteFloat32(v))
	return err
}

func (b *Buffer) WriteLFloat64(v float64) error {
	_, err := b.Write(Le.WriteFloat64(v))
	return err
}

func (b *Buffer) WriteBFloat64(v float64) error {
	_, err := b.Write(Be.WriteFloat64(v))
	return err
}

func (b *Buffer) WriteString(v string) error {
	if b.err != nil {
		return b.err
	}
	_, err := b.buf.WriteString(v)
	return err
}

func (b *Buffer) Bytes() []byte {
	return b.buf.Bytes()
}

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

func (b *Buffer) ReadByte() (byte, error) {
	if b.err != nil {
		return 0, b.err
	}

	return b.buf.ReadByte()
}

func (b *Buffer) ReadUint8() (uint8, error) {
	return Le.ReadUint8(b)
}

func (b *Buffer) ReadLUint16() (uint16, error) {
	return Le.ReadUint16(b)
}

func (b *Buffer) ReadBUint16() (uint16, error) {
	return Be.ReadUint16(b)
}

func (b *Buffer) ReadLUint32() (uint32, error) {
	return Le.ReadUint32(b)
}

func (b *Buffer) ReadBUint32() (uint32, error) {
	return Be.ReadUint32(b)
}

func (b *Buffer) ReadLUint64() (uint64, error) {
	return Le.ReadUint64(b)
}

func (b *Buffer) ReadBUint64() (uint64, error) {
	return Be.ReadUint64(b)
}

func (b *Buffer) ReadLFloat32() (float32, error) {
	return Le.ReadFloat32(b)
}

func (b *Buffer) ReadBFloat32() (float32, error) {
	return Be.ReadFloat32(b)
}
