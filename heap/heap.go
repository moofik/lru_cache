package heap

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"
)

type Element struct {
	element  []byte
	priority int64
}

type MinHeap struct {
	nodes    []*Element
	heapSize int64
	capacity int64
}

func CreateHeap(capacity int64) *MinHeap {
	return &MinHeap{make([]*Element, capacity), 0, capacity}
}

func (h *MinHeap) parent(index int64) int64 {
	return (index - 1) / 2
}

func (h *MinHeap) left(index int64) int64 {
	return 2*index + 1
}

func (h *MinHeap) right(index int64) int64 {
	return 2*index + 2
}

func (h *MinHeap) swap(index int64, newIndex int64) {
	tempVal := h.nodes[index]
	h.nodes[index] = h.nodes[newIndex]
	h.nodes[newIndex] = tempVal
}

func (h *MinHeap) heapify(index int64) {
	l := h.left(index)
	r := h.right(index)
	smallest := index

	if l < h.heapSize && h.nodes[l].priority < h.nodes[index].priority {
		smallest = l
	}

	if r < h.heapSize && h.nodes[r].priority < h.nodes[smallest].priority {
		smallest = r
	}

	if smallest != index {
		h.swap(index, smallest)
		h.heapify(smallest)
	}
}

func (h *MinHeap) Insert(element []byte, priority int64) {
	h.heapSize++
	index := h.heapSize - 1
	h.nodes[index] = &Element{element: element, priority: priority}
	fmt.Println("Heap inserting: to index ", index, ", parent idx: ", h.parent(index), "; comparing ", h.nodes[h.parent(index)].priority, " and ",  h.nodes[index].priority )
	for index != 0 && h.nodes[h.parent(index)].priority > h.nodes[index].priority {
		fmt.Println("SWAP")
		h.swap(index, h.parent(index))
		index = h.parent(index)
	}
}

func (h *MinHeap) RemoveKey(key []byte) {
	extract := false

	for _, node := range h.nodes {
		if bytes.Equal(node.element, key) {
			node.priority = math.MinInt
			extract = true
			break
		}
	}

	if extract {
		h.heapify(0)
		h.ExtractMin()
	}
}

func (h *MinHeap) UpdateKeyPriority(key []byte, priority int64) {
	extract := false

	for _, node := range h.nodes {
		if bytes.Equal(node.element, key) {
			node.priority = priority
			extract = true
			break
		}
	}

	if extract {
		h.heapify(0)
	}
}

func (h *MinHeap) ExtractMin() []byte {
	if h.heapSize <= 0 {
		bs := make([]byte, 8)
		binary.BigEndian.PutUint64(bs, uint64(math.MaxUint64))
		return bs
	}

	if h.heapSize == 1 {
		h.heapSize--
		return h.nodes[0].element
	}

	root := h.nodes[0]
	h.nodes[0] = h.nodes[h.heapSize - 1]
	h.nodes[h.heapSize - 1] = nil
	h.heapSize--
	h.heapify(0)

	return root.element
}

func (h *MinHeap) Dump() {
	for _, i2 := range h.nodes {
		fmt.Printf("%d ", i2)
	}
}