package main

import (
	"math"
	"strconv"
	"fmt"
	"strings"
)

func main(){
	a:=revesre(-145)
	fmt.Println(a)
}

func revesre(n int)(int){
	if n>(-10)&&n<10{
		return n
	}
	n1:=math.Abs(float64(n))
	b:=strconv.Itoa(int(n1))
	if len(b)>10{
		return 0
	}
	v:=""
	for _,i:=range strings.Split(b,""){
		v=i+v
	}
	result,_:=strconv.Atoi(v)
	if n<0{
		result=-result
	}
	if result>-2147483648&&result<2147483648{
		return result
	}else {
		return 0
	}
}

func test(){
	a:="basdd"
	c:=""
	for _,i:=range strings.Split(a,""){
		c=i+c
	}
	fmt.Println(c)
}
