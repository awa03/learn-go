package main

import "fmt"

func sum(nums ... int){
  fmt.Println(nums)
  total := 0

  for _, nums := range nums {
    total += nums
  }
  fmt.Println("total: ", total)
}

func main(){
  sum(3, 32, 32, 1, 4, 2, 1, 5)
  sum(2, 3, 1)

  // pass array through variadic
  nums := []int{1, 3, 4, 1}
  sum(nums...)
}
