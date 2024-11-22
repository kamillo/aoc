from strutils import parseInt
import utils/fileutils

var 
    first = 0
    second = 0
    lines = getLines("input.txt")
    numbers: seq[int]

for line in lines:
    