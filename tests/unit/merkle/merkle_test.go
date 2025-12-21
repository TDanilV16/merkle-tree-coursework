package merkle

import (
	"testing"

	"github.com/TDanilV16/merkle-tree-coursework/internal/merkletree"
)

func TestAdd(t *testing.T) {
	want := 5
	actual, err := merkletree.Add(2, 3)
	if want != actual || err != nil {
		t.Errorf("Add() error = (%v,%v), want (%v,%v)", err, actual, want, err)
	}
}
