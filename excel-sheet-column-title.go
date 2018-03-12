package main

import "fmt"

func main(){
	b:=convertToTitle(289678)
	c:=titleToNumber(b)
	fmt.Println(c)
}

func convertToTitle(n int) string {
	res := ""

	for n > 0 {
		n--
		res = string(byte(n%26)+'A') + res
		n /= 26
	}

	return res
}

func titleToNumber(s string) int {
	res := 0

	for i := 0; i < len(s); i++ {
		temp := int(s[i] - 'A' + 1)
		res = res*26 + temp
	}

	return res
}