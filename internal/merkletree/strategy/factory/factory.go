package factory

import (
	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree/strategy"
	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree/strategy/balancing"
	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree/strategy/sorting"
	"github.com/TDanilV16/merkle-tree-coursework/pkg/hasher"
)

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (f *Factory) BitcoinStyle() strategy.CompositeStrategy {
	return *strategy.NewCompositeStrategy(
		sorting.NewNoSortingStrategy(),
		balancing.NewDuplicateLastStrategy(),
		"bitcoin_style",
		"Bitcoin-style: no sorting, only append last node",
	)
}

func (f *Factory) EthereumStyle(emptyHash hasher.Hash) strategy.CompositeStrategy {
	return *strategy.NewCompositeStrategy(
		sorting.NewLeafSortingStrategy(),
		balancing.NewEmptyHashStrategy(emptyHash),
		"ethereum_style",
		"Ethereum-style: sort leafs, append empty hash",
	)
}

func (f *Factory) CertificateTransparencyStyle(emptyHash hasher.Hash) strategy.CompositeStrategy {
	return *strategy.NewCompositeStrategy(
		sorting.NewAllLevelsSortingStrategy(),
		balancing.NewEmptyHashStrategy(emptyHash),
		"certificate_transparency_style",
		"Certificate-Transparency-Style: sort all levels, append empty hash",
	)
}
