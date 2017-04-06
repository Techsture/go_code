package main

import (
  "fmt"
  "math"
)

const DELTA = 0.00000001
const INITIAL_Z = 1.0

func Sqrt(x float64) float64 {
  z := INITIAL_Z
  step := func() float64 {
    return z - (z*z - x) / (2 * z)
  }
  for zz := step(); math.Abs(zz - z) > DELTA; {
    z = zz
    zz = step()
  }
  return z
}

func main() {
  fmt.Println("Newtonian squareroot:\t\t\t", Sqrt(2))
  fmt.Println("Calculated math function squareroot:\t", math.Sqrt(2))
}