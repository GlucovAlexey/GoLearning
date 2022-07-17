package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	switch {
	case 9 < num && num < 100:
		fmt.Print(num / 10)
	case 99 < num && num < 1000:
		fmt.Print(num / 100)
	case 999 < num && num < 10000:
		fmt.Print(num / 1000)
	default:
		fmt.Print(num)
	}
}
