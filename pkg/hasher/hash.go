package hasher

import (
	"bytes"
	"encoding/hex"
)

type Hash interface {
	Bytes() []byte
	String() string
	Equal(other Hash) bool
	Size() int
	IsZero() bool
	StringShort(length int) string
}

func NewHash(data []byte) Hash {
	if data == nil {
		return &hashImpl{[]byte{}}
	}
	copied := make([]byte, len(data))
	copy(copied, data)
	return &hashImpl{copied}
}

type hashImpl struct {
	bytes []byte
}

func (h *hashImpl) Bytes() []byte {
	if h.bytes == nil {
		return []byte{}
	}
	result := make([]byte, len(h.bytes))
	copy(result, h.bytes)
	return result
}

func (h *hashImpl) String() string {
	return hex.EncodeToString(h.bytes)
}

func (h *hashImpl) Equal(other Hash) bool {
	if other == nil {
		return false
	}

	return bytes.Equal(h.bytes, other.Bytes())
}

func (h *hashImpl) Size() int {
	return len(h.bytes)
}

func (h *hashImpl) IsZero() bool {
	for _, b := range h.bytes {
		if b != 0 {
			return false
		}
	}
	return true
}

func (h *hashImpl) StringShort(length int) string {
	if h.Size() > length {
		return h.String()[:length] + "..."
	}

	return h.String()
}
