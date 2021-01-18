package main

import (
  "fmt"
)

func main() {
  const level1 = 1_950_000
  const level2 = 3_330_000
  const level3 = 6_950_000
  const level4 = 9_000_000
  const level5 = 18_000_000
  const level6 = 40_000_000

  var income = 6_500_000

  if income > level1 {
    fmt.Println("収入はレベル1より高いです。")
  }
}
