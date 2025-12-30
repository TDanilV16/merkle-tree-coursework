package hasher

import (
	"bytes"
	"testing"

	"github.com/TDanilV16/merkle-tree-coursework/pkg/hasher"
)

func TestNewHash_Empty(t *testing.T) {
	hash1 := hasher.NewHash([]byte{})

	if !hash1.IsZero() {
		t.Errorf("NewHash([]byte{}) must return zero hash")
	}

	hash1String := hash1.String()

	if hash1String != "" {
		t.Errorf("NewHash([]byte{}) must return zero hash got %s", hash1String)
	}

	hash2 := hasher.NewHash(nil)
	if !hash2.IsZero() {
		t.Errorf("NewHash(nil) must return zero hash")
	}
}

func TestNewHash_FromBytes(t *testing.T) {
	data := []byte("hello world")
	hash := hasher.NewHash(data)

	if hash.Size() != len(data) {
		t.Errorf("Size() = %d, expected %d", hash.Size(), len(data))
	}

	hashBytes := hash.Bytes()
	if !bytes.Equal(data, hashBytes) {
		t.Errorf("Bytes() = %s, expected %s", hashBytes, data)
	}

	hashBytes[0] = 255
	if hash.Bytes()[0] == 255 {
		t.Errorf("Bytes() must return copy, not actual bytes")
	}
}

func TestHash_String(t *testing.T) {
	tests := []struct {
		name     string
		data     []byte
		expected string
	}{
		{
			name:     "Empty hash",
			data:     []byte{},
			expected: "",
		},
		{
			name:     "One byte",
			data:     []byte{255},
			expected: "ff",
		},
		{
			name:     "Several bytes",
			data:     []byte{1, 2, 3, 254, 255},
			expected: "010203feff",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := hasher.NewHash(tt.data)
			result := hash.String()

			if result != tt.expected {
				t.Errorf("String() = %s, expected %s", result, tt.expected)
			}
		})
	}
}

func TestHash_Equals(t *testing.T) {
	hash1 := hasher.NewHash([]byte{1, 2, 3})
	hash2 := hasher.NewHash([]byte{1, 2, 3})
	hash3 := hasher.NewHash([]byte{4, 5, 6})
	hash4 := hasher.NewHash([]byte{1, 2})

	tests := []struct {
		name      string
		leftHash  hasher.Hash
		rightHash hasher.Hash
		expected  bool
	}{
		{
			name:      "Same hashes are equal",
			leftHash:  hash1,
			rightHash: hash2,
			expected:  true,
		},
		{
			name:      "Different hashes are not equal",
			leftHash:  hash1,
			rightHash: hash3,
			expected:  false,
		},
		{
			name:      "Hashes of data of different lengths are not equal",
			leftHash:  hash1,
			rightHash: hash4,
			expected:  false,
		},
		{
			name:      "Hash of nil is not equal to non-nil hash",
			leftHash:  hash1,
			rightHash: hasher.NewHash(nil),
			expected:  false,
		},
		{
			name:      "Reflexivity",
			leftHash:  hash1,
			rightHash: hash1,
			expected:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.leftHash.Equal(tt.rightHash)
			if result != tt.expected {
				t.Errorf("Equal() = %t, expected %t", result, tt.expected)
			}
		})
	}
}
