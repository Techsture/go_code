package main

import (
  "fmt"
  "math/rand"
  "time"
)

func add(x int, y int) int {
  fmt.Println("X =", x, "Y =", y)
  return x + y
}

func main() {
  rand.Seed(time.Now().UnixNano())
  fmt.Println(add(rand.Intn(10), rand.Intn(10)))
}