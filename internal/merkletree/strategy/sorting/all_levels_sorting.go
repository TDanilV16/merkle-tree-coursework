package sorting

import (
	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree"
	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree/strategy"
)

type AllLevelsSortingStrategy struct{}

func NewAllLevelsSortingStrategy() *AllLevelsSortingStrategy {
	return &AllLevelsSortingStrategy{}
}

func (a AllLevelsSortingStrategy) Sort(nodes []*merkletree.Node) []*merkletree.Node {
	return sortNodes(nodes)
}

func (a AllLevelsSortingStrategy) Name() string {
	return "all_level_sorting"
}

func (a AllLevelsSortingStrategy) Description() string {
	return "All level sorting. Strict determinism"
}

var _ strategy.SortingStrategy = AllLevelsSortingStrategy{}
var _ strategy.SortingStrategy = (*AllLevelsSortingStrategy)(nil)
