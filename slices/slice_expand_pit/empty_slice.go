package main

import "fmt"

//空切片
func main() {
	var errs []error
	errs = append(errs, nil)
	fmt.Println(errs)
	//nil也是一个值，可以对数组进行追加。需要注意的是实际使用时可能存在风险。
}
