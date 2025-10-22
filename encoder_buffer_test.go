package bitflux

import (
	"bytes"
	"testing"
)

func TestEncLEConvenienceConstructors(t *testing.T) {
	// Test NewEncLEBuffer
	enc := NewEncLEBuffer()
	enc.U32(0x12345678)
	enc.U16(0x9ABC)
	enc.U8(0xDE)

	// Get the underlying bytes.Buffer to access the data
	if buf, ok := enc.W.(*bytes.Buffer); ok {
		result := buf.Bytes()
		expected := []byte{0x78, 0x56, 0x34, 0x12, 0xBC, 0x9A, 0xDE}

		if !bytes.Equal(result, expected) {
			t.Errorf("NewEncLEBuffer: got=%v, want=%v", result, expected)
		}
	} else {
		t.Error("Expected underlying writer to be *bytes.Buffer")
	}
}

func TestEncLEFromBytes(t *testing.T) {
	// Test NewEncLEFromBytes
	initialData := []byte{0xFF, 0xEE}
	enc := NewEncLEFromBytes(initialData)
	enc.U32(0x12345678)

	if buf, ok := enc.W.(*bytes.Buffer); ok {
		result := buf.Bytes()
		expected := []byte{0xFF, 0xEE, 0x78, 0x56, 0x34, 0x12}

		if !bytes.Equal(result, expected) {
			t.Errorf("NewEncLEFromBytes: got=%v, want=%v", result, expected)
		}
	} else {
		t.Error("Expected underlying writer to be *bytes.Buffer")
	}
}

func TestEncLEWithCapacity(t *testing.T) {
	// Test NewEncLEWithCapacity
	enc := NewEncLEWithCapacity(16)
	enc.U32(0x12345678)
	enc.U16(0x9ABC)

	if buf, ok := enc.W.(*bytes.Buffer); ok {
		result := buf.Bytes()
		expected := []byte{0x78, 0x56, 0x34, 0x12, 0xBC, 0x9A}

		if !bytes.Equal(result, expected) {
			t.Errorf("NewEncLEWithCapacity: got=%v, want=%v", result, expected)
		}

		// Check that capacity was pre-allocated
		if buf.Cap() < 16 {
			t.Errorf("Expected capacity >= 16, got %d", buf.Cap())
		}
	} else {
		t.Error("Expected underlying writer to be *bytes.Buffer")
	}
}

func TestEncBEConvenienceConstructors(t *testing.T) {
	// Test NewEncBEBuffer
	enc := NewEncBEBuffer()
	enc.U32(0x12345678)
	enc.U16(0x9ABC)
	enc.U8(0xDE)

	if buf, ok := enc.W.(*bytes.Buffer); ok {
		result := buf.Bytes()
		expected := []byte{0x12, 0x34, 0x56, 0x78, 0x9A, 0xBC, 0xDE}

		if !bytes.Equal(result, expected) {
			t.Errorf("NewEncBEBuffer: got=%v, want=%v", result, expected)
		}
	} else {
		t.Error("Expected underlying writer to be *bytes.Buffer")
	}
}

// Benchmark comparison: old way vs new way
func BenchmarkEncLEOldWay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		enc := NewEncLE(&buf)
		enc.U32(0x12345678)
		enc.U16(0x9ABC)
		enc.U8(0xDE)
		_ = buf.Bytes()
	}
}

func BenchmarkEncLENewWay(b *testing.B) {
	for i := 0; i < b.N; i++ {
		enc := NewEncLEBuffer()
		enc.U32(0x12345678)
		enc.U16(0x9ABC)
		enc.U8(0xDE)
		if buf, ok := enc.W.(*bytes.Buffer); ok {
			_ = buf.Bytes()
		}
	}
}

func BenchmarkEncLEWithCapacity(b *testing.B) {
	for i := 0; i < b.N; i++ {
		enc := NewEncLEWithCapacity(16)
		enc.U32(0x12345678)
		enc.U16(0x9ABC)
		enc.U8(0xDE)
		if buf, ok := enc.W.(*bytes.Buffer); ok {
			_ = buf.Bytes()
		}
	}
}
