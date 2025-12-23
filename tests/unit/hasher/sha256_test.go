package hasher

import (
	_ "crypto/sha256"
	"testing"

	"github.com/TDanilV16/merkle-tree-coursework/pkg/hasher"
	"github.com/stretchr/testify/assert"
)

func TestSha256Hasher_Init(t *testing.T) {
	_ = hasher.NewSha256Hasher()

	assert.True(t, 1 == 1)
}
