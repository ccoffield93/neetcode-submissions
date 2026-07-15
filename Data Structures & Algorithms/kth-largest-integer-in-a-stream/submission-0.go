
// Simple heap implementation using container/heap
// An IntHeap is a min-heap of ints.
type IntHeap []int
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // root will always be less than children
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}
func (h *IntHeap) Peek() any {
	arr := *h 
	return arr[0]
} 
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}


type KthLargest struct {
    k int
	minheap IntHeap
}


func Constructor(k int, nums []int) KthLargest {
    obj := KthLargest{}
	obj.k = k
	obj.minheap = nums

	// initialize the heap with the values we start with
	heap.Init(&obj.minheap)

    // we only care about the K largest numbers so pop off 
	// the smallest ones until that's all that's left 
	for len(obj.minheap) > k {
		heap.Pop(&obj.minheap)
	}

	return obj
}

func (this *KthLargest) Add(val int) int {
    heap.Push(&this.minheap, val)

	// pop off the smallest elements until we only have K left
	for len(this.minheap) > this.k {
		heap.Pop(&this.minheap)
	}

	// The ROOT of the tree will be the SMALLEST of the tree
	// But, we only keep the top K largest elements in the tree
	// So, the root is the smallest of the K largest, i.e.
	// the K largest!
	return this.minheap.Peek().(int)
}


