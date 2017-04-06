package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  fmt.Println("Here is a list of 7 random numbers:", rand.Perm(7))
}