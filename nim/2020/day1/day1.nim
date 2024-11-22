from strutils import parseInt
import utils/fileutils

var 
    first = 0
    second = 0
    lines = getLines("input.txt")
    numbers: seq[int]

for line in lines:
    numbers.add(parseInt(line))

for i, number1 in numbers:
    for j, number2 in numbers:
        if i != j and first == 0 and number1+number2 == 2020:
            first = number1 * number2        
        for k, number3 in numbers:
            if i != j and j != k and i != k and number1 + number2 + number3 == 2020:
                second = number1 * number2 * number3
                break
        if second != 0 and first != 0:
            break
    if second != 0 and first != 0:
        break
        
echo("Part 1: ", first)
echo("Part 2: ", second)