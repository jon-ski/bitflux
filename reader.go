package bitflux

import (
	"io"
)

type Reader struct {
	r   io.Reader
	err error
}

func NewReader(r io.Reader) *Reader {
	return &Reader{r: r}
}

func (r *Reader) Read(p []byte) (n int, err error) {
	// fail/return early if error already in reader
	// we do not want to continue reading after an error regardless
	// of new calls to `Read`
	if r.err != nil {
		return 0, r.err
	}

	// Continue as normal
	n, err = r.r.Read(p)
	if err != nil {
		r.err = err
	}
	return n, err
}

func (r *Reader) ReadByte() (byte, error) {
	buf := make([]byte, 1)
	_, err := r.Read(buf)
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}

func (r *Reader) ReadUint8() (uint8, error) {
	return r.ReadByte()
}

func (r *Reader) ReadLUint16() (uint16, error) {
	return Le.ReadUint16(r)
}

func (r *Reader) ReadLUint32() (uint32, error) {
	return Le.ReadUint32(r)
}

func (r *Reader) ReadLUint64() (uint64, error) {
	return Le.ReadUint64(r)
}

func (r *Reader) ReadLFloat32() (float32, error) {
	return Le.ReadFloat32(r)
}

func (r *Reader) ReadLFloat64() (float64, error) {
	return Le.ReadFloat64(r)
}

func (r *Reader) ReadBUint16() (uint16, error) {
	return Be.ReadUint16(r)
}

func (r *Reader) ReadBUint32() (uint32, error) {
	return Be.ReadUint32(r)
}

func (r *Reader) ReadBUint64() (uint64, error) {
	return Be.ReadUint64(r)
}

func (r *Reader) ReadBFloat32() (float32, error) {
	return Be.ReadFloat32(r)
}

func (r *Reader) ReadBFloat64() (float64, error) {
	return Be.ReadFloat64(r)
}
