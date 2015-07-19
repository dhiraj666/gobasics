package main

import "fmt"

func plus(a int, b int) int {

      return a + b
}

func plusplus(a, b, c int) int {

      return a + b + c
}

func main() {
    p := plus(4, 3)
    fmt.Println(p)
    pp := plusplus(4, 3, 2)
    fmt.Println(pp)

}
