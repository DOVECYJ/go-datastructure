package tree

import (
	"fmt"
)

type Color bool

func (c Color) String() string {
	if c {
		return "Black"
	} else {
		return " Red "
	}
}

const (
	RED   Color = false //红色
	BLACK Color = true  //黑色
)

//红黑树节点
type rbnode struct {
	value  int
	color  Color
	left   *rbnode
	right  *rbnode
	parent *rbnode
}

func (n *rbnode) printnode() {
	if n == nil {
		fmt.Printf("<node>--[Empty]\n")
		return
	}
	fmt.Printf("<node>-[%3d #%v]\n", n.value, n.color)
}

//插入节点
func (n *rbnode) insert(value int) *rbnode {
	if n == nil {
		return newBlack(value)
	}
	node := newRed(value)
	p := n
	for p != nil {
		if value < p.value {
			if p.left == nil {
				break
			}
			p = p.left
		} else {
			if p.right == nil {
				break
			}
			p = p.right
		}
	}

	node.parent = p
	if value < p.value {
		p.left = node
	} else {
		p.right = node
	}
	if p.color == BLACK {
		return n
	}

	root := n
	node.fixup(&root)
	return root
}

func XOR(a, b bool) bool {
	return a != b
}

/*
 * 单支节点，必为黑父红子，一步到位
 * 红叶节点，直接删除
 * 黑叶节点，慎重考虑
 * 两个子节点 → 变成上面三种情况
 */
func (n *rbnode) delete(value int) *rbnode {
	p := n
	for p != nil && p.value != value {
		if value < p.value {
			p = p.left
		} else {
			p = p.right
		}
	}
	if p == nil {
		return n
	}
	if p == n && n.isLeaf() {
		return nil
	}
	root := n
	p.del_node(&root)
	return root
}

//删除节点辅助函数
func (n *rbnode) del_node(root **rbnode) {
	if XOR(n.left == nil, n.right == nil) { //只有一个子节点
		n.del_one_branch(root)
	} else if n.left == nil && n.right == nil { //没有子节点
		if n.isRed() {
			n.del_red_leaf()
		} else {
			n.del_black_leaf(root)
		}

	} else if n.left != nil && n.right != nil { //有两个子节点
		min := n.right.find_min()
		n.value = min.value
		min.del_node(root)
	}
}

//颜色修正
func (n *rbnode) fixup(root **rbnode) {
	n.processCase(root)
}

func (n *rbnode) getCase() int {
	if n.parent.isBlack() {
		return 0
	} else if n.uncle().isRed() {
		return 1
	} else if n.parent.isLeft() {
		if n.isRight() {
			return 2
		} else {
			return 3
		}
	} else {
		if n.isLeft() {
			return 4
		} else {
			return 5
		}
	}
}

func (n *rbnode) processCase(root **rbnode) {
	switch n.getCase() {
	case 1:
		n.processCase1(root)
	case 2:
		n.processCase2(root)
	case 3:
		n.processCase3(root)
	case 4:
		n.processCase4(root)
	case 5:
		n.processCase5(root)
	default:
		return
	}
}

/*
 * 父节点和叔父节点均为红色
 * 将父节点和叔父节点设为黑色
 * 将祖父节点设为红色
 * 将当前节点设为祖父节点
 * 若祖父节点就是根节点，将祖父节点设为黑色
 * 在祖父节点上再做平衡
 * 不会改变树的结构，也不改变根节点，树的黑高度增加1
 * 但是可能改变根节点的颜色
 */
func (n *rbnode) processCase1(root **rbnode) {
	n.parent.toBlack()
	n.uncle().toBlack()
	if n.grandpa() != *root {
		n.grandpa().toRed()
		n.grandpa().processCase(root)
	}
}

/*
 * 父节点为红色，叔父节点为黑色，当前节点为右孩子
 * 以父节点为当前节点左旋
 * 变成case 3
 * 在原父节点上再做平衡
 * 不会改变根节点
 */
func (n *rbnode) processCase2(root **rbnode) {
	p := n.parent
	g := n.parent.parent
	if p.isLeft() {
		g.left = p.left_spin()
	} else {
		g.right = p.left_spin()
	}
	p.processCase(root)
}

/*
 * 父节点为红色，叔父节点为黑色，当前节点为左孩子
 * 将父节点设为黑色，祖父节点设为红色
 * 以祖父节点为当前节点右旋
 * 算法结束，树重新平衡
 * 可能改变根节点
 */
func (n *rbnode) processCase3(root **rbnode) {
	g := n.grandpa()
	n.parent.toBlack()
	g.toRed()

	if g == *root {
		*root = g.right_spin()
	} else {
		gp := g.parent
		if g.isLeft() {
			gp.left = g.right_spin()
		} else {
			gp.right = g.right_spin()
		}
	}
}

//case2的对称情况
func (n *rbnode) processCase4(root **rbnode) {
	p := n.parent
	g := n.parent.parent
	if p.isLeft() {
		g.left = p.right_spin()
	} else {
		g.right = p.right_spin()
	}
	p.processCase(root)
}

//case3的对称情况
func (n *rbnode) processCase5(root **rbnode) {
	g := n.grandpa()
	n.parent.toBlack()
	g.toRed()

	if g == *root {
		*root = g.left_spin()
	} else {
		gp := g.parent
		if g.isLeft() {
			gp.left = g.left_spin()
		} else {
			gp.right = g.left_spin()
		}
	}
}

//所谓旋转就是爸爸变孙子，儿子变爸爸
//在红黑树中，左旋不会改变根节点
func (n *rbnode) left_spin() *rbnode {
	right := n.right
	right.parent = n.parent
	n.right = right.left
	if right.left != nil {
		right.left.parent = n
	}
	right.left = n
	n.parent = right
	return right
}

//右旋可能改变根节点
func (n *rbnode) right_spin() *rbnode {
	left := n.left
	left.parent = n.parent
	n.left = left.right
	if left.right != nil {
		left.right.parent = n
	}
	left.right = n
	n.parent = left
	return left
}

/*
 * 红兄弟还是黑兄弟
 * 兄弟有无侄子
 * 原侄子还是近侄子
 * 没有侄子的话，父亲是红是黑
 */
func (n *rbnode) delCase() int {
	if n.brother().isRed() { //兄弟为红色，父亲必为黑，必有两个黑孩子
		return 1
	} else { //黑兄弟
		if n.brother().hasChild() { //兄弟有孩子，必为红孩儿
			if n.isLeft() && n.brother().right != nil || n.isRight() && n.brother().left != nil { //远侄子
				return 2
			} else { //近侄子
				return 3
			}
		} else { //兄弟无孩子
			if n.parent.isRed() { //父红兄黑
				return 4
			} else { //父兄全黑
				return 5
			}
		}
	}
}

//寻找最小节点
func (n *rbnode) find_min() *rbnode {
	p := n
	for p.left != nil {
		p = p.left
	}
	return p
}

/*
 * 删除单支的节点，必为黑父一红子
 * 删除该节点，并将孩子设为黑色
 * 树平衡
 * 如果删除的就是根节点，则会改变根节点。
 */
func (n *rbnode) del_one_branch(root **rbnode) {
	p := n.left
	if p == nil {
		p = n.right
	}
	if p != nil {
		p.toBlack()
		p.parent = n.parent
		if n == *root { //n是头节点
			*root = p
		} else { //n不是头节点
			if n.isLeft() {
				n.parent.left = p
			} else {
				n.parent.right = p
			}
		}
	}
	n.left, n.right = nil, nil
}

/*
 * 不会破坏树的平衡
 * 也不会改变根节点
 */
func (n *rbnode) del_red_leaf() {
	if n.isLeft() {
		n.parent.left = nil
	} else {
		n.parent.right = nil
	}
}

func (n *rbnode) del_black_leaf(root **rbnode) {
	switch n.delCase() {
	case 1:
		n.delCase1(root)
	case 2:
		n.delCase2(root)
	case 3:
		n.delCase3(root)
	case 4:
		n.delCase4()
	case 5:
		n.delCase5(root)
	}
}

/*
 * case.1
 * 兄弟是红色(父亲必黑，兄弟必有两个黑孩子)
 * 父变红，兄变黑，以父为轴旋转
 * 转入case.4
 * 如果父节点是根，则会改变根节点
 */
func (n *rbnode) delCase1(root **rbnode) {
	n.parent.toRed()
	n.brother().toBlack()

	if n.isLeft() { //删除的是左节点
		if n.parent == *root { //父亲是根节点
			*root = n.parent.left_spin()
		} else { //父亲不是根节点
			if n.parent.isLeft() {
				n.grandpa().left = n.parent.left_spin()
			} else {
				n.grandpa().right = n.parent.left_spin()
			}
		}
	} else { //删除的是右节点
		if n.parent == *root { //父亲是根节点
			*root = n.parent.right_spin()
		} else { //父亲不是根节点
			if n.parent.isLeft() {
				n.grandpa().left = n.parent.right_spin()
			} else {
				n.grandpa().right = n.parent.right_spin()
			}
		}
	}

	n.del_black_leaf(root)
}

/*
 * case.2
 * 兄弟是黑色，远侄子为红色(近侄子为空或红)
 * 父亲和兄弟互换颜色，旋转，远侄子设为黑
 * 删除该节点，树平衡
 * 结束
 * 如果父节点是根节点，则会改变根节点
 */
func (n *rbnode) delCase2(root **rbnode) {
	//println("[case.2]")
	n.parent.color, n.brother().color = n.brother().color, n.parent.color

	if n.isLeft() { //删除的是左节点
		n.brother().right.toBlack()
		if n.parent == *root { //父节点是根节点
			*root = n.parent.left_spin()
		} else { //父节点不是根节点
			if n.parent.isLeft() { //父节点是左节点
				n.grandpa().left = n.parent.left_spin()
			} else { //父节点是右节点
				n.grandpa().right = n.parent.left_spin()
			}
		}
		if n.isLeaf() { //如果是叶节点才删除
			n.parent.left = nil //删除节点
		}
	} else { //删除的是右节点
		n.brother().left.toBlack()
		if n.parent == *root { //父节点是根节点
			*root = n.parent.right_spin()
		} else { //父节点不是根节点
			if n.parent.isLeft() { //父节点是左节点
				n.grandpa().left = n.parent.right_spin()
			} else { //父节点是右节点
				n.grandpa().right = n.parent.right_spin()
			}
		}
		if n.isLeaf() { //如果是叶节点才删除
			n.parent.right = nil //删除节点
		}
	}
}

/*
 * case.3
 * 兄弟是黑色，近侄子为红色(远侄子为空)
 * 兄弟和它儿子互换颜色，旋转
 * 转入case.2
 */
func (n *rbnode) delCase3(root **rbnode) {
	//println("[case.3]")
	if n.isLeft() { //删除的为左节点
		n.brother().color, n.brother().left.color = n.brother().left.color, n.brother().color
		n.parent.right = n.brother().right_spin()
	} else { //删除的为右节点
		n.brother().color, n.brother().right.color = n.brother().right.color, n.brother().color
		n.parent.left = n.brother().left_spin()
	}
	n.delCase2(root)
}

/*
 * case.4
 * 父亲是红色，兄弟是黑色，(没有侄儿)
 * 父亲变黑，兄弟变红
 * 删除该节点，树平衡
 * 结束
 */
func (n *rbnode) delCase4() {
	n.parent.toBlack()
	n.brother().toRed()
	if n.isLeaf() { //如果是叶节点才删除
		if n.isLeft() {
			n.parent.left = nil
		} else {
			n.parent.right = nil
		}
	}
}

/*
 * case.5
 * 父兄皆黑，(没有侄儿)
 * 兄弟变红
 * 删除该节点
 * 以父节点为起点再平衡
 */
func (n *rbnode) delCase5(root **rbnode) {
	n.brother().toRed()
	if n.isLeaf() { //如果是叶节点就删除
		if n.isLeft() {
			n.parent.left = nil
		} else {
			n.parent.right = nil
		}
	}
	if n.parent == *root {
		return
	}
	n.parent.del_black_leaf(root)
}

//反转颜色
func (n *rbnode) flip_color() {
	if n == nil {
		return
	}
	n.color = !n.color
}

//将节点变成黑色
func (n *rbnode) toBlack() {
	if n == nil {
		return
	}
	n.color = BLACK
}

//将节点变成红色
func (n *rbnode) toRed() {
	if n == nil {
		return
	}
	n.color = RED
}

//判断节点是否为红节点
func (n *rbnode) isRed() bool {
	if n != nil && n.color == RED {
		return true
	}
	return false
}

//判断节点是否为黑节点
func (n *rbnode) isBlack() bool {
	if n == nil {
		return true
	} else if n.color == BLACK {
		return true
	}
	return false
}

////判断节点是否为左孩子
func (n *rbnode) isLeft() bool {
	if n.parent != nil {
		if n == n.parent.left {
			return true
		} else {
			return false
		}
	}
	return false
}

//判断节点是否为右孩子
func (n *rbnode) isRight() bool {
	if n.parent != nil {
		if n == n.parent.right {
			return true
		} else {
			return false
		}
	}
	return false
}

//判断节点是否有孩子
func (n *rbnode) hasChild() bool {
	if n != nil && n.left == nil && n.right == nil {
		return false
	}
	return true
}

//判断节点是否为叶子节点
func (n *rbnode) isLeaf() bool {
	if n != nil && n.left == nil && n.right == nil {
		return true
	}
	return false
}

//获取节点的祖父节点
func (n *rbnode) grandpa() *rbnode {
	if n != nil && n.parent != nil {
		return n.parent.parent
	}
	return nil
}

//获取节点的叔父节点
func (n *rbnode) uncle() *rbnode {
	if n.grandpa() != nil {
		return n.parent.brother()
	}
	return nil
}

//获取节点的兄弟节点
func (n *rbnode) brother() *rbnode {
	if n.parent != nil {
		if n == n.parent.left {
			return n.parent.right
		} else {
			return n.parent.left
		}
	}
	return nil
}

//for test
func (n *rbnode) printInOrder() {
	if n.left != nil {
		n.left.printInOrder()
	}
	fmt.Printf("[%3d #%v]\n", n.value, n.color)
	if n.right != nil {
		n.right.printInOrder()
	}
}

func (n *rbnode) inOrder(out *[]int) {
	if n.left != nil {
		n.left.inOrder(out)
	}
	(*out) = append((*out), n.value)
	if n.right != nil {
		n.right.inOrder(out)
	}
}

//新建一个红节点
func newRed(value int) *rbnode {
	return &rbnode{value, RED, nil, nil, nil}
}

//新建一个黑节点
func newBlack(value int) *rbnode {
	return &rbnode{value, BLACK, nil, nil, nil}
}

type RBT struct {
	root *rbnode
}

//插入
func (t *RBT) Insert(value int) {
	t.root = t.root.insert(value)
}

//删除
func (t *RBT) Delete(value int) {
	if t.IsEmpty() {
		return
	}
	t.root = t.root.delete(value)
}

//是否为空树
func (t *RBT) IsEmpty() bool {
	if t.root == nil {
		return true
	}
	return false
}

func (t *RBT) FromSlice(slice []int) {
	if len(slice) == 0 {
		return
	}
	for i := range slice {
		t.Insert(slice[i])
	}
}

func (t *RBT) ToSlice() []int {
	slice := []int{}
	t.root.inOrder(&slice)
	return slice
}

func (t *RBT) String() string {
	slice := t.ToSlice()
	return fmt.Sprintf("%v\n", slice)
}

//for test
func (t *RBT) Print() {
	if t.IsEmpty() {
		fmt.Println("[RBTree.Print] Empty Tree.")
	} else {
		t.root.printInOrder()
	}
}

//for test
func (t *RBT) PrintRoot() {
	if t.IsEmpty() {
		fmt.Println("[RBTree.PrintRoot] Empty Root")
	} else {
		fmt.Printf("<root: [%3d #%v]>\n\n", t.root.value, t.root.color)
	}
}

//创建新红黑树
func NewRBT(values ...int) *RBT {
	rbt := new(RBT)
	for _, v := range values {
		rbt.Insert(v)
	}
	return rbt
}
