package main
import (
  "fmt"
)

func main() {
  // Declare an int array
  var arr [5]int
  fmt.Println(arr)

  // print 100
  arr[4] = 100
  fmt.Println("array: ", arr[4])

  // Get the length of the array
  fmt.Println("len: ", len(arr))


  a := [2][3]int{
    {1, 2, 3},
    {1, 2, 3},
    }
  fmt.Println("2d:", a)

  b := [5]int{1, 2, 3, 4, 5}
  fmt.Println("dcl", b)

  c := [...]int{1, 2, 3, 4, 5}
  fmt.Println("dcl", c)

  d := [...]int{100, 3: 400, 500}
  fmt.Println("idx", d)
  //
}
