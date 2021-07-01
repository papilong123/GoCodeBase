package main

import (
	"fmt"
)

func main() {
	//ws := new (restful.WebService)
	//ws.Route(ws.GET("/hello").To(hello))
	//restful.Add(ws)
	//log.Fatal(http.ListenAndServe(":18080", nil))
	var i int
	var f float64
	var b bool
	var s string
	var ab map[string]int
	var ac = 12
	println(ab)
	println(ac)
	fmt.Printf("%v %v %v %q\n\n", i, f, b, s)
	g, h := 123, "hello"
	fmt.Println(g, h)
}
