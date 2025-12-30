package main

import (
	"context"

	"github.com/TDanilV16/merkle-tree-coursework/pkg/config"
	"github.com/TDanilV16/merkle-tree-coursework/pkg/logger"
)

func main() {
	configPath := "configs/config.yaml"
	cfg := config.LoadOrCreate(configPath)

	logger.SetupFromConfig(cfg)
	ctx := context.Background()

	logger.Info(ctx, "Hello, world")
}
