package main

import "fmt"

func create3DArray() interface{} {
	var array = [4][4][4]float32{
		1: {
			0: {
				2: 88.6,
			},
		},
	}
	fmt.Print()
	var a = [4]int{}

	var a1 [4]int = [4]int{}

	var a2 = [4]int

	a3 := [4]int{}

	var a4 [4]int

	a5 := [4]int
	fmt.Print(a,a1,a2,a4,a5,a3)
	return array
}
