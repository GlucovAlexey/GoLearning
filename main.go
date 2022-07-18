package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b)
	for a > 0 {
		d = a % 10
		a = a / 10
		c = b
		for c > 0 {
			if c%10 == d {
				fmt.Print(d, " ")
				break
			}
			c = c / 10
		}
	}
}
