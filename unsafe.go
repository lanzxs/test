package main

import "fmt"

func double(x int) {
	x += x
}
func main() {
	var a = 3
	double(a)
	fmt.Println(a) // 3
}
