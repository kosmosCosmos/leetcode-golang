package main

import "fmt"

func maxProfit(prices []int) int {
	res := 0

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			res += prices[i] - prices[i-1]
		}
	}

	return res
}

func main(){
	c:=maxProfit([]int{1,2,6,2,4,6,3,4})
	fmt.Println(c)
}