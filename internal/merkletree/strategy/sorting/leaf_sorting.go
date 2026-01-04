package sorting

import (
	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree"
	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree/strategy"
)

type LeafSortingStrategy struct{}

func NewLeafSortingStrategy() *LeafSortingStrategy {
	return &LeafSortingStrategy{}
}

func (l LeafSortingStrategy) Sort(nodes []*merkletree.Node) []*merkletree.Node {
	if shouldSort(nodes) {
		return sortNodes(nodes)
	}
	return sortNodes(nodes)
}

func (l LeafSortingStrategy) Name() string {
	return "leaf_sorting"
}

func (l LeafSortingStrategy) Description() string {
	return "Leaf sorting. Ethereum-style"
}

func shouldSort(nodes []*merkletree.Node) bool {
	firstNodeDepth := nodes[0].Depth
	for _, node := range nodes[1:] {
		if node.Depth != firstNodeDepth {
			return false
		}
	}

	return true
}

var _ strategy.SortingStrategy = LeafSortingStrategy{}
var _ strategy.SortingStrategy = (*LeafSortingStrategy)(nil)
