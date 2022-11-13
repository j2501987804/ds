package main

import "fmt"

func main() {
	a := []int{0, 1, 2, 3, 4, 5, 6}
	b := make([]int,0,3)
	a[6] = 0
	copy(b, a[7:])
	fmt.Printf("%+v", b)
}
