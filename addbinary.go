package main

import (
	"strconv"
	"strings"
	"fmt"
)

func main(){
	a:=binary("10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101","110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011")
	fmt.Println(a)
}

func binary(a,b string)string{
	num1,_:=strconv.Atoi(a)
	num2,_:=strconv.Atoi(b)
	res:=num1+num2
	fmt.Println(res)
	var result string
	str:=strconv.Itoa(res)
	array:=strings.Split(str,"")
	intarray:=make([]int,0)
	for _,x:=range array{
		xint,_:=strconv.Atoi(x)
		intarray=append(intarray,xint)
	}
	for i:=len(intarray)-1;i>0;i--{
		if intarray[i]>1{
			intarray[i]=intarray[i]-2
			intarray[i-1]++
		}

	}
	if intarray[0]>1{
		intarray[0]=intarray[0]-2
		intarray=append([]int{1},intarray...)
	}
	for _,y:=range intarray{
		result=result+strconv.Itoa(y)
	}
	return result
}