package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	a:=[]string{"ccsd","cc","ccc"}
	b:=longestCommonPrefix2(a)
	fmt.Println(b)
}

func longestCommonPrefix(a []string) string {
	b := ""
	num := 0
	if len(a) == 1 {
		b = a[0]
	} else if len(a) != 0 {
		sorts:=make(map[int]string,0)
		count:=make([]int,0)
		for _,x:=range a{
			sorts[len(x)]=x
			count=append(count,len(x))
		}
		sort.Ints(count)
		x := strings.Split(sorts[count[0]], "")
		y := strings.Split(sorts[count[len(count)-1]], "")
		if len(x) != 0 && len(y) != 0 {
			if len(x) > len(y) {
				num = len(y)
			} else {
				num = len(x)
			}
			for i := 0; i < num; i++ {
				if x[i] == y[i] {
					b = b + x[i]
				}
			}
		}
	}
	return b
}

func longestCommonPrefix2(strs []string) string {
	short := shortest(strs)
	for i, r := range short {
		for j := 0; j < len(strs); j++ {
			if strs[j][i] != byte(r) {
				return strs[j][:i]
			}
		}
	}
	return short
}

func shortest(strs []string) string {
	fmt.Println(strs)
	if len(strs) == 0 {
		return ""
	}

	res := strs[0]
	for _, s := range strs {
		if len(res) > len(s) {
			res = s
		}
	}

	return res
}