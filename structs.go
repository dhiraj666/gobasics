package main

import "fmt"

type person struct {
     name string
     age int
}

func main() {

    fmt.Println(person{"bob", 20})

    fmt.Println(person{name : "dhiraj", age : 24})

    fmt.Println(person{name : "roshan"})


    fmt.Println(&person{name : "rakesh", age: 30})
    s := person{ name : "shyam", age : 56}
    fmt.Println(s.name)

    sp := &s

    fmt.Println(sp.age)

    sp.age = 57
    fmt.Println(sp.age)

}
