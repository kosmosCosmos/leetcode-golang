package main

import "fmt"

func main(){
	b:=addDigits(231)
	fmt.Println(b)
}

func addDigits(n int) int {
	return (n-1)%9 + 1
}