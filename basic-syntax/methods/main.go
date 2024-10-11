package main
import (
  "fmt"
)

type rect struct {
  width, height int
}

func (r* rect) area() int {
  return r.width * r.height 
}

func (r rect) perim() int {
  return 2 * r.width + 2 * r.height 
}

func main(){
  r := rect{10, 20}
  fmt.Println("Rectangle: (", r.height, ",", r.width, ")")
  fmt.Println("Area:", r.area())
  fmt.Println("Perim:", r.perim())

  fmt.Println("Trying with a pointer now")
  rp := &r
  fmt.Println("Rectangle: (", rp.height, ",", rp.width, ")")
  fmt.Println("Area:", rp.area())
  fmt.Println("Perim:", rp.perim())
}
