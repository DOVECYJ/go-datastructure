package queue

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

/*-+-+-+-+-+-+-+-+-+-+-+-+-+-+队列+-+-+-+-+-+-+-+-+-+-+-+-+-+-*/

type queue struct {
	queue []interface{}
}

//入队
func (q *queue) Push(data interface{}) {
	q.queue = append(q.queue, data)
}

//出队
func (q *queue) Pop() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("Pop with empty queue.")
	}
	data := q.queue[0]
	q.queue = q.queue[1:]
	return data, nil
}

//取队头
func (q *queue) Head() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("Empty queue error.")
	}
	data := q.queue[0]
	return data, nil
}

//队列长度
func (q *queue) Len() int {
	return len(q.queue)
}

//队列是否为空
func (q *queue) Empty() bool {
	if len(q.queue) <= 0 {
		return true
	}
	return false
}

func (q *queue) String() string {
	if q.Empty() {
		return "Empty Queue."
	}
	str := ""
	for _, v := range q.queue {
		str += fmt.Sprintf("%v ", v)
	}
	buf := bytes.NewBufferString(strings.Repeat("_", len(str)))
	buf.WriteString("\n")
	buf.WriteString(str)
	buf.WriteString("\n")
	buf.WriteString(strings.Repeat("¯", len(str)))
	return buf.String()
}

//创建队列
func NewQueue() *queue {
	return &queue{
		queue: []interface{}{},
	}
}

/*-+-+-+-+-+-+-+-+-+-+-+-+-+-+双端队列+-+-+-+-+-+-+-+-+-+-+-+-+-+-*/

type dqueue struct {
	queue []interface{}
}

//在队头插入
func (q *dqueue) PushFront(data interface{}) {
	q.queue = append([]interface{}{data}, q.queue...)
}

//在队尾插入
func (q *dqueue) PushBack(data interface{}) {
	q.queue = append(q.queue, data)
}

//从队头出队
func (q *dqueue) PopFront() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("Pop with empty queue")
	}
	data := q.queue[0]
	q.queue = q.queue[1:]
	return data, nil
}

//从队尾出队
func (q *dqueue) PopBack() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("Pop with empty queue")
	}
	index := len(q.queue) - 1
	data := q.queue[index]
	q.queue = q.queue[0:index]
	return data, nil
}

//取队头
func (q *dqueue) Head() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("Empty dqueue error")
	}
	data := q.queue[0]
	return data, nil
}

//取队尾
func (q *dqueue) Tail() (interface{}, error) {
	if q.Empty() {
		return nil, errors.New("Empty dqueue error")
	}
	index := len(q.queue) - 1
	data := q.queue[index]
	return data, nil
}

//队列长度
func (q *dqueue) Len() int {
	return len(q.queue)
}

//队列是否为空
func (q *dqueue) Empty() bool {
	if len(q.queue) == 0 {
		return true
	}
	return false
}

func (q *dqueue) String() string {
	if q.Empty() {
		return "Empty Queue."
	}
	str := ""
	for _, v := range q.queue {
		str += fmt.Sprintf("%v ", v)
	}
	buf := bytes.NewBufferString(strings.Repeat("_", len(str)))
	buf.WriteString("\n")
	buf.WriteString(str)
	buf.WriteString("\n")
	buf.WriteString(strings.Repeat("¯", len(str)))
	return buf.String()
}

//创建双端队列
func NewDQueue() *dqueue {
	return &dqueue{
		queue: []interface{}{},
	}
}

/*-+-+-+-+-+-+-+-+-+-+-+-+-+-+循环队列+-+-+-+-+-+-+-+-+-+-+-+-+-+-*/

//   1     2     3     □
//  ↑                ↑
// head[0]          tail[3]

type cqueue struct {
	queue []interface{}
	head  int
	tail  int
	count int
	size  int
}

//入队
func (q *cqueue) Push(data interface{}) error {
	if q.count == q.size {
		return errors.New("Queue is full.")
	}

	q.queue[q.tail] = data
	q.tail = (q.tail + 1) % q.size
	q.count++
	return nil
}

//出队
func (q *cqueue) Pop() (interface{}, error) {
	if q.count == 0 {
		return nil, errors.New("Queue is empty.")
	}

	data := q.queue[q.head]
	q.head = (q.head + 1) % q.size
	q.count--
	return data, nil
}

//取队头
func (q *cqueue) Head() (interface{}, error) {
	if q.count == 0 {
		return nil, errors.New("Queue is empty.")
	}

	data := q.queue[q.head]
	return data, nil
}

//队列长度
func (q *cqueue) Len() int {
	return q.count
}

//队列是否为空
func (q *cqueue) Empty() bool {
	return q.count == 0
}

//队列是否已满
func (q *cqueue) Full() bool {
	return q.count == q.size
}

//重新设置循环队列大小
func (q *cqueue) Resize(newSize int) (int, error) {
	if newSize <= q.count {
		return q.size, errors.New("New size is too small.")
	}
	old := q.size
	nq := make([]interface{}, newSize, newSize)
	for i := 0; i < q.count; i++ {
		nq[i] = q.queue[(q.head+i)%q.size]
	}
	q.queue = nq
	q.head, q.tail, q.size = 0, q.count, newSize
	return old, nil
}

func (q *cqueue) String() string {
	if q.Empty() {
		return "Empty Queue."
	}
	str := ""
	for i := 0; i < q.count; i++ {
		str += fmt.Sprintf("%v ", q.queue[(q.head+i)%q.size])
	}

	buf := bytes.NewBufferString(strings.Repeat("_", len(str)))
	buf.WriteString("\n")
	buf.WriteString(str)
	buf.WriteString("\n")
	buf.WriteString(strings.Repeat("¯", len(str)))
	return buf.String()
}

//创建循环队列
func NewCQueue(size int) (*cqueue, error) {
	if size <= 0 {
		return nil, errors.New("Can't create loop queue with zero size.")
	}
	return &cqueue{
		queue: make([]interface{}, size, size),
		head:  0,
		tail:  0,
		count: 0,
		size:  size,
	}, nil
}
