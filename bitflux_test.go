package bitflux

import (
	"bytes"
	"testing"
)

func TestBufferGeneral(t *testing.T) {
	var buf Buffer
	var err error
	err = buf.WriteLUint16(10)
	if err != nil {
		t.Fatalf("WriteLUint16 error: %v", err)
	}

	err = buf.WriteBUint32(12345)
	if err != nil {
		t.Fatalf("WriteBUint32 error: %v", err)
	}

	if buf.Err() != nil {
		t.Fatalf("Unexpected buf.Err(): %v", buf.Err())
	}

	// Check for correctness of result
	result := buf.Bytes()
	expected := []byte{0x0A, 0x00, 0x00, 0x00, 0x30, 0x39}
	if !bytes.Equal(result, expected) {
		t.Fatalf("got=%v, want=%v", result, expected)
	}
}
