package main

import "fmt"

func sum(nums ...int) {

    total := 0
    for _, item := range nums {
        total += item
    }
    fmt.Println(total)
}

func main() {
    sum(2, 3, 4)
    sum(1)
    sum(4, 3, 6)

    numslice := []int{1,2,3,4,5,6}
    sum(numslice...)
}
