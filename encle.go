package bitflux

import (
	"bytes"
	"encoding"
	"io"
	"math"
)

// EncLE is a little-endian binary encoder that writes data to an io.Writer.
// It tracks the number of bytes written and any errors that occur during encoding.
type EncLE struct {
	W   io.Writer // The underlying writer to encode data to
	N   int64     // Number of bytes written
	Err error     // First error encountered during encoding
}

// NewEncLE creates a new little-endian encoder that writes to the provided io.Writer.
func NewEncLE(w io.Writer) *EncLE { return &EncLE{W: w} }

// Convenience constructors for common use cases

// NewEncLEFromBytes creates an encoder that writes to a buffer initialized with existing data
func NewEncLEFromBytes(data []byte) *EncLE {
	return NewEncLE(bytes.NewBuffer(data))
}

// NewEncLEBuffer creates an encoder that writes to a new bytes.Buffer
func NewEncLEBuffer() *EncLE {
	return NewEncLE(&bytes.Buffer{})
}

// NewEncLEWithCapacity creates an encoder that writes to a buffer with pre-allocated capacity
func NewEncLEWithCapacity(cap int) *EncLE {
	return NewEncLE(bytes.NewBuffer(make([]byte, 0, cap)))
}

// push writes the provided byte slice to the underlying writer.
// It handles partial writes and tracks the number of bytes written and any errors.
func (e *EncLE) push(p []byte) {
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

// U8 encodes a uint8 value in little-endian format.
func (e *EncLE) U8(v uint8) {
	var b [1]byte
	b[0] = byte(v)
	e.push(b[:])
}

// U16 encodes a uint16 value in little-endian format.
func (e *EncLE) U16(v uint16) {
	var b [2]byte
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	e.push(b[:])
}

// U32 encodes a uint32 value in little-endian format.
func (e *EncLE) U32(v uint32) {
	var b [4]byte
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	e.push(b[:])
}

// U64 encodes a uint64 value in little-endian format.
func (e *EncLE) U64(v uint64) {
	var b [8]byte
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	b[4] = byte(v >> 32)
	b[5] = byte(v >> 40)
	b[6] = byte(v >> 48)
	b[7] = byte(v >> 56)
	e.push(b[:])
}

// I8 encodes an int8 value in little-endian format.
func (e *EncLE) I8(v int8) {
	e.U8(uint8(v))
}

// I16 encodes an int16 value in little-endian format.
func (e *EncLE) I16(v int16) {
	e.U16(uint16(v))
}

// I32 encodes an int32 value in little-endian format.
func (e *EncLE) I32(v int32) {
	e.U32(uint32(v))
}

// I64 encodes an int64 value in little-endian format.
func (e *EncLE) I64(v int64) {
	e.U64(uint64(v))
}

// F32 encodes a float32 value in little-endian format using IEEE 754 representation.
func (e *EncLE) F32(v float32) {
	e.U32(math.Float32bits(v))
}

// F64 encodes a float64 value in little-endian format using IEEE 754 representation.
func (e *EncLE) F64(v float64) {
	e.U64(math.Float64bits(v))
}

// Write writes raw bytes to the encoder.
func (e *EncLE) Write(p []byte) {
	e.push(p)
}

// To calls WriteTo on the provided WriterTo and updates the encoder's byte count and error state.
func (e *EncLE) To(w io.WriterTo) {
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
func (e *EncLE) Marshal(m encoding.BinaryMarshaler) {
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
