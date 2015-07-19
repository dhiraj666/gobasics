package main

import "fmt"

type rect struct {

    width, height int
}

func (r *rect) area() int {

      return r.width * r.height
}

func main() {

    s := rect{ width : 20, height : 40}
    fmt.Println(s.area())

}
