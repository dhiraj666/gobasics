package main

import "fmt"

const(
    height = 5
    mass = "40 kg"
)

var (
    phone = "Moto 4G LTE"
    price = 7999
)

func main() {
    mainvar := "mainvariable"
    name := "dhiraj - twentyone"

    fmt.Println(mainvar)
    aka := fmt.Sprintf("Number %d", 21)
    fmt.Printf("%s is also known as %s", name, aka)
    fmt.Printf("%d and %s" , height, mass)
}
