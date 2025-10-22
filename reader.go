package bitflux

import (
	"io"
)

// Deprecated: Use DecLE/DecBE instead. Reader will be removed in a future version.
// Reader wraps an io.Reader and provides serialization methods on top of it.
type Reader struct {
	r   io.Reader
	err error
}

// Deprecated: Use NewDecLE/NewDecBE instead. NewReader will be removed in a future version.
// NewReader creates a new Reader that wraps the provided io.Reader.
func NewReader(r io.Reader) *Reader {
	return &Reader{r: r}
}

// Deprecated: Use DecLE/DecBE instead. Read will be removed in a future version.
// Read reads into p bytes and returns the number of bytes read and an error, if any.
// Will NOP if the reader has stored an error.
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

// Deprecated: Use DecLE/DecBE instead. ReadByte will be removed in a future version.
// ReadByte reads a single byte from the underlying reader.
func (r *Reader) ReadByte() (byte, error) {
	buf := make([]byte, 1)
	_, err := r.Read(buf)
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}

// Deprecated: Use DecLE/DecBE instead. ReadUint8 will be removed in a future version.
// ReadUint8 reads a single uint8 (byte) from the underlying reader.
func (r *Reader) ReadUint8() (uint8, error) {
	return r.ReadByte()
}

// Deprecated: Use DecLE/DecBE instead. ReadLUint16 will be removed in a future version.
// ReadLUint16 reads a uint16 (little-endian) from the underlying reader.
func (r *Reader) ReadLUint16() (uint16, error) {
	return Le.ReadUint16(r)
}

// Deprecated: Use DecLE/DecBE instead. ReadLUint32 will be removed in a future version.
// ReadLUint32 reads a uint32 (little-endian) from the underlying reader.
func (r *Reader) ReadLUint32() (uint32, error) {
	return Le.ReadUint32(r)
}

// Deprecated: Use DecLE/DecBE instead. ReadLUint64 will be removed in a future version.
// ReadLUint64 reads a uint64 (little-endian) from the underlying reader.
func (r *Reader) ReadLUint64() (uint64, error) {
	return Le.ReadUint64(r)
}

// Deprecated: Use DecLE/DecBE instead. ReadLFloat32 will be removed in a future version.
// ReadLFloat32 reads a float32 (little-endian) from the underlying reader.
func (r *Reader) ReadLFloat32() (float32, error) {
	return Le.ReadFloat32(r)
}

// Deprecated: Use DecLE/DecBE instead. ReadLFloat64 will be removed in a future version.
// ReadLFloat64 reads a float64 (little-endian) from the underlying reader.
func (r *Reader) ReadLFloat64() (float64, error) {
	return Le.ReadFloat64(r)
}

// Deprecated: Use DecLE/DecBE instead. ReadBUint16 will be removed in a future version.
// ReadBUint16 reads a uint16 (big-endian) from the underlying reader.
func (r *Reader) ReadBUint16() (uint16, error) {
	return Be.ReadUint16(r)
}

// Deprecated: Use DecLE/DecBE instead. ReadBUint32 will be removed in a future version.
// ReadBUint32 reads a uint32 (big-endian) from the underlying reader.
func (r *Reader) ReadBUint32() (uint32, error) {
	return Be.ReadUint32(r)
}

// Deprecated: Use DecLE/DecBE instead. ReadBUint64 will be removed in a future version.
// ReadBUint64 reads a uint64 (big-endian) from the underlying reader.
func (r *Reader) ReadBUint64() (uint64, error) {
	return Be.ReadUint64(r)
}

// Deprecated: Use DecLE/DecBE instead. ReadBFloat32 will be removed in a future version.
// ReadBFloat32 reads a float32 (big-endian) from the underlying reader.
func (r *Reader) ReadBFloat32() (float32, error) {
	return Be.ReadFloat32(r)
}

// Deprecated: Use DecLE/DecBE instead. ReadBFloat64 will be removed in a future version.
// ReadBFloat64 reads a float64 (big-endian) from the underlying reader.
func (r *Reader) ReadBFloat64() (float64, error) {
	return Be.ReadFloat64(r)
}
