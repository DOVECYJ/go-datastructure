package list

import (
	"fmt"
)

type node struct {
	value interface{}
	next  *node
}

type List struct {
	head   *node
	tail   *node
	length int
}

//创建链表
func New(values ...interface{}) *List {
	l := &List{nil, nil, 0}
	if len(values) == 0 {
		return l
	}
	l.AddRange(values...)
	return l
}

//在链表头部插入
func (l *List) Insert(value interface{}) {
	if l.Empty() {
		l.head = &node{value, nil}
		l.tail = l.head
		l.length = 1
	} else {
		p := l.head
		l.head = &node{value, p}
		l.length++
	}
}

//在下标为index之前插入
func (l *List) InsertAt(index int, value interface{}) {
	if index < 0 || index > l.length {
		panic("Index out of range")
	}
	if index == 0 {
		l.Insert(value)
		return
	}
	if index == l.length {
		l.Add(value)
		return
	}
	p := l.head
	for ; index > 1; index-- {
		p = p.next
	}
	q := &node{value, p.next}
	p.next = q
	l.length++
}

//在链表尾部插入
func (l *List) Add(value interface{}) {
	if l.Empty() {
		l.head = &node{value, nil}
		l.tail = l.head
		l.length = 1
	} else {
		l.tail.next = &node{value, nil}
		l.tail = l.tail.next
		l.length++
	}
}

//批量在尾部插入元素
func (l *List) AddRange(values ...interface{}) {
	for i := range values {
		l.Add(values[i])
	}
}

//批量在头部插入
func (l *List) InsertRange(values ...interface{}) {
	for i := range values {
		l.Insert(values[i])
	}
}

//链表中是否包含值为value的元素
func (l *List) HasValue(value interface{}) bool {
	for p := l.head; p != nil; p = p.next {
		if p.value == value {
			return true
		}
	}
	return false
}

//删除第一个值为value的元素
func (l *List) RemoveValue(value interface{}) {
	if l.head.value == value {
		if l.tail == l.head {
			l.tail = nil
		}
		l.head = l.head.next
		l.length--
		return
	}
	p, q := l.head, l.head.next
	for q != nil && q.value != value {
		p, q = q, q.next
	}
	if q != nil {
		p.next = q.next
		if q == l.tail {
			l.tail = p
		}
		l.length--
	}
}

//删除最后一个值为value的元素
func (l *List) RemoveLastValue(value interface{}) {
	var lst, p *node
	for p = l.head; p != nil; p = p.next {
		if p.value == value {
			lst = p
		}
	}
	println(lst.value.(int))
	if lst == l.head {
		if l.tail == lst {
			l.tail = nil
		}
		l.head = l.head.next
		l.length--
		return
	}
	for p = l.head; p.next != lst; p = p.next {
	}
	println(p.value.(int))
	if lst == l.tail {
		l.tail = p
	}
	p.next = lst.next
	l.length--
}

//删除下标为index的元素
func (l *List) RemoveAt(index int) {
	if index < -l.length || index >= l.length {
		panic("[List.RemoveAt] Index out of range")
	}

	if index < 0 {
		index += l.length
	}

	if index == 0 {
		if l.length == 1 {
			l.tail = nil
		}
		l.head = l.head.next
	} else {
		p, q := l.head, l.head.next
		for ; index > 1; index-- {
			p = p.next
			q = q.next
		}
		p.next = q.next
		if q == l.tail {
			l.tail = p
		}
	}
	l.length--
}

//删除倒数第index个元素
func (l *List) RemoveLastAt(index int) {
	index = l.length - index
	l.RemoveAt(index)
}

//清空链表
func (l *List) Clear() {
	l.head = nil
	l.tail = nil
	l.length = 0
}

//判断链表是否为空
func (l *List) Empty() bool {
	if l.head == nil {
		return true
	}
	return false
}

//获取链表长度
func (l *List) Length() int {
	return l.length
}

func (l *List) String() string {
	if l.Empty() {
		return "[Empty List]"
	} else {
		str := "["
		for p := l.head; p != l.tail; p = p.next {
			str += fmt.Sprintf(" %v →", p.value)
		}
		str += fmt.Sprintf(" %v ]", l.tail.value)
		return str
	}
}
