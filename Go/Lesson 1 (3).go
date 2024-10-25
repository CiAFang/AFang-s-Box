package main

import (
	"fmt"
)

func main() {
	var x int
	fmt.Scanln(&x)
	if x <= 1 {
		fmt.Println("No")
	} else {
		for j := 2; j < (x - 1); j++ {
			flag := x % j
			if flag == 0 {
				fmt.Println("No")
				break
			} else {
				fmt.Println("Yes")
				break
			}
		}
	}
}
