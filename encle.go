package bitflux

import (
	"encoding"
	"io"
	"math"
)

type EncLE struct {
	W   io.Writer
	N   int64
	Err error
}

func NewEncLE(w io.Writer) *EncLE { return &EncLE{W: w} }

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

func (e *EncLE) U8(v uint8) {
	var b [1]byte
	b[0] = byte(v)
	e.push(b[:])
}

func (e *EncLE) U16(v uint16) {
	var b [2]byte
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	e.push(b[:])
}

func (e *EncLE) U32(v uint32) {
	var b [4]byte
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
	e.push(b[:])
}

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

func (e *EncLE) I8(v int8) {
	e.U8(uint8(v))
}

func (e *EncLE) I16(v int16) {
	e.U16(uint16(v))
}

func (e *EncLE) I32(v int32) {
	e.U32(uint32(v))
}

func (e *EncLE) I64(v int64) {
	e.U64(uint64(v))
}

func (e *EncLE) F32(v float32) {
	e.U32(math.Float32bits(v))
}

func (e *EncLE) F64(v float64) {
	e.U64(math.Float64bits(v))
}

func (e *EncLE) Write(p []byte) {
	e.push(p)
}

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
