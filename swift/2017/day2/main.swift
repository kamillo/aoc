//
//  main.swift
//  day2
//
//  Created by kamillo on 03/12/2020.
//

import Foundation
if let lines = getLines(fromFile: "input.txt") {
    let lineDiffs = lines.reduce(into: []) { (diffs, line) in
        let lineSorted = String(line).split(separator: " ").map { Int(String($0))! }.sorted()
        diffs.append(lineSorted.last! - lineSorted.first!)
    } as [Int]
    
    var evenlyDivisibleValues = [Int]()
    lines.forEach( { (line) in
        line.split(separator: " ").forEach( { (a) in
            line.split(separator: " ").forEach( { (b) in
                if a != b, let a = Int(a), let b = Int(b) {
                    if a % b == 0 {
                        evenlyDivisibleValues.append(Int(a/b))
                    }
                }
            })
        })
    })
    
    print("Part 1:", lineDiffs.reduce(0, +))
    print("Part 2:", evenlyDivisibleValues.reduce(0, +))
}

