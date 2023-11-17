package main

import (
	"fmt"
	"container/heap"
)

func main() {
	nums := []int{1,1,2,2,2,3,3,3,3,3,4,5,6,6,6,6,6,6,6,6}
	m := map[int]int{}
	for _ , v := range nums{
		m[v]++
	}
	fmt.Println(m)
	h := &my{}
	heap.Init(h)
	for i , v := range m{
		heap.Push(h , [2]int{i,v})
	}
	for h.Len() > 0{
		fmt.Println(heap.Pop(h).([2]int)[0])
	}
	// nums := []int{1, 3, 2, 4, 6, 9, 5}
	// heapify(nums)
	// for len(nums) > 0 {
	// 	fmt.Println(pop(&nums))
	// }
}

type my [][2]int

func (h my) Len() int           { return len(h) }
func (h my) Less(i, j int) bool { return h[i][1] > h[j][1] }
func (h my) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *my) Push(x any) {
	*h = append(*h, x.([2]int))
}
func (h *my) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func heapify(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		down(nums, i, n)
	}
}
func down(nums []int, i, n int) {
	parent := i 
	for{
		left := 2 * parent + 1
		if left >= n || left < 1{break}
		max := left
		if right := left + 1; right < n && nums[left] > nums[right]{
			max = right
		}
		if nums[parent] < nums[max]{break}
		nums[parent] , nums[max] = nums[max] , nums[parent]
		parent = max
	}
}
func pop(nums *[]int) int {
	last := len(*nums) - 1
	(*nums)[0], (*nums)[last] = (*nums)[last], (*nums)[0]
	down(*nums, 0, last)
	ans := (*nums)[last]
	(*nums) = (*nums)[:last]
	return ans
}
