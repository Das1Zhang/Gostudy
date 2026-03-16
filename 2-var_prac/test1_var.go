package main

import "fmt"

/*
go 的多变量声明
*/
var x, y int
var (
	a int
	b bool
)

var z, h = 1, "hello"

func main() {
	var c, d int = 1, 2
	e, f := 3, "zhangsihao"
	fmt.Println(x, y, a, b, z, h, c, d, e, f)
}
