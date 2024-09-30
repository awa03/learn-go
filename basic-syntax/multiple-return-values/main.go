package main
import "fmt"

// similar to c# 
// (int, int) funcName(){}
func vals() (int, int) {
  return 4, 1
}

func main(){
  a, b := vals()
  fmt.Println("a val: ", a)
  fmt.Println("b val: ", b)

  _, c := vals()
  fmt.Println("underscore test (only get c): ", c)
}
