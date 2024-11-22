package main

import (
	"fmt"
	"strings"

	"github.com/kamillo/aoc/utils"
)

type Module struct {
	id                 string
	typ                string
	destinationModules []string
	on                 bool
	recivedLow         map[string]bool
}

var modules = make(map[string]Module)

func main() {
	lines := utils.GetLines("input.txt")

	for _, line := range lines {
		s := strings.Split(line, " -> ")
		id := s[0]
		typ := s[0]
		if id != "broadcaster" {
			typ = s[0][:1]
			id = s[0][1:]
		}

		destinationModules := strings.Split(s[1], ", ")

		modules[id] = Module{id, typ, destinationModules, false, map[string]bool{}}
	}

	for _, module := range modules {
		for _, m := range module.destinationModules {
			if _, ok := modules[m]; ok {
				modules[m].recivedLow[module.id] = true
			}
		}
	}

	totalLow := 0
	totalHigh := 0

	// rx is fed by &zg which is fed by &vm, &lm, &jd and &fv
	// rx <low- &zg
	// <h- &vm
	// <h- &lm
	// <h- &jd
	// <h- &fv
	//

	part2 := map[string]int{}

	for i := 0; ; i++ {
		nextModules = append(nextModules, Pulse{"button", "broadcaster", true})
		for len(nextModules) > 0 {
			for _, pulse := range nextModules {
				if pulse.low {
					if pulse.to == "vm" || pulse.to == "lm" || pulse.to == "jd" || pulse.to == "fv" {
						part2[pulse.to] = i + 1
					}
					totalLow++
				} else {
					totalHigh++
				}
			}

			sendPulse()
		}

		if i == 999 {
			fmt.Println("Part 1:", totalHigh*totalLow)
		}

		if len(part2) == 4 {
			sum := 1
			for _, v := range part2 {
				sum *= v
			}

			fmt.Println("Part 2:", sum)

			break
		}
	}

}

type Pulse struct {
	from string
	to   string
	low  bool
}

func (p Pulse) String() string {
	pulseName := "high"
	if p.low {
		pulseName = "low"
	}
	return fmt.Sprintf("%s --%s--> %s", p.from, pulseName, p.to)
}

var nextModules = []Pulse{}

func sendPulse() {
	tmpPulses := []Pulse{}
	for _, pulse := range nextModules {
		module := modules[pulse.to]
		id := module.id
		low := pulse.low
		from := pulse.from

		switch module.typ {
		case "%":
			if low {
				modules[id] = Module{module.id, module.typ, module.destinationModules, !module.on, map[string]bool{}}
				for _, module := range module.destinationModules {
					tmpPulses = append(tmpPulses, Pulse{id, module, !modules[id].on})
				}
			}

		case "&":
			modules[id].recivedLow[from] = low
			sendLow := true
			for _, recivedLow := range modules[id].recivedLow {
				sendLow = sendLow && !recivedLow
			}

			for _, module := range module.destinationModules {
				tmpPulses = append(tmpPulses, Pulse{id, module, sendLow})
			}

		case "broadcaster":
			for _, module := range modules[id].destinationModules {
				tmpPulses = append(tmpPulses, Pulse{id, module, low})
			}
		}
	}
	nextModules = tmpPulses
}
