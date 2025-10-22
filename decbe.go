package bitflux

import (
	"encoding"
	"io"
	"math"
)

type DecBE struct {
	R   io.Reader
	N   int64
	Err error
}

func NewDecBE(r io.Reader) *DecBE { return &DecBE{R: r} }

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

func (d *DecBE) U8() uint8 {
	var b [1]byte
	d.pull(b[:])
	return b[0]
}

func (d *DecBE) U16() uint16 {
	var b [2]byte
	d.pull(b[:])
	return uint16(b[1]) | uint16(b[0])<<8
}

func (d *DecBE) U32() uint32 {
	var b [4]byte
	d.pull(b[:])
	return uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
}

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

func (d *DecBE) I8() int8 {
	return int8(d.U8())
}

func (d *DecBE) I16() int16 {
	return int16(d.U16())
}

func (d *DecBE) I32() int32 {
	return int32(d.U32())
}

func (d *DecBE) I64() int64 {
	return int64(d.U64())
}

func (d *DecBE) F32() float32 { return math.Float32frombits(d.U32()) }
func (d *DecBE) F64() float64 { return math.Float64frombits(d.U64()) }

func (d *DecBE) Bytes(n int) []byte {
	if n <= 0 {
		return nil
	}
	buf := make([]byte, n)
	d.pull(buf)
	return buf
}

// Skip discards n bytes
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

// FromRF reads via x.ReadFrom and adds to N.
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

// Unmarshal reads n bytes and calls UnmarshalBinary.
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
