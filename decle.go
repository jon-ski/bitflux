package bitflux

import (
	"encoding"
	"io"
	"math"
)

// DecLE is a little-endian binary decoder that reads data from an io.Reader.
// It tracks the number of bytes read and any errors that occur during decoding.
type DecLE struct {
	R   io.Reader // The underlying reader to decode data from
	N   int64     // Number of bytes read
	Err error     // First error encountered during decoding
}

// NewDecLE creates a new little-endian decoder that reads from the provided io.Reader.
func NewDecLE(r io.Reader) *DecLE { return &DecLE{R: r} }

// pull reads the provided byte slice from the underlying reader using io.ReadFull.
// It tracks the number of bytes read and any errors that occur.
func (d *DecLE) pull(p []byte) {
	if d.Err != nil {
		return
	}
	n, err := io.ReadFull(d.R, p)
	d.N += int64(n)
	if err != nil {
		d.Err = err
	}
}

// U8 decodes a uint8 value from little-endian format.
func (d *DecLE) U8() uint8 {
	var b [1]byte
	d.pull(b[:])
	return b[0]
}

// U16 decodes a uint16 value from little-endian format.
func (d *DecLE) U16() uint16 {
	var b [2]byte
	d.pull(b[:])
	return uint16(b[0]) | uint16(b[1])<<8
}

// U32 decodes a uint32 value from little-endian format.
func (d *DecLE) U32() uint32 {
	var b [4]byte
	d.pull(b[:])
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

// U64 decodes a uint64 value from little-endian format.
func (d *DecLE) U64() uint64 {
	var b [8]byte
	d.pull(b[:])
	return uint64(b[0]) |
		uint64(b[1])<<8 |
		uint64(b[2])<<16 |
		uint64(b[3])<<24 |
		uint64(b[4])<<32 |
		uint64(b[5])<<40 |
		uint64(b[6])<<48 |
		uint64(b[7])<<56
}

// I8 decodes an int8 value from little-endian format.
func (d *DecLE) I8() int8 {
	return int8(d.U8())
}

// I16 decodes an int16 value from little-endian format.
func (d *DecLE) I16() int16 {
	return int16(d.U16())
}

// I32 decodes an int32 value from little-endian format.
func (d *DecLE) I32() int32 {
	return int32(d.U32())
}

// I64 decodes an int64 value from little-endian format.
func (d *DecLE) I64() int64 {
	return int64(d.U64())
}

// F32 decodes a float32 value from little-endian format using IEEE 754 representation.
func (d *DecLE) F32() float32 { return math.Float32frombits(d.U32()) }

// F64 decodes a float64 value from little-endian format using IEEE 754 representation.
func (d *DecLE) F64() float64 { return math.Float64frombits(d.U64()) }

// Bytes reads n bytes from the decoder and returns them as a byte slice.
func (d *DecLE) Bytes(n int) []byte {
	if n <= 0 {
		return nil
	}
	buf := make([]byte, n)
	d.pull(buf)
	return buf
}

// Skip discards n bytes from the decoder without storing them.
func (d *DecLE) Skip(n int) {
	if n <= 0 || d.Err != nil {
		return
	}
	var scratch [64]byte
	for n > 0 && d.Err == nil {
		k := min(n, len(scratch))
		d.pull(scratch[:k])
		n -= k
	}
}

// FromRF calls ReadFrom on the provided ReaderFrom and updates the decoder's byte count and error state.
func (d *DecLE) FromRF(r io.ReaderFrom) {
	if d.Err != nil {
		return
	}
	n, err := r.ReadFrom(d.R)
	d.N += n
	if err != nil {
		d.Err = err
	}
}

// Unmarshal reads n bytes and calls UnmarshalBinary on the provided BinaryUnmarshaler.
// It updates the decoder's error state if unmarshaling fails.
func (d *DecLE) Unmarshal(u encoding.BinaryUnmarshaler, n int) {
	if d.Err != nil {
		return
	}
	b := d.Bytes(n)
	if d.Err != nil {
		return
	}
	if err := u.UnmarshalBinary(b); err != nil {
		d.Err = err
	}
}
