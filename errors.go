package main

import "errors"
import "fmt"

func f1(arg int) (int, error) {

      if arg == 42 {
          return -1, errors.New("cant work with 42")
      }

      return arg + 3, nil
}

type argError struct {
    arg int
    prob string
}

func (e *argError) Error() string {

      return fmt.Sprintf("%d - %s",e.arg, e.prob)
}

func f2(arg int) (int, error) {

    if arg == 42 {
        return -1 , &argError{arg,"cant work with it"}

    }

    return arg + 3, nil
}


func main() {

  for _, item := range []int{7, 42} {
      if r, e := f1(item); e != nil {
          fmt.Println("f1 failed: ", e)
      } else {
          fmt.Println("f1 worked:", r)

      }
  }

  for _, item := range []int{7, 42} {
      if r, e := f2(item); e != nil {
          fmt.Println("f1 failed: ", e)
      } else {
          fmt.Println("f1 worked:", r)

      }
  }

  _, e := f2(42)

  if ae, ok := e.(*argError); ok {
      fmt.Println(ae.arg)
      fmt.Println(ae.prob)
    } 

}
