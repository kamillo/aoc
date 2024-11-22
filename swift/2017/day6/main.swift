//
//  main.swift
//  day2
//
//  Created by kamillo on 03/12/2020.
//

import Foundation

class Block {
    var value: Int = 0
    var next: Block?
    
    init(value: Int, next: Block?) {
        self.value = value
        self.next = next
    }
}

guard let lines = getLines(fromFile: "input.txt") else {
    exit(-1)
}

var blocks = lines[0].split(separator: " ").map { Int(String($0))! }
var snapshots: [String:Int] = [:]

var count = 0
while true {
    if snapshots.keys.contains(String(describing:blocks)) {
        break
    }
    snapshots[String(describing:blocks)] = count
    
    count += 1
    
    var value = (0, 0)
    for i in 0..<blocks.count {
        if blocks[i] > value.1 {
            value.1 = blocks[i]
            value.0 = i
        }
    }
    
    blocks[value.0] = 0
    
    for _ in 0..<value.1 {
        value.0 = (value.0 + 1) % blocks.count
        blocks[value.0] += 1
    }
}
print("Part 1:", count)
print("Part 2:", count - snapshots[String(describing:blocks)]!)

