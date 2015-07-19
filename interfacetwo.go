package main

import "fmt"

type Animal interface {

    Speak() string

}

type Dog struct {

}

func (d Dog) Speak() string {

      return "boww!!"
}

type Cat struct {

}

func (c Cat) Speak() string {

      return "meoww!!"
}

type JavaProgrammer struct {

}

func (j JavaProgrammer) Speak() string {

      return "mylifeiscomingtoend"
}


type Golang struct {

}

func (g Golang) Speak() string {

      return "watchout gophers"
}


func main() {

    animals := []Animal{Dog{}, Cat{}, JavaProgrammer{}, Golang{}}


    for _, item := range animals {
        fmt.Println(item.Speak())
    }
}
