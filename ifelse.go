package main

import "fmt"

func main() {

    if 81 % 9 == 0 {
      fmt.Println("it will print")
    } else {
        fmt.Println("sorry else, next time")
    }

    if 5 % 5 == 0 {
        fmt.Println("there is no else")
    }

    if num := 21; num < 0 {
        fmt.Println("negative")
    } else if num < 10 {
        fmt.Println("has 1 digit")
    } else {
        fmt.Println("has multiple digits")
    }

}
