package merkletree

import "github.com/TDanilV16/merkle-tree-coursework/pkg/hasher"

type Node struct {
	Hash      hasher.Hash
	Left      *Node
	Right     *Node
	Depth     int
	Index     int
	IsVirtual bool
}
