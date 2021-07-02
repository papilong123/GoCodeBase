package main

import (
	"fmt"
	"time"
)

func say3(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, (i+1)*100)
	}
}
func say4(s string, ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(s, (i+1)*150)
	}
	ch <- 0
	close(ch)
}

func main() {
	ch := make(chan int)
	go say4("world", ch)
	say3("hello")
	fmt.Println(<-ch)
}
