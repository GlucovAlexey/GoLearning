package main

import "fmt"

func main() {

  var count int
  fmt.Scan(&count)
  workArray := []int{}
  var a int
	
  for i := 0; i <= count; i++ {
	fmt.Scan(&a)
    workArray = append(workArray, a)
  }
  for i := 0; i < count; i++{
    if workArray[i] % 2 == 0{
      fmt.Printf(" %d ", workArray[i])
    }
  }
}
