package main
import "fmt"

func HelloWorld(){
  fmt.Println("Hello, World!")
}

func add(a int, b int){
  fmt.Println(a + b)
}

func plusPlus(a, b, c int) int {
  return a + b + c
}

func main(){
  fmt.Println("Testing Functions")
  fmt.Println("-- Calling Hello World")

  HelloWorld()

  fmt.Println("-- Calling Add")
  add(5, 32)

  fmt.Println("-- Calling plusPlus")
  fmt.Println(plusPlus(3, 2, 4))
}
