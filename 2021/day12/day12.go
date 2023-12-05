package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Graph map[string]map[string]bool

func main() {
	graph := Graph{}

	for _, line := range utils.GetLines("input.txt") {
		split := strings.Split(line, "-")

		if len(graph[split[0]]) == 0 {
			graph[split[0]] = make(map[string]bool)
		}
		if len(graph[split[1]]) == 0 {
			graph[split[1]] = make(map[string]bool)
		}

		graph[split[0]][split[1]] = true
		graph[split[1]][split[0]] = true
	}

	paths := [][]string{}
	path := []string{}
	visited := map[string]bool{}

	paths = findAllPaths(graph, "start", "end", path, visited)
	fmt.Println("Part 1: ", len(paths))

	path = []string{}
	visited = map[string]bool{}

	findAllPaths2(graph, "start", "end", path, visited, "")
	fmt.Println("Part 2: ", cnt)
}

func findAllPaths(graph Graph, start string, end string, path []string, visited map[string]bool) [][]string {
	path = append(path, start)
	if start == strings.ToLower(start) {
		visited[start] = true
	}
	if start == end {
		return [][]string{path}
	}

	paths := [][]string{}
	for node := range graph[start] {
		if !visited[node] {
			newVisited := map[string]bool{}
			for k, v := range visited {
				newVisited[k] = v
			}
			newPaths := findAllPaths(graph, node, end, path, newVisited)
			for _, newPath := range newPaths {
				paths = append(paths, newPath)
			}
		}
	}
	return paths
}

var cnt int = 0

func findAllPaths2(graph Graph, start string, end string, path []string, visited map[string]bool, twice string) {
	if start == end {
		cnt++
		return
	}
	if visited[start] && start == "start" {
		return
	}

	path = append(path, start)
	if start == strings.ToLower(start) {
		visited[start] = true
	}

	for node := range graph[start] {
		ok := false
		if visited[node] && twice == "" {
			twice = node
			ok = true
		}
		if !ok && !visited[node] && node != "start" {
			ok = true
		}

		if ok {
			newVisited := map[string]bool{}
			for k, v := range visited {
				newVisited[k] = v
			}

			findAllPaths2(graph, node, end, path, newVisited, twice)
			if node == twice {
				twice = ""
			}
		}
	}
	return
}
