package main

import "fmt"

/*
	- We will be checking if the number is power of 2
*/

func CheckIfNumberIsPowerOf2(num int) bool {
	return (num <= 0) && (num&(num-1) == 0)
}

func main() {
	var num int = 0
	if CheckIfNumberIsPowerOf2(num) {
		fmt.Printf("%d is power of 2", num)
	} else {
		fmt.Printf("%d is not power of 2", num)
	}
}
