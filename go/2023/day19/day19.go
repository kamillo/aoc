package main

import (
	"fmt"
	"image"
	"regexp"
	"strings"

	"github.com/kamillo/aoc/utils"
)

func main() {
	lines := utils.GetLines("test.txt")

	workflows := map[string][]string{}
	parts := []map[string]int{}
	rulesRegexp := regexp.MustCompile(`^(.*?)\{(.*?)\}$`)
	partsRegexp := regexp.MustCompile(`^\{(.*?)\}$`)
	accepted := []map[string]int{}

	rules := true
	for _, line := range lines {
		if line == "" {
			rules = false
			continue
		}

		if rules {
			matches := rulesRegexp.FindStringSubmatch(line)
			name := matches[1]
			rules := strings.Split(matches[2], ",")
			workflows[name] = rules

		} else {
			matches := partsRegexp.FindStringSubmatch(line)

			ratings := map[string]int{}
			split := strings.Split(matches[1], ",")
			for _, rating := range split {
				r := strings.Split(rating, "=")
				ratings[r[0]] = utils.JustAtoi(r[1])
			}

			parts = append(parts, ratings)
		}
	}

	for _, part := range parts {
		next := "in"
		for next != "R" && next != "A" {
			next = checkRule(workflows[next], part)
			// fmt.Println(part, next)
			if next == "A" {
				accepted = append(accepted, part)
			}
		}
	}

	sum := 0
	for _, part := range accepted {
		fmt.Println(part)
		for _, value := range part {
			sum += value
		}

	}

	fmt.Println("Part 1:", sum)

	routes := [][]map[string]string{}
	route := []map[string]string{}

	var next func(node string, route []map[string]string) string
	next = func(node string, route []map[string]string) string {
		for _, rule := range workflows[node] {
			s := strings.Split(rule, ":")

			if len(s) > 1 && s[1] == "A" || rule == "A" {
				if len(s) > 1 {
					route = append(route, map[string]string{s[1]: s[0]})
				}
				routes = append(routes, route)
				continue
			}

			if len(s) > 1 && s[1] == "R" || rule == "R" {
				continue
			}

			k := s[0]
			v := "*"
			if len(s) > 1 {
				k = s[1]
				v = s[0]
			}

			newRoute := make([]map[string]string, len(route))
			copy(newRoute, route)
			newRoute = append(newRoute, map[string]string{k: v})
			next(k, newRoute)
		}

		return ""
	}

	next("in", route)

	for _, route := range routes {
		rating := map[string]image.Point{}
		rating["x"] = image.Point{1, 4000}
		rating["m"] = image.Point{1, 4000}
		rating["a"] = image.Point{1, 4000}
		rating["s"] = image.Point{1, 4000}

		// modify rating accordding to route

		modifyR := func(rating map[string]image.Point, rule string) map[string]image.Point {
			if rule == "*" {
				return rating
			}

			splitLess := strings.Split(rule, "<")
			splitMore := strings.Split(rule, ">")

			if len(splitLess) == 2 {
				key := splitLess[0]
				value := utils.JustAtoi(splitLess[1])

				rating[key] = image.Point{rating[key].X, value}
			}

			if len(splitMore) == 2 {
				key := splitMore[0]
				value := utils.JustAtoi(splitMore[1])

				rating[key] = image.Point{value, rating[key].Y}
			}

			return rating
		}

		for _, r := range route {

		}

		// calculate combinations
		combinations := 1
		for _, r := range rating {
			combinations *= (r.Y - r.X)
		}

		fmt.Println(route)
	}
}

func checkRule(rules []string, parts map[string]int) string {
	for _, rule := range rules {
		splitLess := strings.Split(rule, "<")
		splitMore := strings.Split(rule, ">")

		if len(splitLess) == 2 {
			key := splitLess[0]

			valueStr := strings.Split(splitLess[1], ":")

			value := utils.JustAtoi(valueStr[0])
			rule := valueStr[1]

			if parts[key] < value {
				return rule
			}

		} else if len(splitMore) == 2 {
			key := splitMore[0]

			valueStr := strings.Split(splitMore[1], ":")
			value := utils.JustAtoi(valueStr[0])
			rule := valueStr[1]

			if parts[key] > value {
				return rule
			}

		} else {
			// fallback
			return rule
		}
	}

	return ""
}
