package bitflux

import (
	"io"
	"math"
)

type le struct{}

func (c le) WriteUint8(v uint8) []byte {
	return []byte{
		byte(v),
	}
}

func (c le) WriteUint16(v uint16) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
	}
}

func (c le) WriteUint32(v uint32) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
		byte(v >> 24),
	}
}

func (c le) WriteUint64(v uint64) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
		byte(v >> 24),
		byte(v >> 32),
		byte(v >> 40),
		byte(v >> 48),
		byte(v >> 56),
	}
}

func (c le) WriteInt8(v int8) []byte {
	return []byte{
		byte(v),
	}
}

func (c le) WriteInt16(v int16) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
	}
}

func (c le) WriteInt32(v int32) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
		byte(v >> 24),
	}
}

func (c le) WriteInt64(v int64) []byte {
	return []byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
		byte(v >> 24),
		byte(v >> 32),
		byte(v >> 40),
		byte(v >> 48),
		byte(v >> 56),
	}
}

func (c le) WriteFloat32(v float32) []byte {
	return c.WriteUint32(math.Float32bits(v))
}

func (c le) WriteFloat64(v float64) []byte {
	return c.WriteUint64(math.Float64bits(v))
}

func (c le) ReadUint8(r io.Reader) (uint8, error) {
	buf := make([]byte, 1)
	_, err := r.Read(buf)
	if err != nil {
		return 0, err
	}
	return buf[0], nil
}

func (c le) ReadUint16(r io.Reader) (uint16, error) {
	buf := make([]byte, 2)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return uint16(buf[0]) | uint16(buf[1])<<8, nil
}

func (c le) ReadUint32(r io.Reader) (uint32, error) {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (uint32(buf[0]) |
		uint32(buf[1])<<8 |
		uint32(buf[2])<<16 |
		uint32(buf[3])<<24), nil
}

func (c le) ReadUint64(r io.Reader) (uint64, error) {
	buf := make([]byte, 8)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (uint64(buf[0]) |
		uint64(buf[1])<<8 |
		uint64(buf[2])<<16 |
		uint64(buf[3])<<24 |
		uint64(buf[4])<<32 |
		uint64(buf[5])<<40 |
		uint64(buf[6])<<48 |
		uint64(buf[7])<<56), nil
}

func (c le) ReadInt8(r io.Reader) (int8, error) {
	buf := make([]byte, 1)
	_, err := r.Read(buf)
	if err != nil {
		return 0, err
	}
	return int8(buf[0]), nil
}

func (c le) ReadInt16(r io.Reader) (int16, error) {
	buf := make([]byte, 2)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return int16(buf[0]) | int16(buf[1])<<8, nil
}

func (c le) ReadInt32(r io.Reader) (int32, error) {
	buf := make([]byte, 4)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (int32(buf[0]) |
		int32(buf[1])<<8 |
		int32(buf[2])<<16 |
		int32(buf[3])<<24), nil
}

func (c le) ReadInt64(r io.Reader) (int64, error) {
	buf := make([]byte, 8)
	if _, err := io.ReadFull(r, buf); err != nil {
		return 0, err
	}
	return (int64(buf[0]) |
		int64(buf[1])<<8 |
		int64(buf[2])<<16 |
		int64(buf[3])<<24 |
		int64(buf[4])<<32 |
		int64(buf[5])<<40 |
		int64(buf[6])<<48 |
		int64(buf[7])<<56), nil
}

func (c le) ReadFloat32(r io.Reader) (float32, error) {
	buf, err := c.ReadUint32(r)
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(buf), nil
}

func (c le) ReadFloat64(r io.Reader) (float64, error) {
	buf, err := c.ReadUint64(r)
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(buf), nil
}
