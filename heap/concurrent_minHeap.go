package heap

import (
	"errors"
	"fmt"
	"sync"
)

type min_heap struct {
	heap []int
	lock sync.RWMutex
}

//创建小顶堆
func NewMinHeap(data ...int) *min_heap {
	h := &min_heap{
		heap: []int{},
	}
	for i := range data {
		h.Put(data[i])
	}
	return h
}

//从切片构建小顶堆
func MinHeapFromSlice(s []int) *min_heap {
	h := NewMinHeap()
	for i := range s {
		h.Put(s[i])
	}
	return h
}

//添加到堆
func (h *min_heap) Put(data int) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if len(h.heap) == 0 {
		h.heap = append(h.heap, data)
	} else {
		h.heap = append(h.heap, data)
		child := len(h.heap) - 1
		parent := (child - 1) / 2
		for child > 0 && h.heap[child] < h.heap[parent] {
			h.heap[child], h.heap[parent] = h.heap[parent], h.heap[child]
			child = parent
			parent = (child - 1) / 2
		}
	}
}

//删除堆顶元素并返回
func (h *min_heap) Get() (data int, err error) {
	h.lock.Lock()
	defer h.lock.Unlock()
	if len(h.heap) == 0 {
		data, err = 0, errors.New("Heap is empty.")
		return
	} else {
		data = h.heap[0]
		last := len(h.heap) - 1
		h.heap[0] = h.heap[last]
		h.heap = h.heap[0:last]
		parent := 0
		for {
			left, right := 2*parent+1, 2*parent+2
			if left >= last {
				break
			}
			index := left
			if right < last && h.heap[right] < h.heap[left] {
				index = right
			}
			if h.heap[parent] > h.heap[index] {
				h.heap[parent], h.heap[index] = h.heap[index], h.heap[parent]
			} else {
				break
			}
			parent = index
		}
	}
	return
}

//获取堆顶元素
func (h *min_heap) Top() (data int, err error) {
	h.lock.RLock()
	if len(h.heap) == 0 {
		h.lock.RUnlock()
		data, err = 0, errors.New("Heap is empty.")
	} else {
		data = h.heap[0]
		h.lock.RUnlock()
	}
	return
}

//堆是否为空
func (h *min_heap) Empty() bool {
	h.lock.RLock()
	empty := (len(h.heap) == 0)
	h.lock.RUnlock()
	return empty
}

//for test
func (h *min_heap) Print() {
	h.lock.RLock()
	fmt.Println(h.heap)
	h.lock.RUnlock()
}
