package main

import "fmt"

func mySqrt(x int) int {
	res := x

	// 牛顿法求平方根
	for res*res > x {
		res = (res + x/res) / 2
	}

	return res
}

func main(){
	c:=mySqrt(180)
	fmt.Print(c)
}
