package bitflux

import (
	"encoding"
	"io"
	"math"
)

// DecBE is a big-endian binary decoder that reads data from an io.Reader.
// It tracks the number of bytes read and any errors that occur during decoding.
type DecBE struct {
	R   io.Reader // The underlying reader to decode data from
	N   int64     // Number of bytes read
	Err error     // First error encountered during decoding
}

// NewDecBE creates a new big-endian decoder that reads from the provided io.Reader.
func NewDecBE(r io.Reader) *DecBE { return &DecBE{R: r} }

// pull reads the provided byte slice from the underlying reader using io.ReadFull.
// It tracks the number of bytes read and any errors that occur.
func (d *DecBE) pull(p []byte) {
	if d.Err != nil {
		return
	}
	n, err := io.ReadFull(d.R, p)
	d.N += int64(n)
	if err != nil {
		d.Err = err
	}
}

// U8 decodes a uint8 value from big-endian format.
func (d *DecBE) U8() uint8 {
	var b [1]byte
	d.pull(b[:])
	return b[0]
}

// U16 decodes a uint16 value from big-endian format.
func (d *DecBE) U16() uint16 {
	var b [2]byte
	d.pull(b[:])
	return uint16(b[1]) | uint16(b[0])<<8
}

// U32 decodes a uint32 value from big-endian format.
func (d *DecBE) U32() uint32 {
	var b [4]byte
	d.pull(b[:])
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

// U64 decodes a uint64 value from big-endian format.
func (d *DecBE) U64() uint64 {
	var b [8]byte
	d.pull(b[:])
	return uint64(b[7]) |
		uint64(b[6])<<8 |
		uint64(b[5])<<16 |
		uint64(b[4])<<24 |
		uint64(b[3])<<32 |
		uint64(b[2])<<40 |
		uint64(b[1])<<48 |
		uint64(b[0])<<56
}

// I8 decodes an int8 value from big-endian format.
func (d *DecBE) I8() int8 {
	return int8(d.U8())
}

// I16 decodes an int16 value from big-endian format.
func (d *DecBE) I16() int16 {
	return int16(d.U16())
}

// I32 decodes an int32 value from big-endian format.
func (d *DecBE) I32() int32 {
	return int32(d.U32())
}

// I64 decodes an int64 value from big-endian format.
func (d *DecBE) I64() int64 {
	return int64(d.U64())
}

// F32 decodes a float32 value from big-endian format using IEEE 754 representation.
func (d *DecBE) F32() float32 { return math.Float32frombits(d.U32()) }

// F64 decodes a float64 value from big-endian format using IEEE 754 representation.
func (d *DecBE) F64() float64 { return math.Float64frombits(d.U64()) }

// Bytes reads n bytes from the decoder and returns them as a byte slice.
func (d *DecBE) Bytes(n int) []byte {
	if n <= 0 {
		return nil
	}
	buf := make([]byte, n)
	d.pull(buf)
	return buf
}

// Skip discards n bytes from the decoder without storing them.
func (d *DecBE) Skip(n int) {
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
func (d *DecBE) FromRF(r io.ReaderFrom) {
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
func (d *DecBE) Unmarshal(u encoding.BinaryUnmarshaler, n int) {
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
