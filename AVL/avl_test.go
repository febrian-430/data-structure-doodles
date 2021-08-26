package avl_test

import (
	"testing"

	. "github.com/febrian-430/data-structure-doodles/AVL"
)

func assert(t *testing.T, node *TreeNode, expect int) {
	if node.Val != expect {
		t.Logf("Expected %v, got %v", expect, node.Val)
		t.FailNow()
	}
}

func TestPush_WhenEmptyPushedAsRoot(t *testing.T) {
	tree := NewAVLTree()
	expect := 5
	tree.Push(expect)
	if tree.Root.Val != expect {
		t.Logf("Expected %v, got %v", expect, tree.Root.Val)
	}
}

func TestPush_WhenRootIsBigger(t *testing.T) {
	tree := NewAVLTree()
	expectLeft := 4

	tree.Push(5)
	tree.Push(4)

	if tree.Root.Left == nil {
		t.Log("Expected not nil, got nil")
		t.FailNow()
	}

	if tree.Root.Left.Val != expectLeft {
		t.Logf("Expected %v, got %v", expectLeft, tree.Root.Left.Val)
		t.Fail()
	}
}

func TestPush_WhenRootIsSmaller(t *testing.T) {
	tree := NewAVLTree()
	expectRight := 6

	tree.Push(5)
	tree.Push(expectRight)

	if tree.Root.Right == nil {
		t.Log("Expected not nil, got nil")
		t.FailNow()
	}

	if tree.Root.Right.Val != expectRight {
		t.Logf("Expected %v, got %v", expectRight, tree.Root.Right.Val)
		t.Fail()
	}
}

func TestPush_LeftSkewer(t *testing.T) {
	tree := NewAVLTree()
	expect := []int{3, 2, 1}

	for _, v := range expect {
		tree.Push(v)
	}
	cur := tree.Root
	for _, v := range expect {
		if cur.Val != v {
			t.Logf("Expected %v, got %v", v, cur.Val)
			t.FailNow()
		}
		cur = cur.Left
	}
}

func TestPush_RightSkewer(t *testing.T) {
	tree := NewAVLTree()
	expect := []int{1, 2, 3}

	for _, v := range expect {
		tree.Push(v)
	}
	cur := tree.Root
	for _, v := range expect {
		if cur.Val != v {
			t.Logf("Expected %v, got %v", v, cur.Val)
			t.FailNow()
		}
		cur = cur.Right
	}
}

func TestPush_LeftRightSkewer(t *testing.T) {
	tree := NewAVLTree()
	expect := []int{2, 4, 3}

	for _, v := range expect {
		tree.Push(v)
	}
	cur := tree.Root
	assert(t, cur, expect[0])
	assert(t, cur.Right, expect[1])
	assert(t, cur.Right.Left, expect[2])
}

func TestPush_RightLeftSkewer(t *testing.T) {
	tree := NewAVLTree()
	expect := []int{3, 1, 2}

	for _, v := range expect {
		tree.Push(v)
	}
	cur := tree.Root
	assert(t, cur, expect[0])
	assert(t, cur.Left, expect[1])
	assert(t, cur.Left.Right, expect[2])
}
