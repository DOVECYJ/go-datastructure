package list

import (
	"errors"
	"sync"
)

type snode struct {
	value interface{}
	next  *node
}

type SList struct {
	head   *node
	tail   *node
	length int
	lock   sync.RWMutex
}

//创建链表
func NewSafeList(values ...interface{}) *SList {
	l := &SList{}
	if len(values) == 0 {
		return l
	}
	l.AddRange(values...)
	return l
}

//在链表头部插入
func (l *SList) Insert(value interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.head == nil {
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
func (l *SList) InsertAt(index int, value interface{}) { // dead lock
	l.lock.Lock()
	defer l.lock.Unlock()
	if index < 0 || index > l.length {
		panic("Index out of range")
	}

	if index == 0 {
		if l.head == nil {
			l.head = &node{value, nil}
			l.tail = l.head
			l.length = 1
		} else {
			p := l.head
			l.head = &node{value, p}
			l.length++
		}
		return
	}
	if index == l.length {
		if l.head == nil {
			l.head = &node{value, nil}
			l.tail = l.head
			l.length = 1
		} else {
			l.tail.next = &node{value, nil}
			l.tail = l.tail.next
			l.length++
		}
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
func (l *SList) Add(value interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	if l.head == nil {
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
func (l *SList) AddRange(values ...interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	for i := range values {
		if l.head == nil {
			l.head = &node{values[i], nil}
			l.tail = l.head
			l.length = 1
		} else {
			l.tail.next = &node{values[i], nil}
			l.tail = l.tail.next
			l.length++
		}
	}
}

//批量在头部插入
func (l *SList) InsertRange(values ...interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	for i := range values {
		if l.head == nil {
			l.head = &node{values[i], nil}
			l.tail = l.head
			l.length = 1
		} else {
			p := l.head
			l.head = &node{value[i], p}
			l.length++
		}
	}
}

//链表中是否包含值为value的元素
func (l *SList) HasValue(value interface{}) bool {
	l.lock.RLock()
	defer l.lock.RUnlock()
	for p := l.head; p != nil; p = p.next {
		if p.value == value {
			return true
		}
	}
	return false
}

//删除第一个值为value的元素
func (l *SList) RemoveValue(value interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
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
func (l *SList) RemoveLastValue(value interface{}) {
	l.lock.Lock()
	defer l.lock.Unlock()
	var lst, p *node
	for p = l.head; p != nil; p = p.next {
		if p.value == value {
			lst = p
		}
	}
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
	if lst == l.tail {
		l.tail = p
	}
	p.next = lst.next
	l.length--
}

//删除下标为index的元素
func (l *SList) RemoveAt(index int) {
	l.lock.Lock()
	defer l.lock.Unlock()
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

//清空链表
func (l *SList) Clear() {
	l.lock.Lock()
	l.head = nil
	l.tail = nil
	l.length = 0
	l.lock.Unlock()
}

//判断链表是否为空
func (l *SList) Empty() bool {
	l.lock.RLock()
	empty := (l.head == nil)
	l.lock.RUnlock()
	return empty
}

//获取链表长度
func (l *SList) Length() int {
	l.lock.RLock()
	length := l.length
	l.lock.RUnlock()
	return length
}

func (l *SList) String() string {
	if l.Empty() {
		return "[Empty List]"
	} else {
		str := "["
		l.lock.RLock()
		for p := l.head; p != l.tail; p = p.next {
			str += fmt.Sprintf(" %v →", p.value)
		}
		str += fmt.Sprintf(" %v ]", l.tail.value)
		l.lock.RUnlock()
		return str
	}
}
