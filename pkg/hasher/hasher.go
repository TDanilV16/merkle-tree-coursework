package hasher

import (
	"crypto/sha256"
	"io"
	"os"
)

type Hasher interface {
	Hash(data []byte) Hash
	HashFile(filepath string) (Hash, error)
	VerifyHash(hash Hash, data []byte) bool
	HashString(str string) Hash
	HashConcat(left, right Hash) Hash
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

func (h *Sha256Hasher) HashConcat(left, right Hash) Hash {
	leftBytes := left.Bytes()
	rightBytes := right.Bytes()

	if left.Size() > right.Size() {
		leftBytes, rightBytes = rightBytes, leftBytes
	}

	combined := make([]byte, 0, left.Size()+right.Size())
	combined = append(combined, leftBytes...)
	combined = append(combined, rightBytes...)

	return h.Hash(combined)
}
