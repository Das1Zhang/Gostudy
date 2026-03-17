package main

import (
	"fmt"
)

func main() {
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 is nil")
	}

	// 没有初始化的话，必须要make分配空间，否则会报错：assignment to entry in nil map
	myMap1 = make(map[string]string, 10)
	myMap1["name"] = "Alice"
	myMap1["age"] = "30"
	myMap1["city"] = "New York"

	fmt.Println("myMap1:", myMap1)

	myMap2 := make(map[string]int)
	myMap2["one"] = 1
	myMap2["two"] = 2
	myMap2["three"] = 3

	fmt.Println("myMap2:", myMap2)

	myMap3 := map[string]string{
		"one":   "php",
		"two":   "java",
		"three": "golang", // 最后一行也需要逗号
	}
	fmt.Println("myMap3:", myMap3)
}
