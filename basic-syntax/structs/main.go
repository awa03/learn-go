package main

import "fmt"

type person struct {
  name string
  age int
}

func newPerson(name string) *person {
  p:= person{name: name}
  p.age = 42
  return &p
}

func main(){
  fmt.Println(person{"Aiden", 20})   
  fmt.Println(person{name: "Aiden", age: 20})
  fmt.Println(person{name: "Fred"})
  fmt.Println(&person{name: "Aiden", age:50})
  fmt.Println(newPerson("Jon"))

  s := person{name: "Sean", age: 50}
  fmt.Println(s.name)

  sp := &s
  fmt.Println(sp.age)

  dog := struct {
    name string
    isGood bool
  }{
    "Rex",
    true,
  }
  fmt.Println(dog)
}

