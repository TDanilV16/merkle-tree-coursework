package hasher

import (
	"context"

	"github.com/TDanilV16/merkle-tree-coursework/pkg/logger"
)

type InstrumentedHasher struct {
	Hasher
	logger *logger.Logger
	ctx    context.Context
}

func NewInstrumentedHasher(hasher Hasher, l *logger.Logger, ctx context.Context) *InstrumentedHasher {
	if hasher == nil {
		hasher = NewSha256Hasher()
	}
	if l == nil {
		l = logger.Default()
	}
	if ctx == nil {
		ctx = context.Background()
	}
	return &InstrumentedHasher{
		hasher,
		l,
		ctx,
	}
}

func (ih *InstrumentedHasher) Hash(data []byte) Hash {
	ih.logger.Debugf(ih.ctx, "Hashing data, data_szie:%d", len(data))

	result := ih.Hasher.Hash(data)

	ih.logger.Debugf(ih.ctx, "Hashing completed, data_szie:%d, hash:%s", result.Size(), result.String())

	return result
}

func (ih *InstrumentedHasher) HashFile(filepath string) (Hash, error) {
	ih.logger.Infof(ih.ctx, "Hashing file, filepath:%s", filepath)

	result, err := ih.Hasher.HashFile(filepath)

	if err != nil {
		ih.logger.Errorf(ih.ctx, "Hashing file failed, filepath:%s err:%s", filepath, err)
	} else {
		ih.logger.Infof(ih.ctx, "Hashing file completed, filepath:%s, result:%d", filepath, result)
	}

	return result, err
}

func (ih *InstrumentedHasher) HashConcat(left, right Hash) (Hash, error) {
	if left.IsZero() {
		ih.logger.Warnf(ih.ctx, "Left hash is zero. Replacing with EmptyHash")
		left = ih.EmptyHash()
	}

	if right.IsZero() {
		ih.logger.Warnf(ih.ctx, "Right hash is zero. Replacing with EmptyHash")
		right = ih.EmptyHash()
	}

	result, err := ih.Hasher.HashConcat(left, right)

	if err != nil {
		ih.logger.Errorf(ih.ctx, "HashConcat failed, err:%s", err)
	}

	return result, err
}
