package treap

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type tnode struct {
	data        int
	priority    int
	left, right *tnode
}

//左旋
func (t *tnode) left_spin() *tnode {
	p := t.right
	t.right = p.left
	p.left = t
	return p
}

//右旋
func (t *tnode) right_spin() *tnode {
	p := t.left
	t.left = p.right
	p.right = t
	return p
}

//堆调整
func (t *tnode) adjust() *tnode {
	if t.left != nil && t.priority > t.left.priority {
		return t.right_spin()
	}
	if t.right != nil && t.priority > t.right.priority {
		return t.left_spin()
	}
	return t
}

//插入节点
func (t *tnode) insert(data int) *tnode {
	if t == nil {
		return &tnode{data, rand.Intn(1024), nil, nil}
	}
	if data < t.data {
		t.left = t.left.insert(data)
	} else {
		t.right = t.right.insert(data)
	}
	return t.adjust()
}

//删除节点
func (t *tnode) delete(data int) *tnode {
	if t == nil {
		return t
	}
	if t.data == data {
		if t.left != nil && t.right != nil {
			if t.left.priority < t.right.priority {
				return t.right_spin().delete(data)
			} else {
				return t.left_spin().delete(data)
			}
		} else if t.left != nil {
			return t.left
		} else {
			return t.right
		}
	} else if t.data > data {
		t.left = t.left.delete(data)
	} else {
		t.right = t.right.delete(data)
	}

	return t
}

func (t *tnode) inOrder(out *[]int) {
	if t == nil {
		return
	}
	if t.left != nil {
		t.left.inOrder(out)
	}
	(*out) = append((*out), t.data)
	if t.right != nil {
		t.right.inOrder(out)
	}
}

//for test
func (t *tnode) printRoot(tip string) {
	if t == nil {
		fmt.Print(" [O Empty] ")
		return
	}
	fmt.Printf(" [%s# %d:%d] ", tip, t.data, t.priority)
	t.printLeft(t.data)
	t.printRight(t.data)
}

func (t *tnode) printLeft(tip int) {
	if t == nil || t.left == nil {
		fmt.Printf(" [<%d>L Empty] ", tip)
		return
	}
	t.left.printRoot(fmt.Sprintf("<%d>L", tip))
}

func (t *tnode) printRight(tip int) {
	if t == nil || t.right == nil {
		fmt.Printf(" [<%d>R Empty] ", tip)
		return
	}
	t.right.printRoot(fmt.Sprintf("<%d>R", tip))
}

/*-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+*/

type Treap struct {
	root *tnode
	lock sync.RWMutex
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

//创建树堆
func NewTreap(values ...int) *Treap {
	t := &Treap{}
	if len(values) == 0 {
		return t
	}
	for i := range values {
		t.Insert(values[i])
	}
	return t
}

//插入
func (t *Treap) Insert(data int) {
	t.lock.Lock()
	t.root = t.root.insert(data)
	t.lock.Unlock()
}

//删除
func (t *Treap) Delete(data int) {
	t.lock.Lock()
	t.root = t.root.delete(data)
	t.lock.Unlock()
}

//获取根节点
func (t *Treap) GetRoot() (int, error) {
	t.lock.RLock()
	if t.root == nil {
		t.lock.RUnlock()
		return 0, errors.New("Treap is empty.")
	} else {
		data = t.root.data
		t.lock.RUnlock()
		return data, nil
	}
}

//从切片构建
func (t *Treap) FromSlice(slice []int) {
	if len(slice) == 0 {
		return
	}
	t.lock.Lock()
	for i := range slice {
		t.root = t.root.insert(slice[i])
	}
	t.lock.Unlock()
}

//转化为切片
func (t *Treap) ToSlice() []int {
	slice := []int{}
	t.lock.RLock()
	t.root.inOrder(&slice)
	t.lock.RUnlock()
	return slice
}

func (t *Treap) String() string {
	slice := t.ToSlice()
	return fmt.Sprintf("%v\n", slice)
}
