package main

import (
	"fmt"
	"time"
)

func main() {
	f1(1, 2, 3, 4)
}

func f1(a int, b ...int) (sum int) {
	fmt.Println(a, b)
	for _, val := range b {
		sum = sum + val
	}
	fmt.Println(sum, b)
	fmt.Println(time.Now())
	return
}
