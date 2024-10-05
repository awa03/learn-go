package main

import (
  "errors" // for custom error handlinga
  "fmt"    // for print statements
)

type argError struct {
  arg     int       // Error argument
  message string    // Error message
}

// allows error interface to be implemented
func (e* argError) Error() string {
  return fmt.Sprintf("%d - %s", e.arg,e.message);
}

// return error if number is 42
func fourtytwo(arg int) (int, error) {
  if arg == 42 {
    return -1, &argError{arg, "cant work with it"}
  }
  return arg + 1, nil
}

func main(){
  fmt.Println("Enter a number: ")
  num := 42
  fmt.Scanf("%d", num)
  _, err := fourtytwo(num) // Error as value
  var ae* argError         // Make a new arument error
  if errors.As(err, &ae) { 
    fmt.Println(ae.arg)
    fmt.Println(ae.message)
  } else {
    fmt.Println("Err does not match argError")
  }

  // Own example
  Own()

}

