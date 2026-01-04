package balancing

import (
	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree"
	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree/strategy"
	"github.com/TDanilV16/merkle-tree-coursework/pkg/hasher"
)

type EmptyHashStrategy struct {
	emptyHash hasher.Hash
}

func NewEmptyHashStrategy(emptyHash hasher.Hash) *EmptyHashStrategy {
	return &EmptyHashStrategy{emptyHash: emptyHash}
}

func (e EmptyHashStrategy) Balance(nodes []*merkletree.Node) []*merkletree.Node {
	nodesCount := len(nodes)

	if nodesCount%2 != 0 {
		lastNode := nodes[len(nodes)-1]

		duplicate := &merkletree.Node{
			Hash:      e.emptyHash,
			Depth:     lastNode.Depth,
			Index:     lastNode.Index,
			IsVirtual: true,
		}

		return append(nodes, duplicate)
	}

	return nodes
}

func (e EmptyHashStrategy) Name() string {
	return "empty_hash"
}

func (e EmptyHashStrategy) Description() string {
	return "Using empty hash for balancing. Certificate Transparency"
}

var _ strategy.BalancingStrategy = &EmptyHashStrategy{}
var _ strategy.BalancingStrategy = (*EmptyHashStrategy)(nil)
