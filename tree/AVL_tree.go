package tree

import (
	"fmt"
)

type avlnode struct {
	value   int
	left    *avlnode
	right   *avlnode
	balance int
}

type AVL struct {
	root *avlnode
}

//左单旋
func (this *avlnode) left_single_spin() *avlnode {
	p := this.right
	this.right = p.left
	p.left = this
	p.balance = p.left.height() - p.right.height()
	this.balance = this.left.height() - this.right.height()
	return p
}

//右单旋
func (this *avlnode) right_single_spin() *avlnode {
	p := this.left
	this.left = p.right
	p.right = this
	p.balance = p.left.height() - p.right.height()
	this.balance = this.left.height() - this.right.height()
	return p
}

//先左旋，再右旋
func (this *avlnode) left_right_double_spin() *avlnode {
	this.left = this.left.left_single_spin()
	return this.right_single_spin()
}

//先右旋。再左旋
func (this *avlnode) right_left_double_spin() *avlnode {
	this.right = this.right.right_single_spin()
	return this.left_single_spin()
}

//节点高度
func (this *avlnode) height() int {
	if this == nil {
		return 0
	}
	lheight, rheight := 0, 0
	lheight = this.left.height() + 1
	rheight = this.right.height() + 1
	if lheight > rheight {
		return lheight
	} else {
		return rheight
	}
}

//插入节点
func (this *avlnode) insert(node int) *avlnode {
	if this == nil {
		return &avlnode{node, nil, nil, 0}
	}
	if node < this.value {
		this.left = this.left.insert(node)
	} else {
		this.right = this.right.insert(node)
	}
	return this.adjust()
}

//删除节点
func (this *avlnode) delete(node int) *avlnode {
	if this == nil {
		return nil
	}

	if node < this.value {
		this.left = this.left.delete(node)
	} else if node > this.value {
		this.right = this.right.delete(node)
	} else {
		if this.left != nil && this.right != nil {
			p, q := this.right.deletemin()
			q.left, q.right = this.left, p
			this.left, this.right = nil, nil
			return q.adjust()
		} else if this.left != nil {
			return this.left
		} else {
			return this.right
		}
	}
	return this.adjust()
}

//删除最小节点
func (this *avlnode) deletemin() (node *avlnode, min *avlnode) {
	if this.left == nil {
		node, min = this.right, this
		min.right = nil
		return
	} else {
		node, min = this.left.deletemin()
		this.left = node
		return this.adjust(), min
	}
}

//调整树，使树平衡
func (this *avlnode) adjust() *avlnode {
	if this == nil {
		return this
	}
	this.balance = this.left.height() - this.right.height()
	if this.balance == 2 {
		if this.left.left.height() > this.left.right.height() { //右单旋
			return this.right_single_spin()
		} else { //先左后右双旋
			return this.left_right_double_spin()
		}
	}
	if this.balance == -2 {
		if this.right.right.height() > this.right.left.height() { //左单旋
			return this.left_single_spin()
		} else { //先右后左双旋
			return this.right_left_double_spin()
		}
	}
	return this
}

func (this *avlnode) inOrder(out *[]int) {
	if this == nil {
		return
	}

	this.left.inOrder(out)
	(*out) = append((*out), this.value)
	this.right.inOrder(out)
}

//for test
func (this *avlnode) printTree() {
	queue := []*avlnode{this}
	curr, last := 0, 1
	for len(queue) != 0 {
		n := queue[0]
		queue = queue[1:len(queue)]
		if n.left != nil {
			queue = append(queue, n.left)
			curr += 1
		}
		if n.right != nil {
			queue = append(queue, n.right)
			curr += 1
		}
		n.print()
		last -= 1
		if last == 0 {
			println()
			last = curr
			curr = 0
		}
	}
	println("\n")
}

func (this *avlnode) print() {
	print(this.value, " ")
}

//插入
func (this *AVL) Insert(node int) {
	this.root = this.root.insert(node)
}

//删除
func (this *AVL) Delete(node int) {
	this.root = this.root.delete(node)
}

func (this *AVL) Clear() {
	this.root = nil
}

//从切片构建
func (this *AVL) FromSlice(nodes []int) {
	if len(nodes) == 0 {
		return
	}
	for i := range nodes {
		this.Insert(nodes[i])
	}
}

//转换为切片
func (this *AVL) ToSlice() []int {
	slice := []int{}
	this.root.inOrder(&slice)
	return slice
}

//空树
func (this *AVL) IsEmpty() bool {
	if this.root == nil {
		return true
	}
	return false
}

//根节点
func (this *AVL) RootValue() int {
	if this.IsEmpty() {
		return -1
	}
	return this.root.value
}

func (this *AVL) String() string {
	slice := this.ToSlice()
	return fmt.Sprintf("%v\n", slice)
}

//for tsst
func (this *AVL) PrintTree() {
	if this.IsEmpty() {
		println("[nil] Empty tree.")
		return
	}
	this.root.printTree()
}

//创建新AVL树
func NewAVL(nodes ...int) *AVL {
	root := &AVL{nil}
	for i := range nodes {
		root.Insert(nodes[i])
	}
	return root
}
