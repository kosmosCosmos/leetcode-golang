package main

import "fmt"

func main(){
	str:="Hello wor ld "
	a:=lastword(str)
	fmt.Println(a)
}

func lastword(str string)int {
	if str == "" {
		return 0
	}
	empty:=0
	l:=len(str)
	for i:=l;i>0;i--{
		if str[i-1]!=uint8(32){
			break
		}
		empty++
	}
	count := 0
	fmt.Println(str[:l-empty])
	for _, x := range str[:l-empty] {
		count++
		if x == int32(32) {
			count = 0
		}
	}
	return count
}

func lastwordV2(str string) int {
	if str == "" {
		return 0
	}
	count := 0
	for i := len(str); i > 0; i-- {
		if str[i-1] == uint8(32) {
			if count != 0 {
				return count
			}
			continue
		}
		count++
	}

	return count
}