package main

import "fmt"

func main(){
	b:=generate(3)
	fmt.Println(b)
}

func generate(numRows int) []int {
	res := []int{}
	res = append(res, 1)
	if numRows == 0 {
		return res
	}

	for i := 0; i < numRows; i++ {
		res = genNext(res)
	}

	return res
}

func genNext(p []int) []int {
	res := make([]int, 1, len(p)+1)
	res = append(res, p...)

	for i := 0; i < len(res)-1; i++ {
		res[i] += res[i+1]
	}

	return res
}