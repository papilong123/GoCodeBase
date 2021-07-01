package main

import "fmt"

func main() {
	//在dfs相关算法中会经常进行这种操作，需要额外注意。
	x := make([]int, 0, 10)
	x = append(x, 1, 2, 3)
	y := append(x, 4)
	z := append(x, 5)
	fmt.Println(x, y, z)

	//数组x的cap为10，在进行append操作后，len变为3。
	//再在x上进行 append 4 赋值给y，此时因为x的cap大于4，所以不需要进行扩容，会在x的基础上进行追加。即x、y实际用于存储的数组指向同一块内存。
	//此时 y为[1 2 3 4]，但此时x为[1 2 3]，因为决定切片的还有len，尽管内存中实际存储的还有4。
	//再在x上进行 append 5 赋值给z，x这块内存数组的第4位将存为5，覆盖原来的4（因为x的len为3，append的为第4位）。
	//这时候x依然受len影响不会变，z变为[1 2 3 5]。此时，x、y、z实际进行存储的数组指向了同一块内存。所以z进行append时会影响y，此时y也为[1 2 3 5]。
}
