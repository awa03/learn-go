package main
import "fmt"

func main(){
  // Print 1 - 3
  i := 0
  for i <= 4 {
    fmt.Println(i)
    i++;
  }

  for j:=0; j <= 4; j++ {
    fmt.Println(j)
  }

  for i := range 4 {
    fmt.Println("range", i)
  }

  for {
    fmt.Println("empty for loop")
    break
  }

  for n := range 6 {
    if n%2 == 0 {
      continue
    }
    fmt.Println(n)
  }
}


