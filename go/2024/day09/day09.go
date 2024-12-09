package main

import (
  "fmt"
  "slices"

  "github.com/kamillo/aoc/utils"
)


type File struct {
  id int
  length int
}

type Block struct {
  files []File 
  free int
}

func main() {
  lines := utils.GetLines("input.txt")
  // lines := utils.GetLines("test.txt")

  part1 := 0  
  part2 := 0

  disk := utils.ToIntArr(lines[0], "")

  part1 = checksum(defragment1(disk))
  part2 = checksum(defragment2(disk))

  fmt.Println("Part 1:", part1)
  fmt.Println("Part 2:", part2)
}

func parse(disk []int) []Block {
  fragmented := []Block{}
  id := 0
  for i := 0; i < len(disk); i += 2 {
    free := 0
    if i + 1 < len(disk) {
      free = disk[i+1]
    }

    file := []File{{id, disk[i]}}
    fragmented = append(fragmented, Block{file, free})
    id++
  }

  return fragmented
}

func defragment1(disk []int) []Block {
  fragmented := parse(disk)
  start := 0
  end := len(fragmented) - 1

  for ; start != end;  {
    if fragmented[start].free > 0 && fragmented[end].files[0].length > 0 {
      move := min(fragmented[start].free, fragmented[end].files[0].length)

      id := fragmented[end].files[0].id
      fragmented[start].free -= move
      fragmented[start].files = append(fragmented[start].files, File{id, move})
      fragmented[end].free += move
      fragmented[end].files[0].length -= move
    } else {
      if fragmented[start].free <= 0 {
        start++
      } else {
        end--
      }
    }
  }

  return fragmented
}


func defragment2(disk []int) []Block {
  fragmented := parse(disk)
  for f := range fragmented {
    fragmented[f].files = append(fragmented[f].files, File{-1, fragmented[f].free})
  }
  end := len(fragmented) - 1

  for ; end > 0;  {
    for fi, src := range slices.Backward(fragmented[end].files) {

      if (src.id == -1) { continue }

      for i := 0; i < end; i++ {
        found := false
        for dsti, dst := range fragmented[i].files {
          if dst.id == -1 && dst.length >= src.length {
            left := dst.length - src.length

            fragmented[i].files[dsti].id = src.id
            fragmented[i].files[dsti].length = src.length

            // add free space that left
            if left > 0 {
              fragmented[i].files = slices.Insert(fragmented[i].files, dsti+1, File{-1, left})
            }

            fragmented[end].files[fi].id = -1
            fragmented[end].free += src.length
            found = true
            break
          }

          if found { break }
        }

        if found { break }
      }
    }

    end--
  }

  return fragmented
}



func checksum(defragmented []Block) int {
  pos := 0
  ret := 0

  for _, d := range defragmented {
    for _, f := range d.files {
      for i := 0; i < f.length; i++ {
        if (f.id != -1) {
          ret += pos * f.id
        }
        pos++
      }

    }
  }

  return ret
}
