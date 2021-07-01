package main

import "fmt"

//函数中的切片扩容
func main() {
	nums1 := []int{1, 2, 3}
	Append(nums1, 4)
	fmt.Println(nums1)
	//当切片存储容量不足时，则会把当前切片进行扩容并产生新的切片。新的切片具有新的地址，但是go是按值传递，所以新的切片地址只是局部变量，是不能够随着函数返回的。
	//若切片容量足够，切片作为函数传参扩容后返回的还是原来的切片内容。因为决定切片的还有len。

	nums2 := make([]int, 3, 4)
	nums2[2] = 2
	Set(nums2)
	fmt.Println(nums2)
	//切片作为参数传递时，传递的本质还是指向实际存储的数组的地址，所以还是能够改变值的。只是在进行append操作时需要注意。
}

func Append(nums []int, a int) {
	nums = append(nums, a)
}

func Set(nums []int) {
	nums[0] = 999
}
