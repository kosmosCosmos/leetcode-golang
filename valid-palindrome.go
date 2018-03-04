package main

import (
	"fmt"
	"strings"
)

func main(){
	b:=isPalindrome("abc")
	fmt.Println(b)
}

func isPalindrome(s string) bool {
	a := strings.ToLower(s)
	n, m := 0, len(a)-1
	for n < m {
		for n < m && !isChar(a[n]) {
			n++
		}
		for n < m && !isChar(a[m]) {
			m--
		}
		if a[n] != a[m] {
			return false
		}
		n++
		m--
	}
	return true
}

func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}


func isPalindrome2(s string) bool {
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for i < j && !isChar(s[i]) {
			i++
		}
		for i < j && !isChar(s[j]) {
			j--
		}
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}

	return true
}
