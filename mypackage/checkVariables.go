package mypackage

import "fmt"

func Variables() {
	var num int = 0
	if num > 0 {
		fmt.Println("Num is positive number")
	} else if num == 0 {
		fmt.Println("Num is zero")
	} else if num < 0 {
		fmt.Println("Num is negative")
	}
}
