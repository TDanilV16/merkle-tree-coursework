package hasher

import (
	_ "crypto/sha256"
	"os"
	"testing"

	"github.com/TDanilV16/merkle-tree-coursework/pkg/hasher"
)

func TestSha256Hasher_Hash(t *testing.T) {
	h := hasher.NewSha256Hasher()

	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{
			name:     "Empty input",
			input:    []byte{},
			expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		{
			name:     "abc string",
			input:    []byte("abc"),
			expected: "ba7816bf8f01cfea414140de5dae2223b00361a396177a9cb410ff61f20015ad",
		},
		{
			name:     "nil",
			input:    nil,
			expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hash := h.Hash(tt.input)

			if hash == nil {
				t.Error("Hash() returned nil")
			}

			if hash.Size() != 32 {
				t.Errorf("SHA256 size must be 32, got %d", hash.Size())
			}

			if hash.String() != tt.expected {
				t.Errorf("SHA256 string must be %s, got %s", tt.expected, hash.String())
			}

			if !hash.Equal(h.Hash(tt.input)) {
				t.Error("SHA256 hash must be determined")
			}
		})
	}
}

func TestSha256Hasher_HashFile(t *testing.T) {
	h := hasher.NewSha256Hasher()

	tmpFile, err := os.CreateTemp("", "hasher_test.txt")
	if err != nil {
		t.Fatalf("Failed to create file: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	testContent := []byte("test content for file hashing\n")
	if _, err := tmpFile.Write(testContent); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	tmpFile.Close()

	fileHash, err := h.HashFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to hash file: %v", err)
	}

	expectedHash := h.Hash(testContent)

	if !fileHash.Equal(expectedHash) {
		t.Errorf("HashFile() = %s, expected %s", fileHash.String(), expectedHash.String())
	}

	_, err = h.HashFile("non-existing-file")
	if err == nil {
		t.Error("HashFile() must return error for non-existing file")
	}
}

func TestSha256Hasher_VerifyHash(t *testing.T) {
	h := hasher.NewSha256Hasher()

	data := []byte("Data to verify")
	correctHash := h.Hash(data)

	wrongData := []byte("wrong data")

	if !h.VerifyHash(correctHash, data) {
		t.Error("VerifyHash() must return true for same data as hash")
	}

	if h.VerifyHash(correctHash, wrongData) {
		t.Error("VerifyHash() must return false for different data")
	}

	wrongSizeHash := hasher.NewHash([]byte{1, 2, 3})
	if h.VerifyHash(wrongSizeHash, data) {
		t.Error("VerifyHash() must return false for hash of wrong size")
	}
}

func TestSha256Hasher_HashConcat(t *testing.T) {
	h := hasher.NewSha256Hasher()

	left := h.Hash([]byte("left"))
	right := h.Hash([]byte("right"))

	parent, err := h.HashConcat(left, right)
	if err != nil || parent.Size() != 32 {
		t.Errorf("HashConcat() must return 32 hashes got %d", parent.Size())
	}

	parent2, _ := h.HashConcat(left, right)
	if err != nil || !parent.Equal(parent2) {
		t.Error("HashConcat() must be determined")
	}

	parentReversed, err := h.HashConcat(right, left)
	if err != nil || !parent.Equal(parentReversed) {
		t.Error("HashConcat() must be commutative")
	}

	_, err = h.HashConcat(left, nil)
	if err == nil {
		t.Error("HashConcat() should throw an error if one part is empty")
	}
}
