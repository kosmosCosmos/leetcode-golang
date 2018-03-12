package main

import (
	"fmt"
)

func main(){
	a:=isAnagram("a","ab")
	fmt.Println(a)
}

func validAnagram(s,l string)(bool){
	var s1,s2 string
	if len(s)>len(l){
		s1=s
		s2=l
	}else {
		s1=l
		s2=s
	}
	a:=trans(s1)
	b:=trans(s2)
	for k,v:=range a{
		v2,ok:=b[k]
		if ok!=true||v2!=v{
			return false
		}
	}
	return true
}

func trans(s string)map[int32]int{
	res:=make(map[int32]int,0)
	for _,x:=range s{
		_,ok:=res[x]
		if ok{
			res[x]++
		}else {
			res[x]=1
		}
	}
	return res
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// 把 string 转换成 []rune 可以适应 Unicode 字符
	sr := []rune(s)
	tr := []rune(t)

	// 因为使用了 []rune，rec 只好使用 map
	// 不然的话，使用 [26]int 数组，效率更高
	rec := make(map[rune]int, len(sr))
	for i := range sr {
		rec[sr[i]]++
		rec[tr[i]]--
	}

	for _, n := range rec {
		if n != 0 {
			return false
		}
	}

	return true
}