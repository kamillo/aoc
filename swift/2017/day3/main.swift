//
//  main.swift
//  day2
//
//  Created by kamillo on 03/12/2020.
//

import Foundation

let num = 277678

var gridSize = 0
var nearestSqrt = 0

if Int(ceil(sqrt(Double(num)))) % 2 == 0 {
    gridSize = Int(ceil(sqrt(Double(num))) + 1)
} else {
    gridSize = Int(ceil(sqrt(Double(num))))
}

nearestSqrt = (gridSize - 2) * (gridSize - 2)

var x = 0
var y = 0
var spiralSide = 0
let remainingEl = (num - nearestSqrt) % (gridSize - 1)

if ( remainingEl != 0) {
    spiralSide = (num - nearestSqrt) / gridSize
    
    switch spiralSide {
    case 0, 2:
        x = Int(gridSize / 2)
        y = abs(Int(gridSize / 2) - remainingEl)
        break;
        
    case 1, 3:
        x = abs(Int(gridSize / 2) - remainingEl)
        y = Int(gridSize / 2)
        break;
        
    default:
        break;
    }
    
} else {
    x = Int(gridSize / 2)
    y = Int(gridSize / 2)
}

print("Part 1:", x + y)

var system: [String:Int] = [:]

var value = 1
var position = (x: 0, y: 0)
var step = -1

func sumNeighbors(pos: (x: Int, y: Int)) -> Int {
    var sum = 0
    
    if pos == (0, 0) {
        return 1
    }
    
    sum += system[String(describing:(pos.x + 1, pos.y))] ?? 0
    sum += system[String(describing:(pos.x + 1, pos.y + 1))] ?? 0
    sum += system[String(describing:(pos.x, pos.y + 1))] ?? 0
    sum += system[String(describing:(pos.x - 1, pos.y))] ?? 0
    sum += system[String(describing:(pos.x - 1, pos.y - 1))] ?? 0
    sum += system[String(describing:(pos.x, pos.y - 1))] ?? 0
    sum += system[String(describing:(pos.x - 1, pos.y + 1))] ?? 0
    sum += system[String(describing:(pos.x + 1, pos.y - 1))] ?? 0

    return sum
}

let X = 277678
x = 0
y = 0
var dx = 0
var dy = -1
for _ in 0..<X*X {
    if (-X/2 <= x && x <= X/2) && (-X/2 <= y && y <= X/2) {
        value = sumNeighbors(pos: (x: x, y: y))
        system[String(describing:(x,y))] = value
        
        if value > X {
            print("Part 2: \(value)")
            break
        }
    }
    
    if x == y || (x < 0 && x == -y) || (x > 0 && x == 1-y) {
        let t = dx
        dx = -dy
        dy = t
    }
    
    x += dx
    y += dy
}
