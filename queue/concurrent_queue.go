package queue

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"sync"
)

/*-+-+-+-+-+-+-+-+-+-+-+-+-+-+队列+-+-+-+-+-+-+-+-+-+-+-+-+-+-*/

type squeue struct {
	queue []interface{}
	lock  sync.RWMutex
}

//入队
func (q *squeue) Push(data interface{}) {
	q.lock.Lock()
	q.queue = append(q.queue, data)
	q.lock.Unlock()
}

//出队
func (q *squeue) Pop() (interface{}, error) {
	q.lock.Lock()
	if len(q.queue) == 0 {
		q.lock.Unlock()
		return nil, errors.New("Pop with empty queue.")
	} else {
		data := q.queue[0]
		q.queue = q.queue[1:]
		q.lock.Unlock()
		return data, nil
	}
}

//取队头
func (q *squeue) Head() (interface{}, error) {
	q.lock.RLock()
	if len(q.queue) == 0 {
		q.lock.RUnlock()
		return nil, fmt.Errorf("Empty queue error.")
	} else {
		data := q.queue[0]
		q.lock.RUnlock()
		return data, nil
	}
}

//队列长度
func (q *squeue) Len() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return len(q.queue)
}

//队列是否为空
func (q *squeue) Empty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if len(q.queue) > 0 {
		return false
	} else {
		return true
	}
}

func (q *squeue) String() string {
	if q.Empty() {
		return "Empty Queue."
	}
	str := ""
	q.lock.RLock()
	for _, v := range q.queue {
		str += fmt.Sprintf("%v ", v)
	}
	q.lock.RUnlock()
	buf := bytes.NewBufferString(strings.Repeat("_", len(str)))
	buf.WriteString("\n")
	buf.WriteString(str)
	buf.WriteString("\n")
	buf.WriteString(strings.Repeat("¯", len(str)))
	return buf.String()
}

//创建队列
func NewSQueue() *squeue {
	return &squeue{
		queue: []interface{}{},
	}
}

/*-+-+-+-+-+-+-+-+-+-+-+-+-+-+双端队列+-+-+-+-+-+-+-+-+-+-+-+-+-+-*/

type sdqueue struct {
	queue []interface{}
	lock  sync.RWMutex
}

//在队头插入
func (q *sdqueue) PushFront(data interface{}) {
	q.lock.Lock()
	q.queue = append([]interface{}{data}, q.queue...)
	q.lock.Unlock()
}

//在队尾插入
func (q *sdqueue) PushBack(data interface{}) {
	q.lock.Lock()
	q.queue = append(q.queue, data)
	q.lock.Unlock()
}

//从队头出队
func (q *sdqueue) PopFront() (interface{}, error) {
	q.lock.Lock()
	if len(q.queue) == 0 {
		q.lock.Unlock()
		return nil, errors.New("Pop with empty queue")
	} else {
		data := q.queue[0]
		q.queue = q.queue[1:]
		q.lock.Unlock()
		return data, nil
	}
}

//从队尾出队
func (q *sdqueue) PopBack() (interface{}, error) {
	q.lock.Lock()
	if len(q.queue) == 0 {
		q.lock.Unlock()
		return nil, errors.New("Pop with empty queue")
	} else {
		index := len(q.queue) - 1
		data := q.queue[index]
		q.queue = q.queue[0:index]
		q.lock.Unlock()
		return data, nil
	}
}

//取队头
func (q *sdqueue) Head() (interface{}, error) {
	q.lock.RLock()
	if len(q.queue) == 0 {
		q.lock.RUnlock()
		return nil, errors.New("Queue is Empty")
	} else {
		data := q.queue[0]
		q.lock.RUnlock()
		return data, nil
	}
}

//取队尾
func (q *sdqueue) Tail() (interface{}, error) {
	q.lock.RLock()
	if len(q.queue) == 0 {
		q.lock.RUnlock()
		return nil, errors.New("Queue is Empty")
	} else {
		index := len(q.queue) - 1
		data := q.queue[index]
		q.lock.RUnlock()
		return data, nil
	}
}

//队列长度
func (q *sdqueue) Len() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return len(q.queue)
}

//队列是否为空
func (q *sdqueue) Empty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()
	if len(q.queue) > 0 {
		return false
	} else {
		return true
	}
}

func (q *sdqueue) String() string {
	if q.Empty() {
		return "Empty Queue."
	}
	str := ""
	q.lock.RLock()
	for _, v := range q.queue {
		str += fmt.Sprintf("%v ", v)
	}
	q.lock.RUnlock()
	buf := bytes.NewBufferString(strings.Repeat("_", len(str)))
	buf.WriteString("\n")
	buf.WriteString(str)
	buf.WriteString("\n")
	buf.WriteString(strings.Repeat("¯", len(str)))
	return buf.String()
}

//创建双端队列
func NewSDQueue() *sdqueue {
	return &sdqueue{
		queue: []interface{}{},
	}
}

/*-+-+-+-+-+-+-+-+-+-+-+-+-+-+循环队列+-+-+-+-+-+-+-+-+-+-+-+-+-+-*/

//   1     2     3     □
//   ↑                 ↑
// head[0]            tail[3]

type scqueue struct {
	queue []interface{}
	head  int
	tail  int
	count int
	size  int
	lock  sync.RWMutex
}

//入队
func (q *scqueue) Push(data interface{}) error {
	q.lock.Lock()
	if q.count == q.size {
		q.lock.Unlock()
		return errors.New("Queue is full.")
	} else {

		q.queue[q.tail] = data
		q.tail = (q.tail + 1) % q.size
		q.count++
		q.lock.Unlock()
		return nil
	}
}

//出队
func (q *scqueue) Pop() (interface{}, error) {
	q.lock.Lock()
	if q.count == 0 {
		q.lock.Unlock()
		return nil, errors.New("Queue is empty.")
	} else {
		data := q.queue[q.head]
		q.head = (q.head + 1) % q.size
		q.count--
		q.lock.Unlock()
		return data, nil
	}
}

//取队头
func (q *scqueue) Head() (interface{}, error) {
	q.lock.RLock()
	if q.count == 0 {
		q.lock.RUnlock()
		return nil, errors.New("Queue is empty.")
	} else {
		data := q.queue[q.head]
		q.lock.RUnlock()
		return data, nil
	}
}

//队列长度
func (q *scqueue) Len() int {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.count
}

//队列是否为空
func (q *scqueue) Empty() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.count == 0
}

//队列是否已满
func (q *scqueue) Full() bool {
	q.lock.RLock()
	defer q.lock.RUnlock()
	return q.count == q.size
}

//重新设置循环队列大小
func (q *scqueue) Resize(newSize int) (int, error) {
	q.lock.RLock()
	if newSize <= q.count {
		q.lock.RUnlock()
		return q.size, errors.New("New size is too small.")
	} else {
		old := q.size
		nq := make([]interface{}, newSize, newSize)
		for i := 0; i < q.count; i++ {
			nq[i] = q.queue[(q.head+i)%q.size]
		}
		q.queue = nq
		q.head, q.tail, q.size = 0, q.count, newSize
		q.lock.RUnlock()
		return old, nil
	}
}

func (q *scqueue) String() string {
	if q.Empty() {
		return "Empty Queue."
	}
	str := ""
	q.lock.RLock()
	for i := 0; i < q.count; i++ {
		str += fmt.Sprintf("%v ", q.queue[(q.head+i)%q.size])
	}
	q.lock.RUnlock()

	buf := bytes.NewBufferString(strings.Repeat("_", len(str)))
	buf.WriteString("\n")
	buf.WriteString(str)
	buf.WriteString("\n")
	buf.WriteString(strings.Repeat("¯", len(str)))
	return buf.String()
}

//创建循环队列
func NewSCQueue(size int) (*scqueue, error) {
	if size <= 0 {
		return nil, fmt.Errorf("can't create loop queue with zero size.")
	}
	return &scqueue{
		queue: make([]interface{}, size, size),
		head:  0,
		tail:  0,
		count: 0,
		size:  size,
	}, nil
}
