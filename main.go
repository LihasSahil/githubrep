package main

import "fmt"

func main() {
	// var m map[int]int // int:int
	m := make(map[int]int)

	m[1] = 3
	m[2] = 2
	m[3] = 0

	fmt.Println(m)

	fmt.Println(m[1])
	val, ok := m[5]
	if ok {
		fmt.Println(val)
	} else {
		fmt.Println("Invalid Key")
	}

	m[3] = 100
	fmt.Println(m)

	// delete(m, 3)
	// fmt.Println(m)

	for k, v := range m {
		if v == 100 {
			fmt.Println("Century at ", k)
			break
		}
	}

}
