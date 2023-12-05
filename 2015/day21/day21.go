package main

import (
	"fmt"

	"github.com/kamillo/aoc/utils"
)

// Hit Points: 109
// Damage: 8
// Armor: 2

type equipement struct {
	name   string
	cost   int
	damage int
	armor  int
}

type player struct {
	hitPoints int
	damage    int
	armor     int
}

func main() {
	boss := player{
		hitPoints: 109,
		damage:    8,
		armor:     2,
	}

	player1 := player{
		hitPoints: 100,
		damage:    0,
		armor:     0,
	}

	armors := parse("armor.txt")
	weapons := parse("weapon.txt")
	ring := parse("ring.txt")

	won := []int{}
	lost := []int{}
	for _, w := range weapons {
		for _, a := range armors {
			for _, r1 := range ring {
				for _, r2 := range ring {
					if r1 == r2 {
						continue
					}

					player1.damage = w.damage + r1.damage + r2.damage
					player1.armor = a.armor + r1.armor + r2.armor

					boss.hitPoints = 109
					player1.hitPoints = 100

					for boss.hitPoints > 0 && player1.hitPoints > 0 {
						boss.hitPoints -= (player1.damage - boss.armor)
						player1.hitPoints -= (boss.damage - player1.armor)

						if boss.hitPoints <= 0 {
							won = append(won, w.cost+a.cost+r1.cost+r2.cost)
						} else if player1.hitPoints <= 0 {
							lost = append(lost, w.cost+a.cost+r1.cost+r2.cost)
						}
					}
				}
			}
		}
	}

	fmt.Println("Part 1: ", utils.MinInArray(won))
	fmt.Println("Part 2: ", utils.MaxInArray(lost))
}

func parse(name string) []equipement {
	result := []equipement{}
	for _, line := range utils.GetLines(name) {
		e := equipement{}
		fmt.Sscanf(line, "%s %d %d %d", &e.name, &e.cost, &e.damage, &e.armor)
		result = append(result, e)
	}

	return result
}
