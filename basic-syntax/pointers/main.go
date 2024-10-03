package main

import (
  "fmt"
)

func zeroval(ival int) {
  ival = 0
}

func zeroptr(iptr *int) {
  *iptr = 0
}

func main(){
  val := 5
  fmt.Println("Before Running Zero Val:", val)
  zeroval(val)
  fmt.Println("After Running Zero Val:", val)

  fmt.Println("Before Running Zero Ptr:", val)
  zeroptr(&val)
  fmt.Println("After Running Zero Ptr:", val)
  fmt.Println("Pointer:", &val)
}
