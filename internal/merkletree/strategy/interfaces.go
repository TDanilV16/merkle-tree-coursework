package strategy

import "github.com/TDanilV16/merkle-tree-coursework/internal/merkletree"

type BalancingStrategy interface {
	Balance(nodes []*merkletree.Node) []*merkletree.Node

	Name() string

	Description() string
}

type SortingStrategy interface {
	Sort(nodes []*merkletree.Node) []*merkletree.Node

	Name() string

	Description() string
}
