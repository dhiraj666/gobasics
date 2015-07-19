/*

interfaces are named collections of
method signatures


*/

package main

import "fmt"
import "math"

type geometry interface {
    area() float64
  //  perim() float64
}

type rect struct {
    width, height float64
}

type circle struct {
    radius float64
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

func (r rect) area() float64 {

    return r.width * r.height
}

func measure(g geometry){
    fmt.Println(g)
    fmt.Println(g.area())
}

func main() {

    r := rect{ width : 20, height : 20}
    //fmt.Println(r.area())

    c := circle{ radius : 4}
    //fmt.Println(c.area())
    measure(r)
    measure(c)
}
