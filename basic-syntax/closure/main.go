package main

import (
  "fmt"
)

// Return a function reference -- allows us to count using this
func intSeq() func() int {
  i := 0
  return func() int {
    i++
    return i
  }
}

func main(){
  nextInt := intSeq()

  // Continue Iteration
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())
  fmt.Println(nextInt())

  // New func 
  newNextInt := intSeq()
  fmt.Println(newNextInt())
  fmt.Println(newNextInt())
  fmt.Println(newNextInt())
  fmt.Println(newNextInt())
  fmt.Println(newNextInt())
  fmt.Println(newNextInt())
}
