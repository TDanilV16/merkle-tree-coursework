package balancing

import (
	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree"
	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree/strategy"
)

type DuplicateLastStrategy struct{}

func NewDuplicateLastStrategy() *DuplicateLastStrategy {
	return &DuplicateLastStrategy{}
}

func (d DuplicateLastStrategy) Balance(nodes []*merkletree.Node) []*merkletree.Node {
	nodesCount := len(nodes)

	if nodesCount%2 != 0 {
		lastNode := nodes[len(nodes)-1]

		duplicate := &merkletree.Node{
			Hash:      lastNode.Hash,
			Depth:     lastNode.Depth,
			Index:     lastNode.Index,
			IsVirtual: true,
		}

		return append(nodes, duplicate)
	}

	return nodes
}

func (d DuplicateLastStrategy) Name() string {
	return "duplicate_last"
}

func (d DuplicateLastStrategy) Description() string {
	return "Using duplicate last strategy. Bitcoin-style"
}

var _ strategy.BalancingStrategy = &DuplicateLastStrategy{}
var _ strategy.BalancingStrategy = (*DuplicateLastStrategy)(nil)
