package bitflux

import (
	"encoding"
	"io"
	"math"
)

type DecLE struct {
	R   io.Reader
	N   int64
	Err error
}

func NewDecLE(r io.Reader) *DecLE { return &DecLE{R: r} }

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

func (d *DecLE) U8() uint8 {
	var b [1]byte
	d.pull(b[:])
	return b[0]
}

func (d *DecLE) U16() uint16 {
	var b [2]byte
	d.pull(b[:])
	return uint16(b[0]) | uint16(b[1])<<8
}

func (d *DecLE) U32() uint32 {
	var b [4]byte
	d.pull(b[:])
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

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

func (d *DecLE) I8() int8 {
	return int8(d.U8())
}

func (d *DecLE) I16() int16 {
	return int16(d.U16())
}

func (d *DecLE) I32() int32 {
	return int32(d.U32())
}

func (d *DecLE) I64() int64 {
	return int64(d.U64())
}

func (d *DecLE) F32() float32 { return math.Float32frombits(d.U32()) }
func (d *DecLE) F64() float64 { return math.Float64frombits(d.U64()) }

func (d *DecLE) Bytes(n int) []byte {
	if n <= 0 {
		return nil
	}
	buf := make([]byte, n)
	d.pull(buf)
	return buf
}

// Skip discards n bytes
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

// FromRF reads via x.ReadFrom and adds to N.
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

// Unmarshal reads n bytes and calls UnmarshalBinary.
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
