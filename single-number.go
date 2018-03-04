package main

import "fmt"

func main(){
	a:=singleNumber([]int{1,2,2,1,4,5,5})
	fmt.Println(a)
}

func singleNumber(nums []int) int {
	res := 0
	for _, n := range nums {
		// n^n == 0
		// a^b^a^b^a == a
		fmt.Println(res,n)
		res ^= n
		fmt.Println(res)
	}
	return res
}