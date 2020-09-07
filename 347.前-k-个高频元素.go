package leetcode

import (
	"container/heap"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	occurences := map[int]int{}

	for _, num := range nums {
		occurences[num]++
	}

	h := &IHeap{}
	heap.Init(h)

	for k, v := range occurences {
		heap.Push(h, [2]int{k, v})

		if h.Len() > k {
			heap.Pop(h)
		}
	}

	ret := make([]int, k)
	fmt.Println("ret: ", len(ret))

	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}

	return ret
}

// IHeap is
type IHeap [][2]int

// Len is
func (h IHeap) Len() int {
	return len(h)
}

// Less is
func (h IHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}

// Swap is
func (h IHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Push is
func (h *IHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

// Pop is
func (h *IHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end
