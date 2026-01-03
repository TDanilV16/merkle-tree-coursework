package main

import (
	"context"
	"fmt"

	"github.com/TDanilV16/merkle-tree-coursework/pkg/config"
	"github.com/TDanilV16/merkle-tree-coursework/pkg/hasher"
	"github.com/TDanilV16/merkle-tree-coursework/pkg/logger"
)

func main() {
	configPath := "configs/config.yaml"
	cfg := config.LoadOrCreate(configPath)

	l := logger.SetupFromConfig(cfg)
	ctx := context.Background()

	l.Info(ctx, "Logger initialized")

	l.Info(ctx, "Hello, world")

	sha256Hasher := hasher.NewSha256Hasher()
	instrumentedHasher := hasher.NewInstrumentedHasher(sha256Hasher, l, ctx)

	hash := instrumentedHasher.Hash([]byte("test"))

	l.Infof(ctx, fmt.Sprintf("hash: %s", hash.String()))
}
