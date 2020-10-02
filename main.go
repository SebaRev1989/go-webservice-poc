package main

import "fmt"

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	fmt.Println(arr[1])

	newArr := [3]int{3, 4, 5}
	fmt.Println(newArr)
}
