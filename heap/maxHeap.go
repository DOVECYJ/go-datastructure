package heap

import (
	"errors"
	"fmt"
	"sync"
)

type max_heap struct {
	heap []int
}

//创建大顶堆
func NewMaxHeap(data ...int) *max_heap {
	h := &max_heap{
		heap: []int{},
	}
	for i := range data {
		h.Put(data[i])
	}
	return h
}

//从切片创建大顶堆
func MaxHeapFromSlice(s []int) *max_heap {
	h := NewMaxHeap()
	for i := range s {
		h.Put(s[i])
	}
	return h
}

//添加到堆
func (h *max_heap) Put(data int) {
	if len(h.heap) == 0 {
		h.heap = append(h.heap, data)
	} else {
		h.heap = append(h.heap, data)
		child := len(h.heap) - 1
		parent := (child - 1) / 2
		for child > 0 && h.heap[child] > h.heap[parent] {
			h.heap[child], h.heap[parent] = h.heap[parent], h.heap[child]
			child = parent
			parent = (child - 1) / 2
		}
	}
}

//删除堆顶元素并返回
func (h *max_heap) Get() (data int, err error) {
	if len(h.heap) == 0 {
		data, err = 0, errors.New("Heap is empty.")
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
			if right < last && h.heap[right] > h.heap[left] {
				index = right
			}
			if h.heap[parent] < h.heap[index] {
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
func (h *max_heap) Top() (data int, err error) {
	if len(h.heap) == 0 {
		data, err = 0, errors.New("Heap is empty.")
	} else {
		data = h.heap[0]
	}
	return
}

//堆是否为空
func (h *max_heap) Empty() bool {
	return len(h.heap) == 0
}

//for test
func (h *max_heap) Print() {
	fmt.Println(h.heap)
}
