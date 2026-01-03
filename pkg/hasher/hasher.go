package hasher

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

type Hasher interface {
	Hash(data []byte) Hash
	HashFile(filepath string) (Hash, error)
	VerifyHash(hash Hash, data []byte) bool
	HashString(str string) Hash
	HashConcat(left, right Hash) (Hash, error)
	EmptyHash() Hash
}

type Sha256Hasher struct{}

func NewSha256Hasher() *Sha256Hasher {
	return &Sha256Hasher{}
}

func (h *Sha256Hasher) Hash(data []byte) Hash {
	if data == nil {
		data = []byte{}
	}
	hash := sha256.Sum256(data)
	return NewHash(hash[:])
}

func (h *Sha256Hasher) HashFile(filepath string) (Hash, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err = io.Copy(hasher, file); err != nil {
		return nil, err
	}

	return NewHash(hasher.Sum(nil)), nil
}

func (h *Sha256Hasher) VerifyHash(expected Hash, data []byte) bool {
	actualHash := h.Hash(data)
	return actualHash.Equal(expected)
}

func (h *Sha256Hasher) HashString(str string) Hash {
	return h.Hash([]byte(str))
}

func (h *Sha256Hasher) HashConcat(left, right Hash) (Hash, error) {
	if left == nil || right == nil {
		return nil, fmt.Errorf("one hash is nil")
	}

	if left.Size() != 32 || right.Size() != 32 {
		return nil, fmt.Errorf("hashes size mismatch: left=%d, right=%d", left.Size(), right.Size())
	}

	if left.IsZero() {
		return nil, fmt.Errorf("left hash is zero: not a valid hash")
	}

	if right.IsZero() {
		return nil, fmt.Errorf("right hash is zero: not a valid hash")
	}

	leftBytes := left.Bytes()
	rightBytes := right.Bytes()

	if bytes.Compare(leftBytes, rightBytes) > 0 {
		leftBytes, rightBytes = rightBytes, leftBytes
	}

	combined := make([]byte, 0, left.Size()+right.Size())
	combined = append(combined, leftBytes...)
	combined = append(combined, rightBytes...)

	return h.Hash(combined), nil
}
func (h *Sha256Hasher) EmptyHash() Hash {
	return h.Hash([]byte{})
}
