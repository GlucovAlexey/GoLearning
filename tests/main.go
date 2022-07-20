package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(n)
	//OK
	if n > 0 && n <= 10000 {
		n = n * n
	}
}
