package main

import "fmt"

func main() {
	workArray := [10]uint8{}
	var a uint8
	for i := 0; i < 10; i++ {
		fmt.Scan(&a)
		workArray[i] = a
	}
	array := [6]uint8{}
	var b uint8
	for i := 0; i < 6; i++ {
		fmt.Scan(&b)
		array[i] = b
	}
	workArrayCopy := workArray
	for i := 0; i < 5; i += 2 {
		workArray[array[i]] = workArrayCopy[array[i+1]]
		workArray[array[i+1]] = workArrayCopy[array[i]]
	}
	for i := 0; i < 10; i++ {
		fmt.Print(workArray[i], " ")
	}
}
