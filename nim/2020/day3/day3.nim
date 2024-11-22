import utils/fileutils

type 
    Slope = object
        right: int
        down: int

proc countTrees(grid: seq[string], slope: Slope): int =
    var
        x = 0
        y = 0
        width = grid[0].len

    while y < grid.len:
        if grid[y][x mod width] == '#':
            result += 1
        y += slope.down
        x += slope.right

var 
    lines = getLines("input.txt")
    slopes: seq[Slope] = @[
        Slope(right: 1, down: 1),
        Slope(right: 3, down: 1),
        Slope(right: 5, down: 1),
        Slope(right: 7, down: 1),
        Slope(right: 1, down: 2),
    ]
    treesProduct = 1

echo("Part 1: ", countTrees(lines, Slope(right: 3, down: 1)))

for slope in slopes:
    treesProduct *= countTrees(lines, slope)

echo("Part 2: ", treesProduct)