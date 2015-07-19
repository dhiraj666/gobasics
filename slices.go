/*

slices are useful
slices are typed only by the elements
they contain ( not the number of elements)

to create
use builtin make

it supports append , copy

also slice[low:high]
slices[upto]
slices[upfrom]
*/

package main

import "fmt"

func main() {

    s := make([]string, 3)

    fmt.Println(s)

    s[0] = "dhiraj"
    s[1] = "kumar"
    s[2] = "das"

    fmt.Println(s)
    fmt.Println(s[2])

    fmt.Println("length :", len(s))


    s = append(s, "nickname roshan")
    s = append(s, "yes no more nickname")

    fmt.Println("appended :", s)


    c := make([]string, len(s))
    copy(c, s)

    fmt.Println("copied : ", c)

    l := s[2:5]

    fmt.Println("slice 1 : ", l)

    l = s[2:]

    fmt.Println("slice 2 : ", l)

    l = s[:5]

    fmt.Println("slice 3: ", l)

    t := []string{"g", "h", "i"}

    fmt.Println("shortcut :" , t)



}
