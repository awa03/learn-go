package main

import (
  "fmt"
)

// named ErrorVal
type ErrorVal struct{
  arg int 
  message string
}

func (e *ErrorVal) Error() {
  fmt.Printf("%d -- %s", e.arg, e.message)
}

func divide(a int, b int) (int, *ErrorVal){
  if (b == 0){
    return -1, &ErrorVal {}
  } else if (a == 1) {
    return -1, &ErrorVal {-1, "Undefined"}
  }
  return a/b, nil
}

func main() {
  num1:= 1
  num2:= 32
  num3:= 21
  num4:=0
  fmt.Printf("Testing: %d / %d\n", num1, num2)
  fmt.Println(divide(num1, num2))

  fmt.Printf("Testing: %d / %d\n", num1, num3)
  fmt.Println(divide(num1, num3))

  fmt.Printf("Testing: %d / %d\n", num1, num4)
  fmt.Println(divide(num1, num4))

  fmt.Printf("Testing: %d / %d\n", num2, num1)
  fmt.Println(divide(num2, num1))

  fmt.Printf("Testing: %d / %d\n", num2, num3)
  fmt.Println(divide(num2, num3))

  fmt.Printf("Testing: %d / %d\n", num2, num4)
  fmt.Println(divide(num2, num4))

  fmt.Printf("Testing: %d / %d\n", num3, num1)
  fmt.Println(divide(num3, num1))

  fmt.Printf("Testing: %d / %d\n", num3, num2)
  fmt.Println(divide(num3, num2))

  fmt.Printf("Testing: %d / %d\n", num3, num4)
  fmt.Println(divide(num3, num4))

  fmt.Printf("Testing: %d / %d\n", num4, num1)
  fmt.Println(divide(num4, num1))

  fmt.Printf("Testing: %d / %d\n", num4, num2)
  fmt.Println(divide(num4, num2))

  fmt.Printf("Testing: %d / %d\n", num4, num3)
  fmt.Println(divide(num4, num3))
}
