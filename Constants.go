package main

import "fmt"
import "math"

const s string = "constantstring"

func main() {
    fmt.Println(s)

    const n = 500000000

    const ex = 3e20 / n

    fmt.Println(ex)

    fmt.Println(int64(ex))

    fmt.Println(math.Sin(n))

}
