package gollections_test

import (
	"sort"
	"testing"

	"github.com/michaelwomack/gollections"
	"github.com/stretchr/testify/require"
)

type Int int

func (i Int) Less(other Int) bool {
	return i < other
}

func TestMinHeap(t *testing.T) {
	mh := gollections.NewMinHeap[Int]([]Int{10, 12})
	expectedSize := 2
	require.Equal(t, expectedSize, mh.Size())
	require.Equal(t, "MinHeap(elements=[10 12])", mh.String())

	items := []int{2, 2, 2, 3, 5, 4, 1, 34}
	for _, item := range items {
		mh.Add(Int(item))
		expectedSize++
		require.Equal(t, expectedSize, mh.Size())
	}

	items = append([]int{10, 12}, items...)
	require.Equal(t, "MinHeap(elements=[1 2 3 2 2 10 5 12 4 34])", mh.String())
	sort.Ints(items)
	for _, item := range items {
		top, ok := mh.Peek()
		require.True(t, ok)
		require.Equal(t, Int(item), top)
		require.Equal(t, expectedSize, mh.Size())

		popped, ok := mh.Pop()
		require.True(t, ok)
		require.Equal(t, Int(item), popped)
		expectedSize--
		require.Equal(t, expectedSize, mh.Size())
	}

	_, ok := mh.Peek()
	require.False(t, ok)

	_, ok = mh.Pop()
	require.False(t, ok)

	require.Equal(t, "MinHeap(elements=[])", mh.String())
}

func TestMaxHeap(t *testing.T) {
	mh := gollections.NewMaxHeap([]Int{10, 12})
	expectedSize := 2
	require.Equal(t, expectedSize, mh.Size())
	require.Equal(t, "MaxHeap(elements=[12 10])", mh.String())

	items := []int{2, 2, 2, 3, 5, 4, 1, 34}
	for _, item := range items {
		mh.Add(Int(item))
		expectedSize++
		require.Equal(t, expectedSize, mh.Size())
	}

	items = append([]int{10, 12}, items...)
	require.Equal(t, "MaxHeap(elements=[34 12 5 4 10 2 3 2 1 2])", mh.String())
	sort.Ints(items)
	for i := len(items) - 1; i >= 0; i-- {
		item := items[i]
		top, ok := mh.Peek()
		require.True(t, ok)
		require.Equal(t, Int(item), top)
		require.Equal(t, expectedSize, mh.Size())

		popped, ok := mh.Pop()
		require.True(t, ok)
		require.Equal(t, Int(item), popped)
		expectedSize--
		require.Equal(t, expectedSize, mh.Size())
	}

	_, ok := mh.Peek()
	require.False(t, ok)

	_, ok = mh.Pop()
	require.False(t, ok)

	require.Equal(t, "MaxHeap(elements=[])", mh.String())
}
