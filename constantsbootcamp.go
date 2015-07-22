/*

constants are declared like variables but with a constantskeyord
constants can only be character string boolean or numeric
and cannot be declared using the :

*/
package main

import "fmt"

const Pi = 3.14

const (

    StatusOK  =                   200
    StatusCreated =               201
    StatusAccepted =              202
    StatusNonAuthoritativeInfo =  203
    StatusNoContent =             204
    StatusPartialContent =        205
    StatusResetContent =          206

)

const (
  Truth = false
  Big = 1 << 100
  Small = Big >> 99
)

func main() {
   const Greeting = "ハローワールド"
   fmt.Println( Greeting )
   fmt.Println( Pi )
   fmt.Println( Truth )
}
