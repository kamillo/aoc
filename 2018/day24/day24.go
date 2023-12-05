package main

import (
	"fmt"
	"sort"
)

const (
	nothing int = iota
	fire
	cold
	slashing
	radiation
	bludgeoning
)

const (
	immune int = iota
	infection
)

type Group struct {
	Id         int
	Type       int
	Units      int
	UnitHP     int
	DamageType int
	Damage     int
	Weakness   map[int]bool
	Immunity   map[int]bool
	Initiative int
}

type SelectionList []Group

func (a SelectionList) Len() int      { return len(a) }
func (a SelectionList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SelectionList) Less(i, j int) bool {
	iPower := a[i].Units * a[i].Damage
	jPower := a[j].Units * a[j].Damage

	if iPower == jPower {
		return a[i].Initiative < a[j].Initiative
	}

	return iPower < jPower
}

type FightList []Group

func (a FightList) Len() int      { return len(a) }
func (a FightList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a FightList) Less(i, j int) bool {
	return a[i].Initiative < a[j].Initiative
}

func main() {
	fight := func(boost int) (int, int) {
		// groups := map[int]Group{
		// 	1: {1, immune, 17, 5390, fire, 4507, map[int]bool{radiation: true, bludgeoning: true}, map[int]bool{}, 2},
		// 	2: {2, immune, 989, 1274, slashing, 25, map[int]bool{bludgeoning: true, slashing: true}, map[int]bool{fire: true}, 3},
		// 	3: {3, infection, 801, 4706, bludgeoning, 116, map[int]bool{radiation: true}, map[int]bool{}, 1},
		// 	4: {4, infection, 4485, 2961, slashing, 12, map[int]bool{fire: true, cold: true}, map[int]bool{radiation: true}, 4},
		// }
		groups := map[int]Group{
			1:  {1, immune, 9936, 1739, slashing, 1, map[int]bool{slashing: true, fire: true}, map[int]bool{}, 11},
			2:  {2, immune, 2990, 9609, cold, 31, map[int]bool{radiation: true}, map[int]bool{fire: true, cold: true}, 1},
			3:  {3, immune, 2637, 9485, radiation, 26, map[int]bool{bludgeoning: true}, map[int]bool{cold: true, slashing: true}, 13},
			4:  {4, immune, 1793, 2680, bludgeoning, 13, map[int]bool{bludgeoning: true}, map[int]bool{cold: true}, 10},
			5:  {5, immune, 8222, 6619, bludgeoning, 6, map[int]bool{}, map[int]bool{fire: true, slashing: true}, 12},
			6:  {6, immune, 550, 5068, radiation, 87, map[int]bool{}, map[int]bool{}, 19},
			7:  {7, immune, 950, 8681, slashing, 73, map[int]bool{radiation: true}, map[int]bool{}, 17},
			8:  {8, immune, 28, 9835, bludgeoning, 2979, map[int]bool{}, map[int]bool{}, 3},
			9:  {9, immune, 3799, 2933, slashing, 7, map[int]bool{}, map[int]bool{}, 16},
			10: {10, immune, 35, 8999, cold, 2505, map[int]bool{bludgeoning: true}, map[int]bool{radiation: true}, 6},
			11: {11, infection, 1639, 28720, cold, 27, map[int]bool{}, map[int]bool{}, 8},
			12: {12, infection, 4968, 16609, fire, 6, map[int]bool{}, map[int]bool{slashing: true, bludgeoning: true, radiation: true}, 2},
			13: {13, infection, 3148, 48970, slashing, 29, map[int]bool{fire: true, bludgeoning: true}, map[int]bool{}, 20},
			14: {14, infection, 1706, 30069, fire, 29, map[int]bool{}, map[int]bool{cold: true, bludgeoning: true}, 7},
			15: {15, infection, 496, 39909, bludgeoning, 133, map[int]bool{radiation: true}, map[int]bool{cold: true}, 4},
			16: {16, infection, 358, 17475, bludgeoning, 82, map[int]bool{}, map[int]bool{}, 5},
			17: {17, infection, 120, 53629, fire, 807, map[int]bool{}, map[int]bool{}, 15},
			18: {18, infection, 402, 44102, bludgeoning, 185, map[int]bool{slashing: true}, map[int]bool{}, 14},
			19: {19, infection, 468, 11284, radiation, 43, map[int]bool{fire: true}, map[int]bool{}, 18},
			20: {20, infection, 4090, 23075, bludgeoning, 10, map[int]bool{}, map[int]bool{radiation: true}, 9},
		}

		for k, g := range groups {
			if g.Type == immune {
				g.Damage += boost
				groups[k] = g
			}
		}

		for {
			selectionImmune := SelectionList{}
			for _, g := range groups {
				if g.Type == immune {
					selectionImmune = append(selectionImmune, g)
				}
			}

			selectionInfection := SelectionList{}
			for _, g := range groups {
				if g.Type == infection {
					selectionInfection = append(selectionInfection, g)
				}
			}

			sort.Sort(sort.Reverse(selectionImmune))
			sort.Sort(sort.Reverse(selectionInfection))

			match := map[int]int{}
			match = assingTargets(selectionImmune, selectionInfection, match)
			match = assingTargets(selectionInfection, selectionImmune, match)

			battle := FightList{}
			for _, g := range groups {
				if _, ok := match[g.Id]; ok {
					battle = append(battle, g)
				}
			}
			sort.Sort(sort.Reverse(battle))

			sumKilled := 0
			for _, b := range battle {
				attacker := b.Id
				if groups[attacker].Units > 0 {
					if defender, ok := match[attacker]; ok {
						damage := calculateDamage(groups[attacker], groups[defender])
						killed := int(damage / groups[defender].UnitHP)
						sumKilled += killed

						if killed > groups[defender].Units {
							killed = groups[defender].Units
						}

						fmt.Println(attacker, "attacks", defender, killed, "units killed")

						g := groups[defender]
						g.Units -= killed
						groups[defender] = g
					}
				}
			}

			army1 := 0
			army2 := 0
			for _, g := range groups {
				if g.Units <= 0 {
					delete(groups, g.Id)
				}

				if g.Type == immune && g.Units > 0 {
					army1 += g.Units
				}
				if g.Type == infection && g.Units > 0 {
					army2 += g.Units
				}
			}

			if army1 == 0 || army2 == 0 || sumKilled == 0 {
				// fmt.Println(army1, army2)
				return army1, army2
			}

			// fmt.Println()
		}
	}

	unitsImmune, unitsInf := fight(0)
	fmt.Println("Part 1:", unitsImmune, unitsInf)

	for i := 1; ; i++ {
		imm, inf := fight(i)
		fmt.Println(imm, inf)
		if imm > inf {
			fmt.Println("Part 2:", imm)
			break
		}

	}
}

func calculateDamage(attacker, defender Group) int {
	if defender.Immunity[attacker.DamageType] {
		return 0
	}

	effectivePower := attacker.Damage * attacker.Units
	if defender.Weakness[attacker.DamageType] {
		effectivePower *= 2
	}

	return effectivePower
}

func assingTargets(attackers, defenders []Group, match map[int]int) map[int]int {
	assigned := map[int]bool{}
	for _, attacker := range attackers {
		maxDamage := 0
		target := -1

		for i, defender := range defenders {
			if maxDamage < calculateDamage(attacker, defender) && !assigned[defender.Id] {
				maxDamage = calculateDamage(attacker, defender)
				target = i
			}
		}

		if target > -1 {
			match[attacker.Id] = defenders[target].Id
			assigned[defenders[target].Id] = true
		}
	}

	return match
}
