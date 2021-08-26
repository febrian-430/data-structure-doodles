package avl

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type AVLTree struct {
	Root *TreeNode
}

func NewAVLTree() AVLTree {
	return AVLTree{}
}

func (avl *AVLTree) push(cur *TreeNode, val int) *TreeNode {
	if cur == nil {
		cur = &TreeNode{
			Val: val,
		}
		return cur
	}

	if val > cur.Val {
		cur.Right = avl.push(cur.Right, val)
	} else {
		cur.Left = avl.push(cur.Left, val)
	}

	return cur
}

func (avl *AVLTree) Push(val int) {
	avl.Root = avl.push(avl.Root, val)
}
