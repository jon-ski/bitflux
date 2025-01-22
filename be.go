package bitflux

import (
	"io"
	"math"
)

type be struct{}

func (c be) WriteUint8(v uint8) []byte {
	return []byte{
		byte(v),
	}
}

func (c be) WriteUint16(v uint16) []byte {
	return []byte{
		byte(v >> 8),
		byte(v),
	}
}

func (c be) WriteUint32(v uint32) []byte {
	return []byte{
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
}

func (c be) WriteUint64(v uint64) []byte {
	return []byte{
		byte(v >> 56),
		byte(v >> 48),
		byte(v >> 40),
		byte(v >> 32),
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
}

func (c be) WriteInt8(v int8) []byte {
	return []byte{
		byte(v),
	}
}

func (c be) WriteInt16(v int16) []byte {
	return []byte{
		byte(v >> 8),
		byte(v),
	}
}

func (c be) WriteInt32(v int32) []byte {
	return []byte{
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
}

func (c be) WriteInt64(v int64) []byte {
	return []byte{
		byte(v >> 56),
		byte(v >> 48),
		byte(v >> 40),
		byte(v >> 32),
		byte(v >> 24),
		byte(v >> 16),
		byte(v >> 8),
		byte(v),
	}
}

func (c be) WriteFloat32(v float32) []byte {
	return c.WriteUint32(math.Float32bits(v))
}

func (c be) WriteFloat64(v float64) []byte {
	return c.WriteUint64(math.Float64bits(v))
}

// ====================================================================
// Read Functions
// ====================================================================

func (c be) ReadUint8(r io.Reader) (uint8, error) {
	buf := make([]byte, 1)
	_, err := r.Read(buf)
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}

func (c be) ReadUint16(r io.Reader) (uint16, error) {
	buf := make([]byte, 2)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (uint16(buf[0])<<8 |
		uint16(buf[1])), nil
}

func (c be) ReadUint32(r io.Reader) (uint32, error) {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (uint32(buf[0])<<24 |
		uint32(buf[1])<<16 |
		uint32(buf[2])<<8 |
		uint32(buf[3])), nil
}

func (c be) ReadUint64(r io.Reader) (uint64, error) {
	buf := make([]byte, 8)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (uint64(buf[0])<<56 |
		uint64(buf[1])<<48 |
		uint64(buf[2])<<40 |
		uint64(buf[3])<<32 |
		uint64(buf[4])<<24 |
		uint64(buf[5])<<16 |
		uint64(buf[6])<<8 |
		uint64(buf[7])), nil
}

func (c be) ReadInt8(r io.Reader) (int8, error) {
	buf := make([]byte, 1)
	_, err := r.Read(buf)
	if err != nil {
		return 0, err
	}
	return int8(buf[0]), nil
}

func (c be) ReadInt16(r io.Reader) (int16, error) {
	buf := make([]byte, 2)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (int16(buf[0])<<8 |
		int16(buf[1])), nil
}

func (c be) ReadInt32(r io.Reader) (int32, error) {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (int32(buf[0])<<24 |
		int32(buf[1])<<16 |
		int32(buf[2])<<8 |
		int32(buf[3])), nil
}

func (c be) ReadInt64(r io.Reader) (int64, error) {
	buf := make([]byte, 8)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (int64(buf[0])<<56 |
		int64(buf[1])<<48 |
		int64(buf[2])<<40 |
		int64(buf[3])<<32 |
		int64(buf[4])<<24 |
		int64(buf[5])<<16 |
		int64(buf[6])<<8 |
		int64(buf[7])), nil
}

func (c be) ReadFloat32(r io.Reader) (float32, error) {
	buf, err := c.ReadUint32(r)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(buf), nil
}

func (c be) ReadFloat64(r io.Reader) (float64, error) {
	buf, err := c.ReadUint64(r)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(buf), nil
}
