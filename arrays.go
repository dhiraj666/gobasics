/*

arrays is a numbered sequence of elements of a specific length

for setting a value
array[index] =

for retriveing a value
array[index]

there is a built in length function
for returning the length of the array
*/


package main

import "fmt"

func main() {

    var a [5]int
    fmt.Println(a)

    a[3] = 8
    fmt.Println(a)

    fmt.Println(a[3])

    fmt.Println("length :", len(a))

    b := [5]int{1,2,3,4,5}
    fmt.Println(b)

    var twoD [2][4]int

    for i := 0 ; i < 2 ; i++ {
        for j := 0 ; j < 4 ; j++ {
              twoD[i][j] = i + j
          }
    }

  fmt.Println("twoD : ",twoD)

}
