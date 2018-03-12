package main

import (
	"fmt"
)

func main(){
	a:=trailingZeroes(29)
	fmt.Println(a)
}

func trailingZeroes(n int) int {
	res := 0

	for n >= 5 {
		n /= 5
		res += n
	}

	return res
}