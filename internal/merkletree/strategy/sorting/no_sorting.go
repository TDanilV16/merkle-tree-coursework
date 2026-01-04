package sorting

import (
	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree"
	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree/strategy"
)

type NoSortingStrategy struct{}

func NewNoSortingStrategy() *NoSortingStrategy {
	return &NoSortingStrategy{}
}

func (n NoSortingStrategy) Sort(nodes []*merkletree.Node) []*merkletree.Node {
	return nodes
}

func (n NoSortingStrategy) Name() string {
	return "no_sorting"
}

func (n NoSortingStrategy) Description() string {
	return "No sorting. strategy. Bitcoin-style"
}

var _ strategy.SortingStrategy = LeafSortingStrategy{}
var _ strategy.SortingStrategy = (*NoSortingStrategy)(nil)
