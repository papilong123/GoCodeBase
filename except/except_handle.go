package main

import "fmt"

//引发panic有两种情况，一是程序主动调用，二是程序产生运行时错误，由运行时检测并退出。
//发生panic后，程序会从调用panic的函数位置或发生panic的地方立即返回，逐层向上执行函数的defer语句，然后逐层打印函数调用堆栈，直到被recover捕获或运行到最外层函数。
//panic不但可以在函数正常流程中抛出，在defer逻辑里也可以再次调用panic或抛出panic。defer里面的panic能够被后续执行的defer捕获。
//recover用来捕获panic，阻止panic继续向上传递。
//recover()和defer一起使用，但是defer只有在后面的函数体内直接被掉用才能捕获panic来终止异常，否则返回nil，异常继续向外传递。

// DivErr 自定义错误信息结构
type DivErr struct {
	datatype int // 错误类型
	v1       int // 记录下出错时的除数、被除数
	v2       int
}

// 实现接口方法 error.Error()
func (divErr DivErr) Error() string {
	if 0 == divErr.datatype {
		return "除零错误"
	} else {
		return "其他未知错误"
	}
}

// 除法
func div(a int, b int) (int, *DivErr) {
	if b == 0 {
		// 返回错误信息
		return 0, &DivErr{0, a, b}
	} else {
		// 返回正确的商
		return a / b, nil
	}
}
func main() {
	// 正确调用
	v, r := div(100, 2)
	if nil != r {
		fmt.Println("(1)fail:", r)
	} else {
		fmt.Println("(1)succeed:", v)
	}
	// 错误调用
	v, r = div(100, 0)
	if nil != r {
		fmt.Println("(2)fail:", r)
	} else {
		fmt.Println("(2)succeed:", v)
	}
}
