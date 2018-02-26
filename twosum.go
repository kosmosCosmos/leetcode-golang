package main

import "fmt"

func main(){
	num :=[]int{3,2,4,6,78,2}
	target := 80
	twoSum(num,target)
}

func twoSum(nums []int, target int)  {
	dict:=make(map[int]int,0)
	for x,i:= range(nums){
		_,ok:=dict[target-i]
		if ok!=true{
			dict[i]=x
		}else {
			fmt.Println(dict[target-i],x)
		}
	}
	fmt.Println(dict)
}
