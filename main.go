package main

import "fmt"

func main() {
	var num, res, greate int
	num = 1
	for num != 0 {
		fmt.Scan(&num)
		if greate < num {
			greate = num
			res = 0
		}
		if greate == num {
			res++
		}
	}
	fmt.Print(res)
}
