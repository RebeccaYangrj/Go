package main

import (
	"fmt"
)

// slice 切片在超过cap后会自己增加一个cap的长度再存储内容 slice是数组形式的
func main() {
	//本身设定slide cap长度
	var nums1 = make([]int, 3, 5)
	fmt.Printf("nums1 len =%d, cap =%d, value =%v\n", len(nums1), cap(nums1), nums1)
	nums1 = append(nums1, 1)
	fmt.Printf("nums1 len =%d, cap =%d, value =%v\n", len(nums1), cap(nums1), nums1)
	nums1 = append(nums1, 2)
	fmt.Printf("nums1 len =%d, cap =%d, value =%v\n", len(nums1), cap(nums1), nums1)
	nums1 = append(nums1, 3)
	fmt.Printf("nums1 len =%d, cap =%d, value =%v\n", len(nums1), cap(nums1), nums1)
	//不设定slide的cap长度，但是给她一些初始len，则以原始的len作为cap长度
	fmt.Println("-------------------")
	var nums2 = make([]int, 3)
	fmt.Printf("nums2 len =%d, cap =%d, value =%v\n", len(nums2), cap(nums2), nums2)
	nums2 = append(nums2, 1)
	fmt.Printf("nums2 len =%d, cap =%d, value =%v\n", len(nums2), cap(nums2), nums2)
	nums2 = append(nums2, 2)
	fmt.Printf("nums2 len =%d, cap =%d, value =%v\n", len(nums2), cap(nums2), nums2)
	nums2 = append(nums2, 3)
	//截取slice切片，如果改变截取的数值，原始数值也会发生改变
	fmt.Println("-------------------")
	s1 := nums1[:3]
	s2 := nums1[1:]
	fmt.Printf("s1 len =%d, cap =%d, value =%v\n", len(s1), cap(s1), s1)
	fmt.Printf("s2 len =%d, cap =%d, value =%v\n", len(s2), cap(s2), s2)
	s1[0] = 100
	fmt.Printf("s1 len =%d, cap =%d, value =%v\n", len(s1), cap(s1), s1)
	fmt.Printf("nums1 len =%d, cap =%d, value =%v\n", len(nums1), cap(nums1), nums1)

}
