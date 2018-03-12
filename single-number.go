package main

import "fmt"

func main(){
	a:=singleNumber([]int{1,2,2,1,4,5,5})
	fmt.Println(a)
}

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		res ^= n
	}
	return res
}