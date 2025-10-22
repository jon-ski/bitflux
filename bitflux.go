package bitflux

import (
	"bytes"
	"encoding"
	"io"
)

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. These global variables will be removed in a future version.
var (
	Le = new(le) // Little-Endian functions
	Be = new(be) // Big-Endian functions
)

var _ io.Reader = &Buffer{}
var _ io.Writer = &Buffer{}
var _ io.ByteReader = &Buffer{}
var _ io.ByteWriter = &Buffer{}
var _ io.ReaderFrom = &Buffer{}
var _ io.WriterTo = &Buffer{}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. Buffer will be removed in a future version.
// Buffer wraps `bytes.Buffer` and provides serialization methods on top of it.
type Buffer struct {
	buf bytes.Buffer
	err error // Latest/Last Error
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. NewBuffer will be removed in a future version.
// NewBuffer creates an initialized internal buffer from `buf`.
// Useful as a reader on an existing buffer.
func NewBuffer(buf []byte) *Buffer {
	return &Buffer{
		buf: *bytes.NewBuffer(buf),
	}
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. Err will be removed in a future version.
// Err returns an error if an error has occurred on the buffer
func (b *Buffer) Err() error {
	return b.err
}

func (b *Buffer) setErr(err error) {
	if err != nil && err != io.EOF {
		b.err = err
	}
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. Write will be removed in a future version.
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

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. WriteByte will be removed in a future version.
// WriteByte writes a byte to the underlying buffer.
func (b *Buffer) WriteByte(v byte) error {
	_, err := b.Write([]byte{v})
	return err
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. WriteUint8 will be removed in a future version.
// WriteUint8 writes a uint8 to the underlying buffer.
func (b *Buffer) WriteUint8(v uint8) error {
	_, err := b.Write(Le.WriteUint8(v))
	return err
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. WriteLUint16 will be removed in a future version.
// WriteLUint16 writes a uint16 (little-endian) to the underlying buffer.
func (b *Buffer) WriteLUint16(v uint16) error {
	_, err := b.Write(Le.WriteUint16(v))
	return err
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. WriteBUint16 will be removed in a future version.
// WriteBUint16 writes a uint16 (big-endian) to the underlying buffer.
func (b *Buffer) WriteBUint16(v uint16) error {
	_, err := b.Write(Be.WriteUint16(v))
	return err
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. WriteLUint32 will be removed in a future version.
// WriteLUint32 writes a uint32 (little-endian) to the underlying buffer.
func (b *Buffer) WriteLUint32(v uint32) error {
	_, err := b.Write(Le.WriteUint32(v))
	return err
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. WriteBUint32 will be removed in a future version.
// WriteBUint32 writes a uint32 (big-endian) to the underlying buffer.
func (b *Buffer) WriteBUint32(v uint32) error {
	_, err := b.Write(Be.WriteUint32(v))
	return err
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. WriteLUint64 will be removed in a future version.
// WriteLUint64 writes a uint64 (little-endian) to the underlying buffer.
func (b *Buffer) WriteLUint64(v uint64) error {
	_, err := b.Write(Le.WriteUint64(v))
	return err
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. WriteBUint64 will be removed in a future version.
// WriteBUint64 writes a uint64 (big-endian) to the underlying buffer.
func (b *Buffer) WriteBUint64(v uint64) error {
	_, err := b.Write(Be.WriteUint64(v))
	return err
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. WriteLFloat32 will be removed in a future version.
// WriteLFloat32 writes a float32 (little-endian) to the underlying buffer.
func (b *Buffer) WriteLFloat32(v float32) error {
	_, err := b.Write(Le.WriteFloat32(v))
	return err
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. WriteBFloat32 will be removed in a future version.
// WriteBFloat32 writes a float32 (big-endian) to the underlying buffer.
func (b *Buffer) WriteBFloat32(v float32) error {
	_, err := b.Write(Be.WriteFloat32(v))
	return err
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. WriteLFloat64 will be removed in a future version.
// WriteLFloat64 writes a float64 (little-endian) to the underlying buffer.
func (b *Buffer) WriteLFloat64(v float64) error {
	_, err := b.Write(Le.WriteFloat64(v))
	return err
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. WriteBFloat64 will be removed in a future version.
// WriteBFloat64 writes a float64 (big-endian) to the underlying buffer.
func (b *Buffer) WriteBFloat64(v float64) error {
	_, err := b.Write(Be.WriteFloat64(v))
	return err
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. WriteString will be removed in a future version.
// WriteString writes a string to the underlying buffer.
func (b *Buffer) WriteString(v string) error {
	if b.err != nil {
		return b.err
	}
	_, err := b.buf.WriteString(v)
	return err
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. Bytes will be removed in a future version.
// Bytes returns the bytes buffered.
func (b *Buffer) Bytes() []byte {
	return b.buf.Bytes()
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. Read will be removed in a future version.
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

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. ReadByte will be removed in a future version.
// ReadByte reads a single byte from the underlying buffer.
func (b *Buffer) ReadByte() (byte, error) {
	if b.err != nil {
		return 0, b.err
	}

	return b.buf.ReadByte()
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. ReadUint8 will be removed in a future version.
// ReadUint8 reads a single uint8 (byte) from the underlying buffer.
func (b *Buffer) ReadUint8() (uint8, error) {
	return Le.ReadUint8(b)
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. ReadLUint16 will be removed in a future version.
// ReadLUint16 reads a uint16 (little-endian) from the underlying buffer.
func (b *Buffer) ReadLUint16() (uint16, error) {
	return Le.ReadUint16(b)
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. ReadBUint16 will be removed in a future version.
// ReadBUint16 reads a uint16 (big-endian) from the underlying buffer.
func (b *Buffer) ReadBUint16() (uint16, error) {
	return Be.ReadUint16(b)
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. ReadLUint32 will be removed in a future version.
// ReadLUint32 reads a uint32 (little-endian) from the underlying buffer.
func (b *Buffer) ReadLUint32() (uint32, error) {
	return Le.ReadUint32(b)
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. ReadBUint32 will be removed in a future version.
// ReadBUint32 reads a uint32 (big-endian) from the underlying buffer.
func (b *Buffer) ReadBUint32() (uint32, error) {
	return Be.ReadUint32(b)
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. ReadLUint64 will be removed in a future version.
// ReadLUint64 reads a uint64 (little-endian) from the underlying buffer.
func (b *Buffer) ReadLUint64() (uint64, error) {
	return Le.ReadUint64(b)
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. ReadBUint64 will be removed in a future version.
// ReadBUint64 reads a uint64 (big-endian) from the underlying buffer.
func (b *Buffer) ReadBUint64() (uint64, error) {
	return Be.ReadUint64(b)
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. ReadLFloat32 will be removed in a future version.
// ReadLFloat32 reads a float32 (little-endian) from the underlying buffer.
func (b *Buffer) ReadLFloat32() (float32, error) {
	return Le.ReadFloat32(b)
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. ReadBFloat32 will be removed in a future version.
// ReadBFloat32 reads a float32 (big-endian) from the underlying buffer.
func (b *Buffer) ReadBFloat32() (float32, error) {
	return Be.ReadFloat32(b)
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. ReadLFloat64 will be removed in a future version.
// ReadLFloat64 reads a float64 (little-endian) from the underlying buffer.
func (b *Buffer) ReadLFloat64() (float64, error) {
	return Le.ReadFloat64(b)
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. ReadBFloat64 will be removed in a future version.
// ReadBFloat64 reads a float64 (big-endian) from the underlying buffer.
func (b *Buffer) ReadBFloat64() (float64, error) {
	return Be.ReadFloat64(b)
}

// --- io.ReaderFrom / io.WriterTo ---

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. ReadFrom will be removed in a future version.
func (b *Buffer) ReadFrom(r io.Reader) (int64, error) {
	if b.err != nil {
		return 0, b.err
	}
	n, err := b.buf.ReadFrom(r)
	b.setErr(err)
	return n, err
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. WriteTo will be removed in a future version.
func (b *Buffer) WriteTo(w io.Writer) (int64, error) {
	// bytes.Buffer.WriteTo ignores b.err. Keep sticky behavior consistent.
	if b.err != nil {
		return 0, b.err
	}
	return b.buf.WriteTo(w)
}

// --- binary marshaling helpers ---

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. WriteBinary will be removed in a future version.
// WriteBinary writes raw MarshalBinary bytes. No length prefix.
func (b *Buffer) WriteBinary(m encoding.BinaryMarshaler) error {
	if b.err != nil {
		return b.err
	}
	data, err := m.MarshalBinary()
	if err != nil {
		b.setErr(err)
		return err
	}
	_, err = b.Write(data)
	return err
}

// Deprecated: Use EncLE/DecLE and EncBE/DecBE instead. ReadBinary will be removed in a future version.
// ReadBinary reads exactly n bytes and calls UnmarshalBinary.
func (b *Buffer) ReadBinary(u encoding.BinaryUnmarshaler, n int) error {
	if b.err != nil {
		return b.err
	}
	buf := make([]byte, n)
	if _, err := io.ReadFull(b, buf); err != nil {
		b.setErr(err)
		return err
	}
	return u.UnmarshalBinary(buf)
}
