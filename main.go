package main

import "fmt"

func main() {
	var num string
	fmt.Scan(&num)
	var a = num[0]
	var b = num[1]
	var c = num[2]

	switch {
	case a != b && b != c && c != a:
		fmt.Print("YES")
	default:
		fmt.Print("NO")
	}
}
