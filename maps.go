/*

maps are go builtin associative datatype

to create empty maps
make(map[key type]valtype)

to set value

name[key] = value

len - lenght
delete - delete removes key value pairs


*/

package main

import "fmt"

func main() {

      m := make(map[string]int)

      m["key1"] = 6
      m["key2"] = 12

      fmt.Println("map : ", m)

      v1 := m["keyy1"]
      fmt.Println("v1:", v1)

      fmt.Println("len: ", len(m))

      delete(m, "key2")
      fmt.Println("after delete:", m)

      _, prs := m["key2"]
      fmt.Println("prs: ", prs)


      n := map[string]int{"keydd" : 1, "keyrr" : 3}
      fmt.Println("map :" , n)
}
