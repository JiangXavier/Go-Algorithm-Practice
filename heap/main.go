package main

import "fmt"

func main() {
	nums := []int{2, 1, 2, 6, 5, 6, 58, 8, 6}
	heapify(nums)
	for len(nums) > 0{
		fmt.Println(pop(&nums))
	}
}

func heapify(nums []int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		down(nums, i, n)
	}
}

func down(nums []int, i, n int) {
	parent := i
	for {
		left := 2*parent + 1
		if left >= n || left < 1{
			break
		}
		max := left
		if right := left + 1; right < n && nums[right] > nums[left] {
			max = right
		}
		if nums[parent] > nums[max] {
			break
		}
		nums[parent], nums[max] = nums[max], nums[parent]
		parent = max
	}
}


func pop(nums *[]int) int {
	last := len(*nums) - 1
	(*nums)[0] , (*nums)[last] = (*nums)[last] , (*nums)[0]
	down(*nums , 0 , last)
	ans := (*nums)[last]
	(*nums) = (*nums)[:last]
	return ans
}

