package utils

type Graph[T comparable] map[T][]T

func BFSPathFinder[T comparable](graph Graph[T], start T, endCondition func(n T) bool) [][]T {
  queue := [][]T{{start}}

  validPaths := [][]T{}

  visited := make(map[T]bool)

  for len(queue) > 0 {
    currentPath := queue[0]
    queue = queue[1:]

    currentNode := currentPath[len(currentPath)-1]

    if visited[currentNode] {
      continue
    }

    visited[currentNode] = true

    if endCondition(currentNode) {
      validPaths = append(validPaths, currentPath)
    }
    
    for _, neighbor := range graph[currentNode] {
      newPath := make([]T, len(currentPath))
      copy(newPath, currentPath)
      newPath = append(newPath, neighbor)

      if !Contains(currentPath, neighbor) {
        queue = append(queue, newPath)
      }
    }
  }

  return validPaths
}

func DFSPathFinder[T comparable](graph Graph[T], start T, endCondition func(n T) bool) [][]T {
	validPaths := [][]T{}

	var dfsRecursive func(T, []T)
	dfsRecursive = func(currentNode T, currentPath []T) {
		if endCondition(currentNode) {
			validPaths = append(validPaths, currentPath)
			return
		}

		for _, neighbor := range graph[currentNode] {
			if !Contains(currentPath, neighbor) {
				newPath := make([]T, len(currentPath))
				copy(newPath, currentPath)
				newPath = append(newPath, neighbor)

				dfsRecursive(neighbor, newPath)
			}
		}
	}

	dfsRecursive(start, []T{start})

	return validPaths
}

func Contains[T comparable](slice []T, item T) bool {
  for _, v := range slice {
    if v == item {
      return true
    }
  }
  return false
}

