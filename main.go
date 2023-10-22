package main

import (
	"fmt"

	Golang "github.com/abhinandpn/LeetCode/Golang"
)

func main() {
	var ar = []int{2, 3, 4, 2, 2, 3, 7, 3, 8, 3,2}
	res := Golang.MajorityElement(ar)
	fmt.Println(res)
}
