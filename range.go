package main

import "fmt"

func main() {

      nums := []int{2,3,4}
      sum := 0

      for _, item := range nums {
          sum += item
      }

      fmt.Println("sum: ", sum)

      for index, item := range nums {
          if item == 3 {
               fmt.Println("index:" ,nums[index])
          }
      }

      maps := make(map[string]string)

      maps["d"] = "dhiraj"
      maps["r"] = "roshan"
      maps["g"] = "golang"

      for index, item := range maps {
         fmt.Printf("%s - > %s\n", index, item)
      }

      for index, item := range "golanggophers" {
        fmt.Println(index,item)
      }



}
