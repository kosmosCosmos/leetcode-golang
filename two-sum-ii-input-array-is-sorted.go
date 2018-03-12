package main

import "fmt"

func main(){
	twoSum([]int{0,0,3,6},0)
	//fmt.Println(a)
}

func twoSum(nums []int,target int)[]int{
	res:=make([]int,2)
	b:=make(map[int]int,0)
	for i,x:=range nums{
		b[x]=i+1
	}
	for y,n:=range nums{
		num1:=target-n
		value,ok:=b[num1]
		if ok{
			res[0]=y+1
			res[1]=value
			break
		}
	}
	return res
}
