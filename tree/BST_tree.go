package tree

import (
	"fmt"
)

type bstnode struct {
	value int
	left  *bstnode
	right *bstnode
}

type BST struct {
	root *bstnode
}

//创建二叉搜索树
func NewBST(nodes ...int) *BST {
	head := &BST{nil}
	if len(nodes) == 0 {
		return head
	} else {
		head.Insert(nodes[0], nodes[1:len(nodes)]...)
	}
	return head
}

//插入
func (tree *bstnode) insert(node int) {
	if node < tree.value {
		if tree.left == nil {
			tree.left = &bstnode{node, nil, nil}
		} else {
			tree.left.insert(node)
		}
	} else {
		if tree.right == nil {
			tree.right = &bstnode{node, nil, nil}
		} else {
			tree.right.insert(node)
		}
	}
}

func (tree *bstnode) inOrder(out *[]int) {
	if tree == nil {
		return
	}
	if tree.left != nil {
		tree.left.inOrder(out)
	}
	(*out) = append((*out), tree.value)
	if tree.right != nil {
		tree.right.inOrder(out)
	}
}

//删除最小节点
func (tree *bstnode) deletemin() *bstnode {
	p, q := tree, tree.right
	if q.left == nil {
		p.right = q.right
		q.right = nil
		return q
	}
	for q.left != nil {
		p = q
		q = q.left
	}
	p.left = q.right
	q.right = nil
	return q
}

//删除节点
func (tree *bstnode) delete(node int) *bstnode {
	var p **bstnode = &tree
	var q = tree
	for q != nil && node != q.value {
		if node < q.value {
			p = &(q.left)
			q = q.left
		} else if node > q.value {
			p = &(q.left)
			q = q.right
		}
	}
	if q == nil {
		return tree
	}
	if q.left != nil && q.right != nil {
		*p = q.deletemin()
		(*p).left, (*p).right = q.left, q.right
		q.left, q.right = nil, nil
	} else if q.left != nil {
		*p = q.left
		q.left = nil
	} else {
		*p = q.right
		q.right = nil
	}
	return tree
}

//插入
func (bst *BST) Insert(node int, nodes ...int) {
	if bst.IsEmpty() {
		bst.root = &bstnode{node, nil, nil}
		for i := 0; i < len(nodes); i++ {
			bst.root.insert(nodes[i])
		}
	} else {
		bst.root.insert(node)
		for i := 0; i < len(nodes); i++ {
			bst.root.insert(nodes[i])
		}
	}
}

//删除
func (bst *BST) Delete(node int) {
	if bst.root == nil {
		return
	}
	bst.root = bst.root.delete(node)
}

func (bst *BST) Clear() {
	bst.root = nil
}

//是否为空树
func (bst *BST) IsEmpty() bool {
	if bst.root == nil {
		return true
	}
	return false
}

//转换为切片
func (bst *BST) ToSlice() []int {
	slice := []int{}
	bst.root.inOrder(&slice)
	return slice
}

//从切片构建
func (bst *BST) FromSlice(nodes []int) {
	if len(nodes) == 0 {
		return
	}
	bst.Insert(nodes[0], nodes[1:]...)
}

//查找
func (bst *BST) Search(node int) bool {
	p := bst.root
	for p != nil {
		if node < p.value {
			p = p.left
		} else if node > p.value {
			p = p.right
		} else {
			return true
		}
	}
	return false
}

//最小值
func (bst *BST) Min() int {
	if bst.IsEmpty() {
		panic("[tree.Min] Empty Tree")
	}
	p := bst.root
	for p.left != nil {
		p = p.left
	}
	return p.value
}

//最大值
func (bst *BST) Max() int {
	if bst.IsEmpty() {
		panic("[tree.Min] Empty Tree")
	}
	p := bst.root
	for p.right != nil {
		p = p.right
	}
	return p.value
}

func (bst *BST) String() string {
	slice := bst.ToSlice()
	return fmt.Sprintf("%v\n", slice)
}
