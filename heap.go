package gollections

import (
	"fmt"
)

type Lesser[T any] interface {
	Less(other T) bool
}

type MinHeap[T Lesser[T]] struct {
	elements []T
}

func NewMinHeap[T Lesser[T]](items []T) *MinHeap[T] {
	h := &MinHeap[T]{
		elements: make([]T, 1, len(items)+1),
	}
	h.Add(items...)
	return h
}

func (h *MinHeap[T]) Size() int {
	return len(h.elements) - 1
}

// Add will insert each element to the heap in O(logn) time.
// If there are k elements, it will be O(k * logn) time.
func (h *MinHeap[T]) Add(elements ...T) {
	for _, e := range elements {
		h.add(e)
	}
}

func (h *MinHeap[T]) add(element T) {
	h.elements = append(h.elements, element)
	index := h.Size()
	parentIdx := index / 2
	for h.elements[index].Less(h.elements[parentIdx]) && index > 1 {
		h.elements[index], h.elements[parentIdx] = h.elements[parentIdx], h.elements[index]
		index = parentIdx
		parentIdx = index / 2
	}
}

func (h *MinHeap[T]) Peek() (T, bool) {
	if h.Size() == 0 {
		return h.elements[0], false
	}

	return h.elements[1], true
}

func (h *MinHeap[T]) Pop() (T, bool) {
	if h.Size() == 0 {
		return h.elements[0], false
	}

	// Move bottom right element (last item in slice) to place of root element to delete (first item).
	root := 1
	removed := h.elements[root]
	h.elements[root] = h.elements[h.Size()]
	h.elements = h.elements[:h.Size()]

	// When the newly placed root element is a parent, check if left or right child is
	// less. If the left and right child are less, the smaller one will be swapped with the parent.
	parent := root
	for parent <= h.Size()/2 {
		left := parent * 2
		right := left + 1
		hasRight := right <= h.Size()
		shouldSwapLeft := h.elements[left].Less(h.elements[parent])
		shouldSwapRight := hasRight && h.elements[right].Less(h.elements[parent])

		if !(shouldSwapLeft || shouldSwapRight) {
			break
		}

		if hasRight && h.elements[right].Less(h.elements[left]) {
			h.elements[right], h.elements[parent] = h.elements[parent], h.elements[right]
			parent = right
		} else {
			h.elements[left], h.elements[parent] = h.elements[parent], h.elements[left]
			parent = left
		}
	}

	return removed, true
}

func (h *MinHeap[T]) String() string {
	return fmt.Sprintf("MinHeap(elements=%v)", h.elements[1:])
}

type MaxHeap[T Lesser[T]] struct {
	elements []T
}

func NewMaxHeap[T Lesser[T]](items []T) *MaxHeap[T] {
	h := &MaxHeap[T]{
		elements: make([]T, 1, len(items)+1),
	}
	h.Add(items...)
	return h
}

func (h *MaxHeap[T]) Size() int {
	return len(h.elements) - 1
}

func (h *MaxHeap[T]) Peek() (T, bool) {
	if h.Size() == 0 {
		return h.elements[0], false
	}

	return h.elements[1], true
}

func (h *MaxHeap[T]) Add(elements ...T) {
	for _, e := range elements {
		h.add(e)
	}
}

func (h *MaxHeap[T]) add(element T) {
	h.elements = append(h.elements, element)
	index := h.Size()
	parent := index / 2
	for index > 1 && h.elements[parent].Less(h.elements[index]) {
		h.elements[parent], h.elements[index] = h.elements[index], h.elements[parent]
		index = parent
		parent = index / 2
	}
}

func (h *MaxHeap[T]) Pop() (T, bool) {
	if h.Size() == 0 {
		return h.elements[0], false
	}

	root := 1
	removed := h.elements[root]
	h.elements[root] = h.elements[h.Size()]
	h.elements = h.elements[:h.Size()]

	parent := root
	for parent <= h.Size()/2 {
		left := parent * 2
		right := left + 1
		hasRight := right <= h.Size()
		shouldSwapLeft := h.elements[parent].Less(h.elements[left])
		shouldSwapRight := hasRight && h.elements[parent].Less(h.elements[right])

		if !(shouldSwapLeft || shouldSwapRight) {
			break
		}

		if hasRight && h.elements[left].Less(h.elements[right]) {
			h.elements[parent], h.elements[right] = h.elements[right], h.elements[parent]
			parent = right
		} else {
			h.elements[parent], h.elements[left] = h.elements[left], h.elements[parent]
			parent = left
		}
	}

	return removed, true
}

func (h *MaxHeap[T]) String() string {
	return fmt.Sprintf("MaxHeap(elements=%v)", h.elements[1:])
}
