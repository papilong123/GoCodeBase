package main

import "fmt"

//append后赋值给匿名变量
func main() {
	//append产生一个新切片，但是因为赋值给了匿名变量，所以不会影响原变量。
	nums := []int{1, 2, 3}
	_ = append(nums, 4)
	fmt.Println(nums)
}
