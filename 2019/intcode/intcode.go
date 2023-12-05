package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

// ABCDE
// 1002
//
// DE - two-digit opcode,      02 == opcode 2
// C - mode of 1st parameter,  0 == position mode  , 2 == relative mode
// B - mode of 2nd parameter,  1 == immediate mode
// A - mode of 3rd parameter,  0 == position mode,

type IntCode struct {
	Codes []int
	input []int
	base  int
	ptr   int
}

func ParseInput(commands string) []int {
	splitted := strings.Split(commands, ",")
	ints := make([]int, len(splitted)*1000)
	for i := 0; i < len(splitted); i++ {
		ints[i], _ = strconv.Atoi(splitted[i])
	}

	return ints
}

func Make(codes []int) IntCode {
	i := IntCode{}
	//i.Codes = make([]int, len(codes))
	i.Codes = append(i.Codes, codes...)
	i.input = make([]int, 0)
	return i
}

func (c *IntCode) Put(in []int) {
	c.input = append(c.input, in...)
}

func (c *IntCode) GetInput() []int {
	return c.input
}

func (c *IntCode) GetLine() (string, State) {
	line := ""
	for {
		if out, state := c.Get(); out == 10 || state != Output {
			return line, state
		} else {
			//fmt.Println(out, string(out))
			line += string(out)
		}
	}
}

func (c *IntCode) PutLine(in string) {
	runes := []rune(in)
	var result []int

	for i := 0; i < len(runes); i++ {
		result = append(result, int(runes[i]))
	}

	result = append(result, '\n')
	c.Put(result)
}

func (c IntCode) Run(debug bool) {
	for {
		if res, state := c.Get(); state != Exit {
			if debug {
				fmt.Println(res)
			}
		} else {
			return
		}
	}
}

type State int

const (
	Output State = 0
	Input  State = 1
	Exit   State = 2
)

func (c *IntCode) Get() (out int, state State) {
	ret := 0

	for {
		pparam1, param2, param3 := 0, 0, 0

		if c.Codes[c.ptr]%100 == 99 {
			return ret, Exit
		}

		if c.ptr+1 < len(c.Codes) {
			pparam1 = c.ptr + 1
			if ((c.Codes[c.ptr] / 100) % 10) == 0 {
				if c.Codes[c.ptr+1] < len(c.Codes) {
					pparam1 = c.Codes[c.ptr+1]
				}
			} else if ((c.Codes[c.ptr] / 100) % 10) == 2 {
				if c.base+c.Codes[c.ptr+1] < len(c.Codes) {
					pparam1 = c.base + c.Codes[c.ptr+1]
				}
			}
		}

		if c.ptr+2 < len(c.Codes) {
			param2 = c.Codes[c.ptr+2]
			if ((c.Codes[c.ptr] / 1000) % 10) == 0 {
				if c.Codes[c.ptr+2] < len(c.Codes) {
					param2 = c.Codes[c.Codes[c.ptr+2]]
				}
			} else if ((c.Codes[c.ptr] / 1000) % 10) == 2 {
				if c.base+c.Codes[c.ptr+2] < len(c.Codes) {
					param2 = c.Codes[c.base+c.Codes[c.ptr+2]]
				}
			}
		}

		if c.ptr+3 < len(c.Codes) {
			param3 = c.ptr + 3
			if ((c.Codes[c.ptr] / 10000) % 10) == 0 {
				if c.Codes[c.ptr+3] < len(c.Codes) {
					param3 = c.Codes[c.ptr+3]
				}
			} else if ((c.Codes[c.ptr] / 10000) % 10) == 2 {
				param3 = c.base + c.Codes[c.ptr+3]
			}
		}
		param1 := c.Codes[pparam1]
		switch c.Codes[c.ptr] % 100 {
		case 1:
			c.Codes[param3] = param1 + param2
			c.ptr += 4
		case 2:
			c.Codes[param3] = param1 * param2
			c.ptr += 4
		case 3:
			if len(c.input) == 0 {
				return ret, Input
			}
			c.Codes[pparam1] = c.input[0]
			c.input = append(c.input[:0], c.input[1:]...)
			c.ptr += 2
		case 4:
			//fmt.Println("out: ", param1)
			ret = param1
			c.ptr += 2
			return param1, Output
		case 5: // jump-if-true: if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			c.ptr += 3
			if param1 != 0 {
				c.ptr = param2
			}
		case 6: // jump-if-false: if the first parameter is zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
			c.ptr += 3
			if param1 == 0 {
				c.ptr = param2
			}
		case 7: // less than: if the first parameter is less than the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			c.Codes[param3] = 0
			if param1 < param2 {
				c.Codes[param3] = 1
			}
			c.ptr += 4
		case 8: // equals: if the first parameter is equal to the second parameter, it stores 1 in the position given by the third parameter. Otherwise, it stores 0.
			c.Codes[param3] = 0
			if param1 == param2 {
				c.Codes[param3] = 1
			}
			c.ptr += 4
		case 9:
			c.base += param1
			c.ptr += 2
		case 99: // halt
			return ret, Exit
		}
	}
}
