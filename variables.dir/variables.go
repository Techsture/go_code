package main

import "fmt"

func main() {
  var a string = "initial"
  fmt.Println("A string:", a)

  var b, c int = 1, 2
  fmt.Println("Two integers:", b, ",", c)

  var d = true
  fmt.Println("A boolean:", d)

  var e int
  fmt.Println("An uninitialized integer:", e)

  f := "short"
  fmt.Println("Another string:",f)
}