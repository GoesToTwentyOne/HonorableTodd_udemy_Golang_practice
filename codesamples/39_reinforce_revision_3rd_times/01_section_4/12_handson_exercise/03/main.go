package main

import "fmt"

var x int = 42
var y bool = true
var z string = "James Bond"

func main() {
	s := fmt.Sprintf("%v \t %v \t %v\n", x, y, z)
	fmt.Println(s)

}
