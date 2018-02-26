package main

import (
	"strconv"
	"strings"
)

func main() {
	palindrome2(1123)
	//a := palindrome2(11)
	//fmt.Println(a)
}

func palindrome(b int) (bool) {
	a := 0
	c := b
	if b < 10 && b >= 0 {
		return true
	}
	if b < 0 {
		return false
	}
	for {
		a = a*10 + b%10
		b = b / 10
		if b == 0 {
			break
		}
	}
	if a == c {
		return true
	} else {
		return false
	}
}

func palindrome2(x int)(bool){
	if x<0{
		return false
	}
	b:=strconv.Itoa(x)
	c:=strings.Split(b,"")
	for i,j:=0,len(c)-1;i<j;i,j=i+1,j-1{
		if c[i]!=c[j]{
			return false
		}
	}
	return true
}