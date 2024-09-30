package main

import (
  "fmt"
)

func main(){
  nums := []int{1, 3, 4, 2, 1, 4}

  // _ is the index, not needed so wont include
  fmt.Println("Printing the num (itr)")
  for _, num := range nums {
    fmt.Println(num)
  }

  // i is the index now, no used iter
  fmt.Println("Printing the index")
  for i := range nums {
    fmt.Println(i)
  }

  fmt.Println("Printing both index and the num (itr)")
  for i, num := range nums {
    fmt.Println(i, num)
  }

}
