package main

import "fmt"

func main() {
	var s1 = []int{1, 2, 3, 4, 5, 6}
	var s2 []int

	var n = copy(s2, s1)
	fmt.Println(n)
}
func solve() []int {
	// Put your code here
	var a = []int{98, 87, 76}
	fmt.Print()
	return a
}
