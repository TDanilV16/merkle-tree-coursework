package balancing

import (
	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree"
	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree/strategy"
)

type NoBalanceStrategy struct{}

func NewNoBalanceStrategy() *NoBalanceStrategy {
	return &NoBalanceStrategy{}
}

func (n NoBalanceStrategy) Balance(nodes []*merkletree.Node) []*merkletree.Node {
	return nodes
}

func (n NoBalanceStrategy) Name() string {
	return "no_balance"
}

func (n NoBalanceStrategy) Description() string {
	return "Returns nodes as it is"
}

var _ strategy.BalancingStrategy = &NoBalanceStrategy{}
var _ strategy.BalancingStrategy = (*NoBalanceStrategy)(nil)
