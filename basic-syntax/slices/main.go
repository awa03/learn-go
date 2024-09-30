package main

import (
  "fmt"
  "slices"
)

// Note that while slices are different types than arrays, they are rendered similarly by fmt.Println.
func main(){
  var s[] string;
  fmt.Println("Created String", s);
  fmt.Println("Is the string nil by defaults", s == nil)

  s = make([]string , 3)
  fmt.Println("Length", len(s), "Cap", cap(s))

  s[0] = "a"
  fmt.Println("0 index now set ", s)
  s[1] = "b"
  fmt.Println("0-2 index now set ", s)
  s[2] = "c"
  fmt.Println("0 & 1 index now set ", s)
  
  fmt.Println("new length:", len(s))

  // copy
  c := make([]string, len(s))
  l := c[1:]
  fmt.Println("var l", l) 

  // declare an array
  t := []string{"g", "h", "i"}
  fmt.Println("dcl:", t)

  // slice 
  t2 := []string{"g", "h", "i"}
  if slices.Equal(t, t2) {
    fmt.Println("t == t2")
  }

  twoD := make([][]int, 3)
  for i := 0; i < 3; i++ {
    innerLen := i + 1
    twoD[i] = make([]int, innerLen)
    for j := 0; j < innerLen; j++ {
      twoD[i][j] = i + j
    }
  }
  fmt.Println("2d: ", twoD)
}
