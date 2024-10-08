package main

import (
  "fmt"
  "unicode/utf8"
)

func main(){
  const s = "สวัสดี"
  fmt.Println("Len:", len(s))

  for i := 0; i < len(s); i++ {
    fmt.Printf("%x\n", s[i])
  }
  fmt.Println("Rune Count:", utf8.RuneCountInString(s))

  for idx, runeValue := range s {
    fmt.Printf("%#U starts at %d\n", runeValue, idx)
  }
}
