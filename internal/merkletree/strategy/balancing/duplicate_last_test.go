package balancing

import (
	"testing"

	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree"
	"github.com/TDanilV16/merkle-tree-coursework/pkg/hasher"
)

func TestDuplicateLastStrategy(t *testing.T) {
	strategy := NewDuplicateLastStrategy()

	sha256Hasher := hasher.NewSha256Hasher()

	t.Run("Name_and_Description", func(t *testing.T) {
		if strategy.Name() != "duplicate_last" {
			t.Errorf("Name_and_Description should be `duplicate_last`, got %s", strategy.Name())
		}

		desc := strategy.Description()
		if desc != "Using duplicate last strategy. Bitcoin-style" {
			t.Errorf("Description should be `Using duplicate last strategy. Bitcoin-style`, got %s", desc)
		}
	})

	t.Run("Single_node", func(t *testing.T) {
		node := &merkletree.Node{
			Hash:  sha256Hasher.Hash([]byte("single")),
			Depth: 0,
			Index: 0,
		}

		nodes := []*merkletree.Node{node}

		result := strategy.Balance(nodes)

		expected := []*merkletree.Node{node, {
			Hash:      node.Hash,
			Depth:     node.Depth,
			Index:     node.Index,
			IsVirtual: true,
		}}

		if len(result) != len(expected) {
			t.Errorf("Balance() should return %d nodes, got %d", len(expected), len(result))
		}

		duplicate := result[1]

		if duplicate == nil {
			t.Errorf("Balance() should return a duplicate node")
		}

		if !duplicate.Hash.Equal(node.Hash) {
			t.Errorf("Balance() should return a duplicate node with same hash")
		}

		if duplicate.Depth != node.Depth {
			t.Errorf("Balance() should return a duplicate node with same depth")
		}

		if duplicate.Index != node.Index {
			t.Errorf("Balance() should return a duplicate node with same index")
		}

		if !duplicate.IsVirtual {
			t.Errorf("Balance() should return a duplicate node with isVirtual=true")
		}
	})

	t.Run("Even_number_of_nodes", func(t *testing.T) {
		nodes := []*merkletree.Node{
			{
				Hash:  sha256Hasher.Hash([]byte("a")),
				Depth: 0,
				Index: 0,
			},
			{
				Hash:  sha256Hasher.Hash([]byte("b")),
				Depth: 0,
				Index: 0,
			},
		}

		result := strategy.Balance(nodes)

		if result == nil || len(result) != len(nodes) {
			t.Errorf("Balance() should not change nodes if its length is already even %d nodes, got %d", len(nodes), len(result))
		}
	})

	t.Run("Empty_nodes_slice", func(t *testing.T) {
		var nodes []*merkletree.Node
		result := strategy.Balance(nodes)
		if len(result) != 0 {
			t.Errorf("Balance() of empty slice should return an empty slice, got %d", len(result))
		}
	})
}
