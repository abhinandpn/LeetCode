package main

import (
	"fmt"

	Golang "github.com/abhinandpn/LeetCode/Golang"
)

func main() {
	// Solution for 169
	var ar = []int{2, 3, 4, 2, 2, 3, 7, 3, 8, 3, 2}
	res := Golang.RemoveDuplicates(ar)
	fmt.Println(res)
}
