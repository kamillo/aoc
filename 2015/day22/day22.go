package main

import (
	"fmt"
)

// Hit Points: 55
// Damage: 8

type Spell struct {
	name       string
	cost       int
	effectTime int
}

func main() {

	spells := []string{"Missile", "Drain", "Shield", "Poison", "Recharge"}

	//moves := utils.CombinationsWithRepetitions(10, spells)
	//fmt.Println(moves)

	min := 1000000
	n := 11
	k := len(spells)
	pn := make([]int, n)
	p := make([]string, n)
	for {
		// generate permutaton
		for i, x := range pn {
			p[i] = spells[x]
		}
		//fmt.Println(p)
		//for _, move := range moves {
		effects := map[string]int{}
		manaSpent := 0
		mana := 500
		hp := 50
		bossHp := 55
		bossDamage := 8
		spells := []string{}

		for _, spell := range p {

			if manaSpent >= min {
				break
			}

			// skip if arleady active
			if effects[spell] > 1 {
				continue
			}

			spells = append(spells, spell)
			// Part2
			hp--
			if hp <= 0 {
				break
			}

			if effects["Poison"] > 0 {
				bossHp -= 3
			}

			if effects["Recharge"] > 0 {
				mana += 101
			}

			for k := range effects {
				effects[k]--
			}

			if bossHp <= 0 {
				// fmt.Println("boss dead ", manaSpent, spells)
				if manaSpent < min {
					min = manaSpent
				}
				break
			}

			switch spell {
			case "Missile":
				mana -= 53
				manaSpent += 53
				bossHp -= 4
				break

			case "Drain":
				mana -= 73
				manaSpent += 73
				bossHp -= 2
				hp += 2
				break

			case "Shield":
				mana -= 113
				manaSpent += 113
				effects[spell] = 6
				break

			case "Poison":
				mana -= 173
				manaSpent += 173
				effects[spell] = 6
				break

			case "Recharge":
				mana -= 229
				manaSpent += 229
				effects[spell] = 5
				break
			}

			if mana < 0 {
				break
			}

			if effects["Poison"] > 0 {
				bossHp -= 3
			}

			if bossHp <= 0 {
				// fmt.Println("boss dead ", manaSpent, spells)
				if manaSpent < min {
					min = manaSpent
				}
				break
			}

			if effects["Recharge"] > 0 {
				mana += 101
			}

			if effects["Shield"] > 0 {
				hp -= (bossDamage - 7)
			} else {
				hp -= bossDamage
			}

			if hp <= 0 {
				break
			}

			for k := range effects {
				effects[k]--
			}
		}

		for i := 0; ; {
			pn[i]++
			if pn[i] < k {
				break
			}
			pn[i] = 0
			i++
			if i == n {
				fmt.Println("Part 1: ", min)
				return // all permutations generated
			}
		}
	}
}
