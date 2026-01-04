package sorting

import (
	"bytes"
	"sort"

	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree"
)

func sortNodes(nodes []*merkletree.Node) []*merkletree.Node {
	sort.Slice(nodes, func(i, j int) bool { return bytes.Compare(nodes[i].Hash.Bytes(), nodes[j].Hash.Bytes()) == -1 })
	return nodes
}
