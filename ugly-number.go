package main

import "fmt"

func main(){
	a:=isUgly(120)
	fmt.Println(a)
}

func isUgly(num int)  bool{
	if num<0{
		return false
	}
	if num<=6{
		return true
	}

	if num%2==0{
		return isUgly(num%2)
	}

	if  num%3==0{
		return isUgly(num%3)
	}

	if num%6==0{
		return isUgly(num%6)
	}

	return false
}