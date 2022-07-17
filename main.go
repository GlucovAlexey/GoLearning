package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)
	first := num / 100000
	seconf := num / 10000 % 10
	third := num / 1000 % 10
	forth := num / 100 % 10
	fifth := num / 10 % 10
	sixth := num % 10
	switch {
	case (first + seconf + third) == (forth + fifth + sixth):
		fmt.Print("YES")
	default:
		fmt.Print("NO")
	}
}
