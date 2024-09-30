package main
import (
  "fmt"
)

func pow(a int, b int) int {
  if b == 0 {
    return 1;
  } else if b == 1 || a == 1 {
    return a
  }
  return pow(a * a, b-1)
}

func main(){
  fmt.Println("Recursively Calling Power Function")
  res := pow(1, 3)
  fmt.Println("1^3", res)

  res = pow(3, 1)
  fmt.Println("3^1", res)

  res = pow(3, 2)
  fmt.Println("3^2", res)

  res = pow(3, 4)
  fmt.Println("3^4", res)
  
}
