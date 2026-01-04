package strategy

import "github.com/TDanilV16/merkle-tree-coursework/internal/merkletree"

type CompositeStrategy struct {
	sortingStrategy   SortingStrategy
	balancingStrategy BalancingStrategy
	name              string
	description       string
}

func NewCompositeStrategy(
	sorting SortingStrategy,
	balancing BalancingStrategy,
	name string,
	description string,
) *CompositeStrategy {
	return &CompositeStrategy{
		sortingStrategy:   sorting,
		balancingStrategy: balancing,
		name:              name,
		description:       description,
	}
}

func (s *CompositeStrategy) BalanceNodes(nodes []*merkletree.Node) []*merkletree.Node {
	sorted := s.sortingStrategy.Sort(nodes)

	return s.balancingStrategy.Balance(sorted)
}

func (s *CompositeStrategy) Name() string { return s.name }

func (s *CompositeStrategy) Description() string { return s.description }
