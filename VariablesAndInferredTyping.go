/*

variables and inferred typing

var statement declares the variables with the type declared last

*/

package main

import "fmt"

/*var (
  name string = "dhiraj kumar das"
  age int = 24
  location string = "karapakkam,chennai"

)

var (
    name, location string
    age int

)

var (
    name = "dhiraj kumar das"
    age = 24
    location = "karapakkam chhennai"

)

var (

   name, location, age = "dhiraj kumar das" , "karapakkam", 24

)
*/


func main() {
    name, location := "dhirajkumardas", "chennai"
    age := 24
    fmt.Printf("%s %s %d", name, location, age)

    action := func ()  {
        fmt.Println("dhiraj das ")
    }
    action()
}


// variable can contain any type even functions
