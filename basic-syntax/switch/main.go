package main

import (
  "time"
  "fmt"
)

func main(){
  i := 2
  switch i {
    case 0:
      println("number is 0")
    case 1:
      println("number is 1")
    case 2:
      println("number is 2")
    default: 
      break
  }
  switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
      println("It's the weekend")
    default:
      println("It's a weekday")
  }

  whatAmI := func(i interface{}) {
    switch t := i.(type) {
      case bool:
        fmt.Println("I'm a bool")
      case int:
        fmt.Println("I'm an int")
      default:
        fmt.Printf("Don't know type %T\n", t)
    }
  }

  whatAmI(true)
  whatAmI(44)
  whatAmI(423.342)

}
