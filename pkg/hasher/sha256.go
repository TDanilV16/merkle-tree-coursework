package hasher

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
)

type Hasher interface {
	Hash(data []byte) []byte
	HashString(data string) string
	HashFile(data string) ([]byte, error)
}

func NewSha256Hasher() *Sha256Hasher {
	return &Sha256Hasher{}
}

type Sha256Hasher struct{}

func (h *Sha256Hasher) Hash(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

func (h *Sha256Hasher) HashString(data string) string {
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}

func (h *Sha256Hasher) HashFile(filepath string) ([]byte, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hasher := sha256.New()
	if _, err = io.Copy(hasher, file); err != nil {
		return nil, err
	}

	return hasher.Sum(nil), nil
}

func (h *Sha256Hasher) VerifyHash(hash []byte, data []byte) bool {
	actualHash := h.Hash(data)

	return bytes.Equal(actualHash, hash)
}
