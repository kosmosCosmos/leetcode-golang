package main

import "fmt"

func main(){
	nums := []int{1,3,3,3,3,4,5,8}
	target:=3
	a:=removeDuplicates(nums,target)
	fmt.Println(a)
}

func removeDuplicates(nums []int,target int) int {
	res := 0
	i := 0
	for ; i < len(nums); i++ {
		if nums[i] == target {
			continue
		}
			nums[res] = nums[i]
		res++
	}

	return res
}
