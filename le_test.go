package bitflux

import (
	"bytes"
	"math"
	"testing"
)

func TestLeWriteUint8(t *testing.T) {
	inputs := []uint8{0, 1, 2, 3, 255}
	expects := []byte{0x00, 0x01, 0x02, 0x03, 0xff}
	for i := range inputs {
		result := Le.WriteUint8(inputs[i])
		if expects[i] != result[0] {
			t.Fatalf("Expected %x, got %x", expects[i], result)
		}
	}
}

func TestLeWriteUint16(t *testing.T) {
	tests := []struct {
		name      string
		input     uint16
		want      []byte
		expectErr bool
	}{
		{"zero", 0, []byte{0x00, 0x00}, false},
		{"one", 1, []byte{0x01, 0x00}, false},
		{"ten", 10, []byte{0x0A, 0x00}, false},
		{"256", 256, []byte{0x00, 0x01}, false},
		{"1000", 1000, []byte{0xe8, 0x03}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Le.WriteUint16(tt.input)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func TestLeWriteUint32(t *testing.T) {
	tests := []struct {
		name      string
		input     uint32
		want      []byte
		expectErr bool
	}{
		{"zero", 0, []byte{0x00, 0x00, 0x00, 0x00}, false},
		{"one", 1, []byte{0x01, 0x00, 0x00, 0x00}, false},
		{"max", 4_294_967_295, []byte{0xff, 0xff, 0xff, 0xff}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Le.WriteUint32(tt.input)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func TestLeWriteUint64(t *testing.T) {
	tests := []struct {
		name      string
		input     uint64
		want      []byte
		expectErr bool
	}{
		{"zero", 0, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, false},
		{"one", 1, []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, false},
		{"max", math.MaxUint64, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Le.WriteUint64(tt.input)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func TestLeWriteInt16(t *testing.T) {
	tests := []struct {
		name      string
		input     int16
		want      []byte
		expectErr bool
	}{
		{"zero", 0, []byte{0x00, 0x00}, false},
		{"one", 1, []byte{0x01, 0x00}, false},
		{"max", math.MaxInt16, []byte{0xff, 0x7f}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Le.WriteInt16(tt.input)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func TestLeWriteInt32(t *testing.T) {
	tests := []struct {
		name      string
		input     int32
		want      []byte
		expectErr bool
	}{
		{"zero", 0, []byte{0x00, 0x00, 0x00, 0x00}, false},
		{"one", 1, []byte{0x01, 0x00, 0x00, 0x00}, false},
		{"max", math.MaxInt32, []byte{0xff, 0xff, 0xff, 0x7f}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Le.WriteInt32(tt.input)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("got=%v, want=%v", got, tt.want)
			}
		})
	}
}

func TestLeWriteInt64(t *testing.T) {
	tests := []struct {
		name      string
		input     int64
		want      []byte
		expectErr bool
	}{
		{"zero", 0, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, false},
		{"one", 1, []byte{0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}, false},
		{"max", math.MaxInt64, []byte{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Le.WriteInt64(tt.input)
			if !bytes.Equal(got, tt.want) {
				t.Errorf("got=%v, want=%v", got, tt.want)
			}
		})
	}
}
