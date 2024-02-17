package main

import "fmt"

func main() {
	var sum int = 0
	for i := 1; i <= 10; i++ {
		number := 0
		fmt.Scan(&number)
		sum += number
	}
	fmt.Println(sum)
}
