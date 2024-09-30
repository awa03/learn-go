package main

import (
	"fmt"
	"maps"
)

func main(){
  // Phone contacts
  m := make(map[string]string) 
  m["aiden"] = "432-613-6143"
  m["joey"] = "829-521-6342"

  fmt.Println("Contacts map", m)

  v1 := m["joey"]
  fmt.Println("joey:", v1)

  _, v2 := m["aiden"]
  fmt.Println("aiden?:", v2)
  
  fmt.Println("before deletion", m)
  delete(m, "aiden")
  fmt.Println("after deletion", m)

  fmt.Println("Size:", len(m))

  fmt.Println("Before Clearing:", m)
  clear(m)
  fmt.Println("After Clearing:", m)

  n := map[string]int {"aiden": 10, "joey": 13, "greg" : 2 }
  fmt.Println(n)

  fmt.Println("making same map")
  d := map[string]int {"aiden": 10, "joey": 13, "greg" : 2 }

  if maps.Equal(d, n) {
    fmt.Println("They are equal")
  }
} 
