package bitflux

import (
	"bytes"
	"encoding"
	"io"
	"math"
)

// EncBE is a big-endian binary encoder that writes data to an io.Writer.
// It tracks the number of bytes written and any errors that occur during encoding.
type EncBE struct {
	W   io.Writer // The underlying writer to encode data to
	N   int64     // Number of bytes written
	Err error     // First error encountered during encoding
}

// NewEncBE creates a new big-endian encoder that writes to the provided io.Writer.
func NewEncBE(w io.Writer) *EncBE { return &EncBE{W: w} }

// Convenience constructors for common use cases

// NewEncBEFromBytes creates an encoder that writes to a buffer initialized with existing data
func NewEncBEFromBytes(data []byte) *EncBE {
	return NewEncBE(bytes.NewBuffer(data))
}

// NewEncBEBuffer creates an encoder that writes to a new bytes.Buffer
func NewEncBEBuffer() *EncBE {
	return NewEncBE(&bytes.Buffer{})
}

// NewEncBEWithCapacity creates an encoder that writes to a buffer with pre-allocated capacity
func NewEncBEWithCapacity(cap int) *EncBE {
	return NewEncBE(bytes.NewBuffer(make([]byte, 0, cap)))
}

// push writes the provided byte slice to the underlying writer.
// It handles partial writes and tracks the number of bytes written and any errors.
func (e *EncBE) push(p []byte) {
	if e.Err != nil || len(p) == 0 {
		return
	}
	for off := 0; off < len(p); {
		n, err := e.W.Write(p[off:])
		e.N += int64(n)
		off += n
		if err != nil {
			e.Err = err
			return
		}
		if n == 0 { // defensive: writer made no progress
			e.Err = io.ErrShortWrite
			return
		}
	}
}

// U8 encodes a uint8 value in big-endian format.
func (e *EncBE) U8(v uint8) {
	var b [1]byte
	b[0] = byte(v)
	e.push(b[:])
}

// U16 encodes a uint16 value in big-endian format.
func (e *EncBE) U16(v uint16) {
	var b [2]byte
	b[1] = byte(v)
	b[0] = byte(v >> 8)
	e.push(b[:])
}

// U32 encodes a uint32 value in big-endian format.
func (e *EncBE) U32(v uint32) {
	var b [4]byte
	b[3] = byte(v)
	b[2] = byte(v >> 8)
	b[1] = byte(v >> 16)
	b[0] = byte(v >> 24)
	e.push(b[:])
}

// U64 encodes a uint64 value in big-endian format.
func (e *EncBE) U64(v uint64) {
	var b [8]byte
	b[7] = byte(v)
	b[6] = byte(v >> 8)
	b[5] = byte(v >> 16)
	b[4] = byte(v >> 24)
	b[3] = byte(v >> 32)
	b[2] = byte(v >> 40)
	b[1] = byte(v >> 48)
	b[0] = byte(v >> 56)
	e.push(b[:])
}

// I8 encodes an int8 value in big-endian format.
func (e *EncBE) I8(v int8) {
	e.U8(uint8(v))
}

// I16 encodes an int16 value in big-endian format.
func (e *EncBE) I16(v int16) {
	e.U16(uint16(v))
}

// I32 encodes an int32 value in big-endian format.
func (e *EncBE) I32(v int32) {
	e.U32(uint32(v))
}

// I64 encodes an int64 value in big-endian format.
func (e *EncBE) I64(v int64) {
	e.U64(uint64(v))
}

// F32 encodes a float32 value in big-endian format using IEEE 754 representation.
func (e *EncBE) F32(v float32) {
	e.U32(math.Float32bits(v))
}

// F64 encodes a float64 value in big-endian format using IEEE 754 representation.
func (e *EncBE) F64(v float64) {
	e.U64(math.Float64bits(v))
}

// Write writes raw bytes to the encoder.
func (e *EncBE) Write(p []byte) {
	e.push(p)
}

// To calls WriteTo on the provided WriterTo and updates the encoder's byte count and error state.
func (e *EncBE) To(w io.WriterTo) {
	if e.Err != nil {
		return
	}
	n, err := w.WriteTo(e.W)
	e.N += n
	if err != nil {
		e.Err = err
	}
}

// Marshal calls MarshalBinary on the provided BinaryMarshaler and writes the result to the encoder.
// It updates the encoder's byte count and error state.
func (e *EncBE) Marshal(m encoding.BinaryMarshaler) {
	if e.Err != nil {
		return
	}
	buf, err := m.MarshalBinary()
	if err != nil {
		e.Err = err
		return
	}
	e.push(buf)
}
