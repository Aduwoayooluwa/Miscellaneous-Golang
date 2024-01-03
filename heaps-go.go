package main

import (
	"fmt"
)

// heaps can be visualized as tree. theay are an array that operates as a tree.

// maxheap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// insert method adds an element to the heap.
func (h *MaxHeap) Insert(key int) {
		// adding the element to the array
		h.array = append(h.array, key)
		h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	if len(h.array) == 0 {
		fmt.Println("cannot extract array is length 0")
		return -1

	}
	return extracted
}

// maxheapifyup will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}

}
//  maxheapify Down 
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l,r := left(index), right(index)
	childToCompare := 0

		//  loop while index has at least one child
	for l <= lastIndex {
		//  when left child is the only child only one child
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			//  when left child is larger
			childToCompare = l
		} else {
			// whe right child is larger
			childToCompare = r
		}
		// compare array vale of current index to larger child swap if smaller 
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return 
		}
	}


}

// get parent index
func parent(i int) int {
	return (i-1)/2
}

//  get left child index
func left (i int) int {
	return 2*i + 1
}

// get the right child index
func right (i int) int {
	return 2 * i + 2
}

//  swap keys on the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}


func main() {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	for _,v :=range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

	for i:=0; i< 5; i++ {
		m.Extract()
		fmt.Println(m)
	}

}