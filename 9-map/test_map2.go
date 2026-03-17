package main

import "fmt"

func main() {
	cityMap := make(map[string]string)

	cityMap["beijing"] = "北京"
	cityMap["shanghai"] = "上海"
	cityMap["guangzhou"] = "广州"

	for key, value := range cityMap {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
	}

}
