package main

import (
  "fmt"
  "math"

  "github.com/kamillo/aoc/utils"
)


func main() {
  lines := utils.GetLines("input.txt")
  // lines := utils.GetLines("test.txt")

  input := utils.ToIntArr(lines[0], " ")

  part1 := 0  
  part2 := 0

  stones := map[int]int{}
  for _, i := range input {
    stones[i]++
  }

  part1 = run(stones, 25)
  part2 = run(stones, 75)
  fmt.Println("Part 1:", part1)
  fmt.Println("Part 2:", part2)
}


func run(stones map[int]int, blinks int) int {
  for i := 0; i < blinks; i++ {
    newStones := map[int]int{}

    for k, v := range stones {
      newStones[k] += v

      if k == 0 {
        newStones[k] -= v
        newStones[1] += v

      } else if utils.NumDigits(k) % 2 == 0 {
        l, r := splitNumber(k)
        newStones[l] += v
        newStones[r] += v
        newStones[k] -= v

      } else {
        newStones[k] -= v
        newStones[k * 2024] += v
      }
      if newStones[k] <= 0 {
        delete(newStones, k)
      }
    }

    stones = newStones
  }

  res := 0
  for _, v := range stones {
    res += v
  }

  return res
}

func splitNumber(num int) (int, int) {
  digits := utils.NumDigits(num)
  divisor := int(math.Pow(10, float64(digits/2)))

  leftHalf := num / divisor
  rightHalf := num % divisor

  return leftHalf, rightHalf
}
