package main

func main(){
  // i hate this but the else block must be on the same line as if closing -- same with if else
  if 7 % 2 == 0 {
    println("number is even")
  } else {
    println("number is odd")
  }
  
  if num:=9; num < 0 {
    println(num, "is negative")
  } else if num < 10 {
    println(num, "has 1 digit")
  } else {
    println(num, "has multiple digits")
  }
}
