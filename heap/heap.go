package main

import (
	"container/heap"
	"fmt"
)

type hp32 []int32

func (h hp32) Len() int           { return len(h) }
func (h hp32) Less(i, j int) bool { return h[i] > h[j] } // > 为最大堆
func (h hp32) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp32) Push(v any)        { *h = append(*h, v.(int32)) }
func (h *hp32) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp32) push(v int32)      { heap.Push(h, v) }
func (h *hp32) pop() int32        { return heap.Pop(h).(int32) } // 稍微封装一下，方便使用

func main() {
	// h := &hp32{1,4,2,8,6}
	// heap.Init(h)
	// for h.Len() > 0{
	//     fmt.Println(h.pop())
	// }
	nums := []int{2, 4, 1, 3, 6, 9, 5}
	heapify(nums)
	for len(nums) > 0 {
		fmt.Println(pop(&nums))
	}
}

func heapify(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		min_down(nums, i, n)
		//max_down(nums, i, n)
	}
}

func max_down(nums []int, i, n int) bool {
	parent := i
	for {
		left := 2*parent + 1
		if left >= n || left < 1 {
			break
		}
		max := left
		if right := left + 1; right < n && nums[right] > nums[left] {
			max = right
		}
		if nums[max] < nums[parent] {
			break
		}
		nums[parent], nums[max] = nums[max], nums[parent]
		parent = max
	}
	return parent > i
}

func pop(nums *[]int) int {
	last := len(*nums) - 1
	(*nums)[0], (*nums)[last] = (*nums)[last], (*nums)[0]
	min_down(*nums, 0, last)
	//max_down(*nums, 0, last)
	rst := (*nums)[last]
	(*nums) = (*nums)[:last]
	return rst
}

func min_down(nums []int, i, n int) bool {
	parent := i
	for {
		left := 2*parent + 1
		if left >= n || left < 1 {
			break
		}
		max := left
		if right := left + 1; right < n && nums[right] < nums[left] {
			max = right
		}
		if nums[max] > nums[parent] {
			break
		}
		nums[parent], nums[max] = nums[max], nums[parent]
		parent = max
	}
	return parent > i
}
