package main

import "fmt"

func main(){
	a:=[]int{1,2,4,5}
	b:=[]int{1,3,6,8,9}
	merge(a,4,b,5)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	// 深度复制 nums1
	temp := make([]int, m)
	copy(temp, nums1)

	j, k := 0, 0
	for i := 0; i < len(nums1); i++ {
		// nums2用完了
		if k >= n {
			nums1[i] = temp[j]
			j++
			continue
		}
		// temp 用完了
		if j >= m {
			nums1[i] = nums2[k]
			k++
			continue
		}
		// 比较后，放入
		if temp[j] < nums2[k] {
			nums1[i] = temp[j]
			j++
		} else {
			nums1[i] = nums2[k]
			k++
		}
	}

}